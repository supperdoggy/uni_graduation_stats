package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/supperdoggy/diploma_university_statistics_tool/api/pkg/models/rest"
)

func (h *handlers) ListSchools(c *gin.Context) {
	var resp rest.ListSchoolsResponse
	schools, err := h.svc.ListSchools(c.Request.Context())
	if err != nil {
		resp.Error = err.Error()
		c.JSON(500, resp)
		return
	}

	resp.Schools = schools
	resp.Count = len(schools)
	c.JSON(200, resp)
}

func (h *handlers) ListSchoolsTopCompanies(c *gin.Context) {
	var req rest.ListSchoolsTopCompaniesRequest
	var resp rest.ListSchoolsTopCompaniesResponse

	if err := c.Bind(&req); err != nil {
		resp.Error = err.Error()
		c.JSON(400, resp)
		return
	}

	if req.School == "" {
		resp.Error = "school is empty"
		c.JSON(400, resp)
		return
	}

	companies, err := h.svc.ListSchoolsTopCompanies(c.Request.Context(), req.School)
	if err != nil {
		resp.Error = err.Error()
		c.JSON(500, resp)
		return
	}

	resp.Companies = companies
	resp.Count = len(companies)
	c.JSON(200, resp)
}

func (h *handlers) ListJobTitlesBySchool(c *gin.Context) {
	var req rest.ListJobsBySchoolRequest
	var resp rest.ListJobsBySchoolResponse

	if err := c.Bind(&req); err != nil {
		resp.Error = err.Error()
		c.JSON(400, resp)
		return
	}

	if req.School == "" {
		resp.Error = "school is empty"
		c.JSON(400, resp)
		return
	}

	jobs, err := h.svc.ListJobsBySchool(c.Request.Context(), req.School)
	if err != nil {
		resp.Error = err.Error()
		c.JSON(500, resp)
		return
	}

	resp.Jobs = jobs
	resp.Count = len(jobs)
	c.JSON(200, resp)
}

func (h *handlers) CorrelationBetweenDegreeAndTitle(c *gin.Context) {
	var req rest.CorrelationDegreeAndTitleRequest
	var resp rest.CorrelationDegreeAndTitleResponse

	if err := c.Bind(&req); err != nil {
		resp.Error = err.Error()
		c.JSON(400, resp)
		return
	}

	if req.School == "" {
		resp.Error = "school is empty"
		c.JSON(400, resp)
		return
	}

	correlation, err := h.svc.CorrelationBetweenDegreeAndTitle(c.Request.Context(), req.School)
	if err != nil {
		resp.Error = err.Error()
		c.JSON(500, resp)
		return
	}

	resp.Correlations = correlation
	resp.Count = len(correlation)

	c.JSON(200, resp)
}
