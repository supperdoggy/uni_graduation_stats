package service

import (
	"context"

	"github.com/supperdoggy/diploma_university_statistics_tool/api/pkg/models/db"
	"github.com/supperdoggy/diploma_university_statistics_tool/api/pkg/utils"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) CreateUser(ctx context.Context, password, email, fullname string) (*db.User, error) {
	err := utils.ValidateUserEmailAndPassword(email, password)
	if err != nil {
		return nil, err
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		s.log.Error("error hashing password", zap.Error(err))
		return nil, err
	}

	resp, err := s.db.CreateUser(ctx, email, fullname, hashed)
	if err != nil {
		s.log.Error("error CreateUser", zap.Error(err))
		return nil, err
	}

	return resp, nil
}

func (s *service) DeleteUser(ctx context.Context, id string) (*string, error) {
	err := s.db.DeleteUser(ctx, id)
	if err != nil {
		s.log.Error("error DeleteUser", zap.Error(err))
		return nil, err
	}

	return &id, nil
}

func (s *service) UpdateUser(ctx context.Context, id, password, email string) (*db.User, error) {
	err := utils.ValidateUserEmailAndPassword(email, password)
	if err != nil {
		return nil, err
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		s.log.Error("error hashing password", zap.Error(err))
		return nil, err
	}

	err = s.db.UpdateUser(ctx, id, email, hashed)
	if err != nil {
		s.log.Error("error updating user", zap.Any("id", id), zap.Error(err))
		return nil, err
	}

	user, err := s.db.GetUser(ctx, id)
	if err != nil {
		s.log.Error("error getting user", zap.Any("id", id), zap.Error(err))
		return nil, err
	}

	return user, nil
}

func (s *service) GetUser(ctx context.Context, id string) (*db.User, error) {
	user, err := s.db.GetUser(ctx, id)
	if err != nil {
		s.log.Error("error getting user", zap.Any("id", id), zap.Error(err))
		return nil, err
	}

	return user, nil
}
