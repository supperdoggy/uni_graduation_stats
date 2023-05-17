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
	hndls := handlers.NewHandlers(srv, logger, config.AuthorizationEnabled)

	r := gin.Default()

	// auth
	auth := r.Group("/auth")
	{
		auth.POST("/register", hndls.Register)
		auth.POST("/login", hndls.Login)
		auth.POST("/check_token", hndls.CheckToken)
		auth.POST("/new_email_code", hndls.NewEmailCode)
		auth.POST("/check_email_code", hndls.CheckEmailCode)
	}

	// api
	apiv1 := r.Group("/api/v1")

	users := apiv1.Group("/users")
	users.Use(hndls.Middleware)
	{
		users.POST("/create", hndls.CreateUser)
		users.DELETE("/delete", hndls.DeleteUser)
		users.PATCH("/update", hndls.UpdateUser)
		users.GET("/get/:id", hndls.GetUser)
	}

	// Schools
	schools := apiv1.Group("/schools")
	schools.Use(hndls.Middleware)
	{
		schools.GET("/list", hndls.ListSchools)
		schools.POST("/top_companies", hndls.ListSchoolsTopCompanies)
	}

	// Companies
	companies := apiv1.Group("/companies")
	companies.Use(hndls.Middleware)
	{
		companies.GET("/list", hndls.ListCompanies)
		companies.POST("/top_schools", hndls.ListCompaniesTopSchools)
	}

	if err := r.Run(fmt.Sprintf(":%s", config.Port)); err != nil {
		panic(err)
	}

}
