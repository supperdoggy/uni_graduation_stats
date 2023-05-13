package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/supperdoggy/diploma_university_statistics_tool/api/pkg/models/rest"
	"go.uber.org/zap"
)

func (h *handlers) ListCompanies(c *gin.Context) {
	companies, err := h.svc.ListCompanies(c.Request.Context())
	if err != nil {
		h.log.Error("failed to get companies", zap.Error(err))
		c.JSON(http.StatusInternalServerError, rest.ListCompaniesResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, rest.ListCompaniesResponse{
		Companies: companies,
		Count:     len(companies),
	})
}
