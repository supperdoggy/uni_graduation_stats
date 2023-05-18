package service

import (
	"context"

	"github.com/supperdoggy/diploma_university_statistics_tool/api/pkg/clients/email"
	"github.com/supperdoggy/diploma_university_statistics_tool/api/pkg/clients/storage"
	"github.com/supperdoggy/diploma_university_statistics_tool/models/db"
	"github.com/supperdoggy/diploma_university_statistics_tool/models/rest"
	"go.uber.org/zap"
)

type IService interface {

	// user
	CreateUser(ctx context.Context, password, email, fullName string) (*db.User, error)
	DeleteUser(ctx context.Context, id string) (*string, error)
	UpdateUser(ctx context.Context, id, password, email string) (*db.User, error)
	GetUser(ctx context.Context, id string) (*db.User, error)

	// auth
	NewToken(ctx context.Context, userID string) (token string, err error)
	CheckToken(ctx context.Context, token string) (userID string, err error)
	Login(ctx context.Context, email, password string) (userID, token string, err error)
	Register(ctx context.Context, email, fullName, password string) (userID, token string, err error)

	// email validation
	NewEmailCode(ctx context.Context, email string) error
	CheckEmailCode(ctx context.Context, email, code string) error

	// Schools
	ListSchools(ctx context.Context) ([]rest.ListSchools, error)
	ListSchoolsTopCompanies(ctx context.Context, school string) ([]rest.ListSchoolsTopCompanies, error)
	ListJobsBySchool(ctx context.Context, school string) ([]rest.ListJobsBySchool, error)
	CorrelationBetweenDegreeAndTitle(ctx context.Context, school string) ([]rest.CorrelationDegreeAndTitle, error)
	DegreesBySchool(ctx context.Context, school string) ([]rest.SchoolDegrees, error)

	// Companies
	ListCompanies(ctx context.Context) ([]rest.ListCompanies, error)
	ListCompaniesTopSchools(ctx context.Context, company string) ([]rest.ListCompaniesTopSchools, error)
	TopHiredDegreesByCompany(ctx context.Context, company, school string) ([]rest.TopHiredDegrees, error)
}

type service struct {
	// storage interface
	db storage.IMongoDB

	// logger
	log *zap.Logger

	emailClient email.IEmailClient
}

func NewService(db storage.IMongoDB, log *zap.Logger, e email.IEmailClient) IService {
	return &service{
		db:  db,
		log: log,
	}
}
