package postgres

import (
	"fmt"
	"github.com/firekitz/fk-daemon-iam/config"
	"github.com/golang/protobuf/jsonpb"
	_ "github.com/lib/pq"
	"google.golang.org/protobuf/types/known/structpb"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
	"time"
)

var (
	once      sync.Once
	ProjectDB *gorm.DB
	IamDB     *gorm.DB
)

func DatabaseInit() {
	once.Do(func() {
		ConnectIamDBDatabase()
		ConnectProjectDBDatabase()
	})
}

func ConnectIamDBDatabase() {
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%v sslmode=disable",
		config.LoadedConfig.PG_CORE_IAM_USER,
		config.LoadedConfig.PG_CORE_IAM_PASSWORD,
		config.LoadedConfig.PG_CORE_IAM_DATABASE,
		config.LoadedConfig.PG_CORE_IAM_HOST,
		config.LoadedConfig.PG_CORE_IAM_PORT)
	var err error
	IamDB, err = gorm.Open(postgres.Open(dbInfo), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Unable to connection to database: %v\n", err))
	}
	db, err := IamDB.DB()
	if err != nil {
		panic(fmt.Sprintf("Unable to get database: %v\n", err))
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(1000)
	db.SetConnMaxLifetime(time.Minute)
}

func ConnectProjectDBDatabase() {
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%v sslmode=disable",
		config.LoadedConfig.PG_PROJECT_USER,
		config.LoadedConfig.PG_PROJECT_PASSWORD,
		config.LoadedConfig.PG_PROJECT_DATABASE,
		config.LoadedConfig.PG_PROJECT_HOST,
		config.LoadedConfig.PG_PROJECT_PORT)
	var err error
	ProjectDB, err = gorm.Open(postgres.Open(dbInfo), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Unable to connection to database: %v\n", err))
	}
	db, err := ProjectDB.DB()
	if err != nil {
		panic(fmt.Sprintf("Unable to get database: %v\n", err))
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(1000)
	db.SetConnMaxLifetime(time.Minute)
}

func ToJSON(pb *structpb.Struct) (string, error) {
	marshaller := jsonpb.Marshaler{}
	res, err := marshaller.MarshalToString(pb)
	return res, err
}
