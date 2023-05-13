package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/supperdoggy/diploma_university_statistics_tool/api/pkg/service"
	"go.uber.org/zap"
)

type IHandlers interface {
	// Universities
	ListUniversities(c *gin.Context)

	// Companies
	ListCompanies(c *gin.Context)
	ListCompaniesTopUniversities(c *gin.Context)
}

type handlers struct {
	// service interface
	svc service.IService

	// logger
	log *zap.Logger
}

func NewHandlers(svc service.IService, log *zap.Logger) IHandlers {
	return &handlers{
		svc: svc,
		log: log,
	}
}
