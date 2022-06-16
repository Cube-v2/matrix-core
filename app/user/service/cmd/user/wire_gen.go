// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/the-zion/matrix-core/app/user/service/internal/biz"
	"github.com/the-zion/matrix-core/app/user/service/internal/conf"
	"github.com/the-zion/matrix-core/app/user/service/internal/data"
	"github.com/the-zion/matrix-core/app/user/service/internal/server"
	"github.com/the-zion/matrix-core/app/user/service/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, auth *conf.Auth, logger log.Logger, registry *nacos.Registry) (*kratos.App, func(), error) {
	db := data.NewDB(confData, logger)
	cmdable := data.NewRedis(confData, logger)
	producer := data.NewRocketmqProducer(confData, logger)
	txCode := data.NewPhoneCode(confData)
	goMail := data.NewGoMail(confData)
	cos := data.NewCosClient(confData)
	dataData, cleanup, err := data.NewData(db, cmdable, producer, txCode, goMail, cos, logger)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	userUseCase := biz.NewUserUseCase(userRepo, logger)
	authRepo := data.NewAuthRepo(dataData, logger)
	transaction := data.NewTransaction(dataData)
	authUseCase := biz.NewAuthUseCase(auth, authRepo, userRepo, transaction, logger)
	profileRepo := data.NewProfileRepo(dataData, logger)
	profileUseCase := biz.NewProfileUseCase(profileRepo, logger)
	userService := service.NewUserService(userUseCase, authUseCase, profileUseCase, logger)
	httpServer := server.NewHTTPServer(confServer, userService, logger)
	grpcServer := server.NewGRPCServer(confServer, userService, logger)
	mqConsumerServer := server.NewMqConsumerServer(confData, userService, logger)
	app := newApp(logger, registry, httpServer, grpcServer, mqConsumerServer)
	return app, func() {
		cleanup()
	}, nil
}
