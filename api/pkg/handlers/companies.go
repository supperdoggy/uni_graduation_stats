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

func (h *handlers) ListCompaniesTopSchools(c *gin.Context) {
	var req rest.ListCompaniesTopSchoolsRequest
	var resp rest.ListCompaniesTopSchoolsResponse

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

	schools, err := h.svc.ListCompaniesTopSchools(c.Request.Context(), req.Company)
	if err != nil {
		h.log.Error("failed to get companies", zap.Error(err))
		c.JSON(http.StatusInternalServerError, rest.ListCompaniesTopSchoolsResponse{
			Error: err.Error(),
		})
		return
	}

	resp.Schools = schools
	resp.Count = len(schools)

	c.JSON(http.StatusOK, resp)
}
