package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/supperdoggy/diploma_university_statistics_tool/api/pkg/service"
	"go.uber.org/zap"
)

type IHandlers interface {
	// users
	CreateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	GetUser(c *gin.Context)

	// auth
	Login(c *gin.Context)
	Register(c *gin.Context)
	CheckToken(c *gin.Context)
	Middleware(c *gin.Context)

	// email validation
	NewEmailCode(c *gin.Context)
	CheckEmailCode(c *gin.Context)

	// Schools
	ListSchools(c *gin.Context)
	ListSchoolsTopCompanies(c *gin.Context)
	ListJobTitlesBySchool(c *gin.Context)
	CorrelationBetweenDegreeAndTitle(c *gin.Context)
	SchoolDegrees(c *gin.Context)

	// Companies
	ListCompanies(c *gin.Context)
	ListCompaniesTopSchools(c *gin.Context)
	TopHiredDegrees(c *gin.Context)
}

type handlers struct {
	// service interface
	svc service.IService

	// auth
	authEnabled bool

	// logger
	log *zap.Logger
}

func NewHandlers(svc service.IService, log *zap.Logger, authEnabled bool) IHandlers {
	return &handlers{
		svc:         svc,
		log:         log,
		authEnabled: authEnabled,
	}
}
