package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/supperdoggy/diploma_university_statistics_tool/api/pkg/cfg"
	"github.com/supperdoggy/diploma_university_statistics_tool/api/pkg/handlers"
	"github.com/supperdoggy/diploma_university_statistics_tool/api/pkg/service"
	"github.com/supperdoggy/diploma_university_statistics_tool/api/pkg/storage"
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

	srv := service.NewService(db, logger)
	hndls := handlers.NewHandlers(srv, logger)

	r := gin.Default()

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/list_universities", hndls.ListUniversities)

		apiv1.GET("/list_companies", hndls.ListCompanies)
	}

	if err := r.Run(fmt.Sprintf(":%s", config.Port)); err != nil {
		panic(err)
	}

}
