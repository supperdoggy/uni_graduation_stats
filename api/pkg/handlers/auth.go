package handlers

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/supperdoggy/diploma_university_statistics_tool/models/rest"
	"go.uber.org/zap"
)

func (h *handlers) Login(c *gin.Context) {
	var (
		req  rest.LoginReq
		resp rest.LoginResp
		err  error
	)

	if err := c.Bind(&req); err != nil {
		h.log.Error("error Bing", zap.Error(err))
		resp.Error = "error reading request"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.UserID, resp.Token, err = h.svc.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		h.log.Info("error Login", zap.Error(err))
		resp.Error = "Login error"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *handlers) CheckToken(c *gin.Context) {
	var (
		req  rest.CheckTokenReq
		resp rest.CheckTokenResp
		err  error
	)

	if err := c.Bind(&req); err != nil {
		h.log.Error("error Bing", zap.Error(err))
		resp.Error = "error reading request"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.UserID, err = h.svc.CheckToken(c.Request.Context(), req.Token)
	if err != nil {
		h.log.Info("error Login", zap.Error(err))
		resp.Error = "CheckToken error"
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	resp.OK = true

	c.JSON(http.StatusOK, resp)
}

func (h *handlers) Middleware(c *gin.Context) {
	if !h.authEnabled {
		c.Next()
		return
	}

	// read from the header the token
	token := c.GetHeader("Authorization")
	if token == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// check the token
	userID, err := h.svc.CheckToken(c.Request.Context(), token)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// set the userID to the context
	c.Set("userID", userID)
	c.Next()
}

func (h *handlers) Register(c *gin.Context) {
	var (
		req  rest.RegisterReq
		resp rest.RegisterResp
		err  error
	)

	if err := c.Bind(&req); err != nil {
		h.log.Error("error Bing", zap.Error(err))
		resp.Error = "error reading request"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.FirstName, validation.Required, validation.
			Match(regexp.MustCompile("^[a-zA-ZА-Яа-я]+$"))),
		validation.Field(&req.LastName, validation.Required, validation.
			Match(regexp.MustCompile("^[a-zA-ZА-Яа-я]+$"))),
		validation.Field(&req.Email, validation.Required, is.Email),
		validation.Field(&req.Password, validation.Required, validation.Length(6, 100)),
	); err != nil {
		h.log.Error("error Bing", zap.Error(err))
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	fullname := fmt.Sprintf("%s %s", req.FirstName, req.LastName)

	resp.UserID, resp.Token, err = h.svc.Register(c.Request.Context(), req.Email, fullname, req.Password)
	if err != nil {
		h.log.Error("error Bing", zap.Error(err))
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *handlers) NewEmailCode(c *gin.Context) {
	var (
		req  rest.NewEmailCodeReq
		resp rest.NewEmailCodeResp
		err  error
	)

	if err := c.Bind(&req); err != nil {
		h.log.Error("error Bing", zap.Error(err))
		resp.Error = "error reading request"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Email, validation.Required, is.Email),
	); err != nil {
		h.log.Error("error Bing", zap.Error(err))
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	err = h.svc.NewEmailCode(c.Request.Context(), req.Email)
	if err != nil {
		h.log.Error("error Bing", zap.Error(err))
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.OK = true

	c.JSON(http.StatusOK, resp)
}

func (h *handlers) CheckEmailCode(c *gin.Context) {
	var (
		req  rest.CheckEmailCodeReq
		resp rest.CheckEmailCodeResp
		err  error
	)

	if err := c.Bind(&req); err != nil {
		h.log.Error("error Bing", zap.Error(err))
		resp.Error = "error reading request"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Email, validation.Required, is.Email),
		validation.Field(&req.Code, validation.Required),
	); err != nil {
		h.log.Error("error Bing", zap.Error(err))
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	err = h.svc.CheckEmailCode(c.Request.Context(), req.Email, req.Code)
	if err != nil {
		h.log.Error("error Bing", zap.Error(err))
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.OK = true

	c.JSON(http.StatusOK, resp)
}
