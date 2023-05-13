package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/supperdoggy/diploma_university_statistics_tool/api/pkg/models/rest"
)

func (h *handlers) ListUniversities(c *gin.Context) {
	var resp rest.ListUniversitiesResponse
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
