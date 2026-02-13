package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mod/models"
	"go.mod/services"
)

type JobTitleController struct {
	service services.JobTitleService
}

func NewJobTitleController(s services.JobTitleService) *JobTitleController {
	return &JobTitleController{service: s}
}

// POST /job-titles

// CreateJobTitle godoc
// @Summary Create a new job title
// @Description Create a new job title in the system
// @Tags JobTitles
// @Accept json
// @Produce json
// @Param job body models.JobTitle true "Job Title"
// @Success 201 {object} models.JobTitle
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /job-titles [post]
func (c *JobTitleController) Create(ctx *gin.Context) {
	var job models.JobTitle

	if err := ctx.ShouldBindJSON(&job); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.Create(ctx, &job); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, job)
}

// GET /job-titles/:id
// GetJobTitle godoc
// @Summary Get a job title by ID
// @Description Get a specific job title by ID
// @Tags JobTitles
// @Produce json
// @Param id path int true "Job Title ID"
// @Success 200 {object} models.JobTitle
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /job-titles/{id} [get]
func (c *JobTitleController) Get(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	job, err := c.service.Get(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "job title not found"})
		return
	}

	ctx.JSON(http.StatusOK, job)
}

// GET /job-titles?page=1&page_size=10
// ListJobTitles godoc
// @Summary List job titles
// @Description Get a paginated list of job titles
// @Tags JobTitles
// @Produce json
// @Param page query int false "Page number"
// @Param page_size query int false "Number of items per page"
// @Success 200 {object} dto.JobTitleListResponse
// @Failure 500 {object} map[string]string
// @Router /job-titles [get]
func (c *JobTitleController) List(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	result, err := c.service.List(ctx, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Returning generic response at runtime
	ctx.JSON(http.StatusOK, result)
}

// PUT /job-titles/:id
// UpdateJobTitle godoc
// @Summary Update a job title
// @Description Update an existing job title by ID
// @Tags JobTitles
// @Accept json
// @Produce json
// @Param id path int true "Job Title ID"
// @Param job body models.JobTitle true "Job Title"
// @Success 200 {object} models.JobTitle
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /job-titles/{id} [put]
func (c *JobTitleController) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var job models.JobTitle
	if err := ctx.ShouldBindJSON(&job); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	job.ID = id

	if err := c.service.Update(ctx, &job); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, job)
}

// DELETE /job-titles/:id
// DeleteJobTitle godoc
// @Summary Delete a job title
// @Description Delete a job title by ID
// @Tags JobTitles
// @Param id path int true "Job Title ID"
// @Success 200
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /job-titles/{id} [delete]

func (c *JobTitleController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := c.service.Delete(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
