package server

import (
	"context"
	"errors"
	_ "github.com/Jeffail/gabs/v2"
	"github.com/firekitz/fk-daemon-iam/config"
	postgres "github.com/firekitz/fk-daemon-iam/internal/Database"
	iampb "github.com/firekitz/fk-daemon-iam/internal/proto/iam"
	log "github.com/firekitz/fk-lib-log-go"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
	"strings"
	"time"
)

const (
	ERROR_INVALID_TOKEN = 1
	ERROR_EXPIRED_TOKEN = 2
)

const (
	USER_ACCOUNT    = 0
	SERVICE_ACCOUNT = 1
)

type iamServer struct {
	iampb.IamServer
}

type FKClaims struct {
	ProjectId   int64  `json:"pi,omitempty"`
	GroupId     int64  `json:"gi,omitempty"`
	DomainId    int64  `json:"di,omitempty"`
	AccountId   int64  `json:"ai,omitempty"`
	AccountType int64  `json:"at,omitempty"`
	TokenType   string `json:"type,omitempty"`
	jwt.StandardClaims
}

type AuthTokenClaims struct {
	UserID string   `json:"id"`
	Name   string   `json:"name"`
	Email  string   `json:"mail"`
	Role   []string `json:"role"`

	jwt.StandardClaims
}

type UserGroupUser struct {
	UserGroupId int64 `json:"ai,omitempty"`
}

type AccountGroupUser struct {
	GroupId int64 `json:"ai,omitempty"`
}

type Permission struct {
	PermissionId int64 `json:"ai,omitempty"`
}

func CheckError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func contains(s []Permission, e int64) bool {
	for _, a := range s {
		if a.PermissionId == e {
			return true
		}
	}
	return false
}

func TestTokenBuild() {
	at := FKClaims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt: 1636680665,
			Subject:  "0",
		},
	}

	atoken := jwt.NewWithClaims(jwt.SigningMethodHS256, &at)
	signedAuthToken, err := atoken.SignedString([]byte("fk-jwt-secret"))

	if err != nil {
		log.I("permissions: %s", err)
	}
	//var js = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIwIiwiaWF0IjoxNjM2NjgwODk5fQ.TGKBvHW3ged2bPorIkj93ZpV9sWD9J35SkWNu_5L6GM"
	//var go = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2MzY2ODA2NjUsInN1YiI6IjAifQ.38_AAA4kdG1qwqGEpfJ1q76e3oO3qL6Tb-yiWBbEwKs"

	var newToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOjAsImlhdCI6MTYzNjYxNzI5NywiZXhwIjoyMjY3NzY5MjkwLCJkaSI6MTAwLCJwaSI6NSwiYWkiOjIsImF0IjoxLCJ0eXBlIjoiYWNjZXNzIn0.p4ax681nuiGE-vdEUpMOxD9Sn6peFbiy6HnUrtVp4-c"

	key := func(token *jwt.Token) (interface{}, error) {
		logrus.Infof("in key")
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			var ErrUnexpectedSigningMethod = errors.New("unexpected signing method")
			log.I("in key 0")
			return nil, ErrUnexpectedSigningMethod
		}
		log.I("in key 1")
		return []byte("ZHR4LWp3dC1zZWNyZXQ="), nil
	}

	tok, err := jwt.ParseWithClaims(newToken, &FKClaims{}, key)
	println(tok.Valid)

	log.I("permissions: %s", signedAuthToken)
}

func (server *iamServer) Auth(ctx context.Context, req *iampb.AuthRequest) (*iampb.AuthResponse, error) {
	accessToken := req.GetAccessToken()
	permissions := req.GetPermissions()
	logrus.Infof("accessToken: %s", accessToken)
	logrus.Infof("permissions: %s", permissions)

	token, err := jwt.ParseWithClaims(accessToken, &FKClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			ErrUnexpectedSigningMethod := errors.New("unexpected signing method")
			return nil, ErrUnexpectedSigningMethod
		}
		return []byte(config.LoadedConfig.JWT_SECRET), nil
	})
	if token == nil {
		return &iampb.AuthResponse{
			StatusCode:   ERROR_INVALID_TOKEN,
			ErrorMessage: err.Error(),
		}, nil
	}

	if err != nil {
		if strings.Contains(err.Error(), "expired") {
			return &iampb.AuthResponse{
				StatusCode:   ERROR_EXPIRED_TOKEN,
				ErrorMessage: "Expired Token",
			}, nil
		} else {
			return &iampb.AuthResponse{
				StatusCode:   ERROR_INVALID_TOKEN,
				ErrorMessage: err.Error(),
			}, nil
		}
	}

	claims := token.Claims.(*FKClaims)
	i, err := strconv.ParseInt(strconv.FormatInt(claims.ExpiresAt, 10), 10, 64)
	if err != nil {
		return &iampb.AuthResponse{
			StatusCode:   ERROR_INVALID_TOKEN,
			ErrorMessage: err.Error(),
		}, nil
	}

	tm := time.Unix(i, 0)
	if tm.Before(time.Now()) {
		return &iampb.AuthResponse{
			StatusCode:   ERROR_EXPIRED_TOKEN,
			ErrorMessage: "Expired Token",
		}, nil
	}

	if token.Valid != true {
		return &iampb.AuthResponse{
			StatusCode:   ERROR_INVALID_TOKEN,
			ErrorMessage: "Invalid Token",
		}, nil
	}

	accountId := claims.AccountId
	projectId := claims.ProjectId
	groupId := claims.GroupId
	domainId := claims.DomainId
	accountType := claims.AccountType
	issuedAt := claims.IssuedAt
	expiresAt := claims.ExpiresAt

	var rows []Permission
	query := ""

	if accountType == USER_ACCOUNT {
		var v UserGroupUser
		result := postgres.ProjectDB.Table("group_user").
			Select("group_id").
			Where("project_id = ?", projectId).
			Where("user_id = ?", accountId).
			First(&v)
		if result.Error != nil {
			if result.Error.Error() == "record not found" {
				return nil, status.Error(codes.NotFound, "User not found")
			}
			return nil, status.Error(codes.Internal, result.Error.Error())
		}
		groupId := v.UserGroupId
		query = `
			SELECT
				 DISTINCT RP.permission_id
			FROM project_group_role AS PGR
				  JOIN role R on R.id = PGR.role_id
				  JOIN role_permission RP on R.id = RP.role_id
			WHERE PGR.project_group_id = ` + strconv.Itoa(int(groupId)) + `
		`
	} else {
		query = `
			SELECT
				 DISTINCT RP.permission_id
			FROM role_member AS RM
		    JOIN role R on R.id = RM.role_id
		    JOIN role_permission RP on R.id = RP.role_id
			WHERE RM.account_id = ` + strconv.Itoa(int(accountId))
	}

	result := postgres.IamDB.Raw(query).Scan(&rows)
	if result.Error != nil {
		return nil, status.Error(codes.Internal, result.Error.Error())
	}

	for _, permissionId := range permissions {
		if !contains(rows, permissionId) {
			return nil, status.Error(codes.PermissionDenied, "Has no permissions")
		}
	}

	return &iampb.AuthResponse{
		StatusCode:  200,
		DomainId:    domainId,
		ProjectId:   projectId,
		GroupId:     groupId,
		AccountId:   accountId,
		AccountType: accountType,
		IssuedAt:    issuedAt,
		ExpiresAt:   expiresAt,
	}, nil
}

func (server *iamServer) VerifyToken(ctx context.Context, req *iampb.VerifyTokenRequest) (*iampb.VerifyTokenResponse, error) {
	accessToken := req.GetAccessToken()
	logrus.Infof("accessToken: %s", accessToken)

	token, err := jwt.ParseWithClaims(accessToken, &FKClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			ErrUnexpectedSigningMethod := errors.New("unexpected signing method")
			return nil, ErrUnexpectedSigningMethod
		}
		return []byte(config.LoadedConfig.JWT_SECRET), nil
	})
	if token == nil {
		CheckError(err)
	}
	if err != nil {
		expired := strings.Contains(err.Error(), "expired")
		if expired {
			return &iampb.VerifyTokenResponse{
				StatusCode:   ERROR_EXPIRED_TOKEN,
				ErrorMessage: err.Error(),
			}, nil
		}
		return &iampb.VerifyTokenResponse{
			StatusCode:   ERROR_INVALID_TOKEN,
			ErrorMessage: err.Error(),
		}, nil
	}

	claims := token.Claims.(*FKClaims)
	i, err := strconv.ParseInt(strconv.FormatInt(claims.ExpiresAt, 10), 10, 64)
	CheckError(err)

	tm := time.Unix(i, 0)
	if token.Valid != true {
		return &iampb.VerifyTokenResponse{
			StatusCode:   ERROR_INVALID_TOKEN,
			ErrorMessage: err.Error(),
		}, nil
	}
	if tm.Before(time.Now()) {
		return &iampb.VerifyTokenResponse{
			StatusCode:   ERROR_EXPIRED_TOKEN,
			ErrorMessage: err.Error(),
		}, nil
	}

	accountId := claims.AccountId
	projectId := claims.ProjectId
	groupId := claims.GroupId
	domainId := claims.DomainId
	accountType := claims.AccountType
	issuedAt := claims.IssuedAt
	expiresAt := claims.ExpiresAt

	return &iampb.VerifyTokenResponse{
		DomainId:    domainId,
		ProjectId:   projectId,
		GroupId:     groupId,
		AccountId:   accountId,
		AccountType: accountType,
		IssuedAt:    issuedAt,
		ExpiresAt:   expiresAt,
	}, nil
}

func (server *iamServer) CreateToken(ctx context.Context, req *iampb.CreateTokenRequest) (*iampb.CreateTokenResponse, error) {

	currentTime := time.Now()
	accountId := req.GetAccountId()

	at := FKClaims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  currentTime.Unix(),
			ExpiresAt: currentTime.Add(time.Minute * 30).Unix(),
			Subject:   strconv.Itoa(int(accountId)),
		},
		ProjectId:   req.GetProjectId(),
		GroupId:     req.GetGroupId(),
		DomainId:    req.GetDomainId(),
		AccountId:   req.GetAccountId(),
		AccountType: req.GetAccountType(),
		TokenType:   "access",
	}

	atoken := jwt.NewWithClaims(jwt.SigningMethodHS256, &at)
	accessToken, err := atoken.SignedString([]byte(config.LoadedConfig.JWT_SECRET))

	if err != nil {
		return nil, err
	}

	rt := FKClaims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  currentTime.Unix(),
			ExpiresAt: currentTime.Add(time.Hour * 24 * 14).Unix(),
			Subject:   string(req.GetAccountId()),
		},
		ProjectId:   req.GetProjectId(),
		GroupId:     req.GetGroupId(),
		DomainId:    req.GetDomainId(),
		AccountId:   req.GetAccountId(),
		AccountType: req.GetAccountType(),
		TokenType:   "refresh",
	}

	rToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &rt)
	refreshToken, err := rToken.SignedString([]byte(config.LoadedConfig.JWT_SECRET))

	return &iampb.CreateTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
