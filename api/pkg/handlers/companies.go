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

func (h *handlers) ListCompaniesTopUniversities(c *gin.Context) {
	var req rest.ListCompaniesTopUniversitiesRequest
	var resp rest.ListCompaniesTopUniversitiesResponse

	if err := c.Bind(&req); err != nil {
		h.log.Error("failed to bind request", zap.Error(err))
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	if req.Company == "" {
		h.log.Error("company is empty")
		resp.Error = "company is empty"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	schools, err := h.svc.ListCompaniesTopUniversities(c.Request.Context(), req.Company)
	if err != nil {
		h.log.Error("failed to get companies", zap.Error(err))
		c.JSON(http.StatusInternalServerError, rest.ListCompaniesTopUniversitiesResponse{
			Error: err.Error(),
		})
		return
	}

	resp.Universities = schools
	resp.Count = len(schools)

	c.JSON(http.StatusOK, resp)
}
