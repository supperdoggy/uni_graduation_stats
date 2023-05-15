package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/supperdoggy/diploma_university_statistics_tool/api/pkg/cfg"
	"github.com/supperdoggy/diploma_university_statistics_tool/api/pkg/clients/email"
	"github.com/supperdoggy/diploma_university_statistics_tool/api/pkg/clients/storage"
	"github.com/supperdoggy/diploma_university_statistics_tool/api/pkg/handlers"
	"github.com/supperdoggy/diploma_university_statistics_tool/api/pkg/service"
	"go.uber.org/zap"
)

func main() {

	// initializing pkg instances
	ctx := context.Background()
	config := cfg.NewConfig()

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	db, err := storage.NewMongoDB(ctx, config.MongoDB, logger)
	if err != nil {
		logger.Fatal("error initializing mongodb", zap.Error(err))
	}

	emailClient := email.NewClient(*logger, config.EmailCheckService)
	srv := service.NewService(db, logger, emailClient)
	hndls := handlers.NewHandlers(srv, logger)

	r := gin.Default()

	// auth
	auth := r.Group("/auth")
	auth.POST("/register", hndls.Register)
	auth.POST("/login", hndls.Login)
	auth.POST("/check_token", hndls.CheckToken)
	auth.POST("/new_email_code", hndls.NewEmailCode)
	auth.POST("/check_email_code", hndls.CheckEmailCode)

	// users
	apiUser := r.Group("/user")
	apiUser.Use(hndls.Middleware)
	{
		apiUser.POST("/create", hndls.CreateUser)
		apiUser.DELETE("/delete", hndls.DeleteUser)
		apiUser.PATCH("/update", hndls.UpdateUser)
		apiUser.GET("/get/:id", hndls.GetUser)
	}

	apiv1 := r.Group("/api/v1")
	apiv1.Use(hndls.Middleware)
	{
		// Universities
		apiv1.GET("/list_universities", hndls.ListUniversities)

		// Companies
		apiv1.GET("/list_companies", hndls.ListCompanies)
		apiv1.POST("/list_companies_top_universities", hndls.ListCompaniesTopUniversities)
	}

	if err := r.Run(fmt.Sprintf(":%s", config.Port)); err != nil {
		panic(err)
	}

}
