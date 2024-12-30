package transport

import (
	"github.com/poolcamacho/jobs-service/internal/domain"
	"github.com/poolcamacho/jobs-service/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// JobHandler handles HTTP requests related to jobs
// This struct acts as the controller for handling HTTP endpoints related to jobs.
type JobHandler struct {
	service service.JobService // Dependency on JobService for business logic
}

// NewJobHandler creates a new JobHandler instance
// This is a constructor function to initialize the JobHandler with a JobService dependency.
func NewJobHandler(service service.JobService) *JobHandler {
	return &JobHandler{service: service}
}

// HealthCheck provides a simple health status of the service
// @Summary Check service health
// @Description Returns the health status of the jobs service
// @Tags Health
// @Produce json
// @Success 200 {object} map[string]string "Service is healthy"
// @Router /health [get]
func (h *JobHandler) HealthCheck(c *gin.Context) {
	// Respond with a simple JSON indicating the service is running
	c.JSON(http.StatusOK, gin.H{"status": "healthy"})
}

// GetJobs handles the retrieval of all jobs
// @Summary Get all jobs
// @Description Retrieve a list of all jobs in the system
// @Tags Jobs
// @Produce json
// @Success 200 {array} domain.Job "List of jobs"
// @Failure 500 {object} map[string]string "Failed to fetch jobs"
// @Router /jobs [get]
func (h *JobHandler) GetJobs(c *gin.Context) {
	// Fetch all jobs using the service
	jobs, err := h.service.GetAllJobs()
	if err != nil {
		// Return 500 Internal Server Error if fetching fails
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch jobs"})
		return
	}
	// Return the list of jobs in JSON format
	c.JSON(http.StatusOK, jobs)
}

// CreateJob handles the creation of a new job
// @Summary Create a new job
// @Description Add a new job by providing title, description, and salary range
// @Tags Jobs
// @Accept json
// @Produce json
// @Param request body domain.Job true "Job Creation Request"
// @Success 201 {object} map[string]string "Job created successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Failed to create job"
// @Router /jobs [post]
func (h *JobHandler) CreateJob(c *gin.Context) {
	var job domain.Job
	// Bind the incoming JSON request to the Job struct
	if err := c.ShouldBindJSON(&job); err != nil {
		// Return 400 Bad Request if JSON is invalid
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate required fields
	if job.Description == "" || job.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title and description are required"})
		return
	}

	// Add the job using the service
	if err := h.service.AddJob(&job); err != nil {
		// Return 500 Internal Server Error if creation fails
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create job"})
		return
	}

	// Return 201 Created status with a success message
	c.JSON(http.StatusCreated, gin.H{"message": "job created successfully"})
}
