package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/supperdoggy/diploma_university_statistics_tool/api/pkg/utils"
	"github.com/supperdoggy/diploma_university_statistics_tool/models/rest"
	"go.uber.org/zap"
)

func (h *handlers) CreateUser(c *gin.Context) {
	var (
		req  rest.CreateUserRequest
		resp rest.CreateUserResponse
		ctx  context.Context
	)
	if err := c.Bind(&req); err != nil {
		h.log.Error("error Bing", zap.Error(err))
		resp.Error = "error reading request"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	user, err := h.svc.CreateUser(ctx, req.Password, req.Email, req.FullName)
	if err != nil {
		h.log.Error("error CreateUser", zap.Error(err))
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.ID = user.ID
	c.JSON(http.StatusOK, resp)
}

func (h *handlers) DeleteUser(c *gin.Context) {
	var (
		req  rest.DeleteUserRequest
		resp rest.DeleteUserResponse
		ctx  context.Context
		err  error
	)
	if err := c.Bind(&req); err != nil {
		h.log.Error("error Bing", zap.Error(err))
		resp.Error = "error reading request"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	h.log.Info("DeleteUser", zap.Any("req", req))

	id, err := h.svc.DeleteUser(ctx, req.ID)
	if err != nil {
		h.log.Error("error deleting user", zap.Error(err), zap.Any("req", req))
		resp.Error = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp.ID = id
	c.JSON(http.StatusOK, resp)
}

func (h *handlers) UpdateUser(c *gin.Context) {
	var (
		req  rest.UpdateUserRequest
		resp rest.UpdateUserResponse
		ctx  context.Context
		err  error
	)
	if err := c.Bind(&req); err != nil {
		h.log.Error("error Bing", zap.Error(err))
		resp.Error = "error reading request"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	user, err := h.svc.UpdateUser(ctx, req.ID, req.Password, req.Email)
	if err != nil {
		h.log.Error("error UpdateUser", zap.Error(err), zap.Any("id", req.ID))
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.User = utils.MapDBUserToResponseUser(*user)
	c.JSON(http.StatusBadRequest, resp)
}

func (h *handlers) GetUser(c *gin.Context) {
	var (
		resp rest.GetUserResponse
		ctx  context.Context
		err  error
	)
	userID := c.Param("id")

	user, err := h.svc.GetUser(ctx, userID)
	if err != nil {
		h.log.Error("error GetUser", zap.Error(err), zap.Any("id", userID))
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.User = utils.MapDBUserToResponseUser(*user)
	c.JSON(http.StatusOK, resp)
}
