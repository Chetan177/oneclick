package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/google/uuid"

	"github.com/chetan177/oneclick/models"

)

func (s *Server) createProjectHandler(c echo.Context) error {
	project := &models.ProjectRequest{}
	if err := c.Bind(project); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
	}

	if err := s.ValidateRequest(c, project); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
	}

	existingProject, err := s.DB.GetProjectByName(project.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
	}
	if existingProject != nil {
		return c.JSON(http.StatusConflict, models.ErrorResponse{
			Error: "Project with the same name already exists",
		})
	}


	uuid := uuid.New()	
	projectDb := &models.Project{
		ID: uuid.String(),
		Name: project.Name,
		Description: project.Description,
	}
	err = s.DB.CreateProject(projectDb)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, projectDb)
}

func (s *Server) getProjectsHandler(c echo.Context) error {
	projects, err := s.DB.GetProjects()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, projects)
}

func (s *Server) getProjectHandler(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Project ID is required",
		})
	}
	project, err := s.DB.GetProject(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, project)
}

func (s *Server) updateProjectHandler(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Project ID is required",
		})
	}
	project := &models.ProjectRequest{}
	if err := c.Bind(project); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
	}

	if err := s.ValidateRequest(c, project); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
	}

	projectDb, err := s.DB.GetProject(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
	}
	projectDb.Name = project.Name
	projectDb.Description = project.Description
	err = s.DB.UpdateProject(id, projectDb)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, projectDb)
}

func (s *Server) deleteProjectHandler(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Project ID is required",
		})
	}
	err := s.DB.DeleteProject(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Project deleted successfully"})
}

