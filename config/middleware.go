package config

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	logrusCustomFunc = func(code codes.Code) logrus.Level {
		if code == codes.OK {
			return logrus.InfoLevel
		}
		return logrus.ErrorLevel
	}
	recoveryCustomFunc grpc_recovery.RecoveryHandlerFunc
)

func Logger() (*logrus.Entry, []grpc_logrus.Option) {
	// Logrus entry is used, allowing pre-definition of certain fields by the user.
	logrus.ErrorKey = "grpc.error"
	logrusEntry := logrus.NewEntry(logrus.StandardLogger())
	// Shared options for the logger, with a custom gRPC code to log level function.
	opts := []grpc_logrus.Option{
		grpc_logrus.WithLevels(logrusCustomFunc),
	}
	// Make sure that log statements internal to gRPC library are logged using the logrus Logger as well.
	grpc_logrus.ReplaceGrpcLogger(logrusEntry)

	return logrusEntry, opts
}

func Recovery() []grpc_recovery.Option {
	// Define customfunc to handle panic
	recoveryCustomFunc = func(p interface{}) (err error) {
		return status.Errorf(codes.Unknown, "panic triggeerd: %v", p)
	}
	// Shared options for the logger, with a custom gRPC code to log level function.
	opts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(recoveryCustomFunc),
	}
	return opts
}

func Middleware() *grpc.Server {
	logrusEntry, loggerOpts := Logger()
	recoveryOpts := Recovery()
	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_recovery.UnaryServerInterceptor(recoveryOpts...),
			grpc_logrus.UnaryServerInterceptor(logrusEntry, loggerOpts...),
			grpc_validator.UnaryServerInterceptor(),
		)),
	)
	return server
}