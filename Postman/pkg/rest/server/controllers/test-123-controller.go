package controllers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/unnagirirao/Automation_Testing/postman/pkg/rest/server/models"
	"github.com/unnagirirao/Automation_Testing/postman/pkg/rest/server/services"
	"net/http"
	"strconv"
)

type Test123Controller struct {
	test123Service *services.Test123Service
}

func NewTest123Controller() (*Test123Controller, error) {
	test123Service, err := services.NewTest123Service()
	if err != nil {
		return nil, err
	}
	return &Test123Controller{
		test123Service: test123Service,
	}, nil
}

func (test123Controller *Test123Controller) CreateTest123(context *gin.Context) {
	// validate input
	var input models.Test123
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger test123 creation
	if _, err := test123Controller.test123Service.CreateTest123(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Test123 created successfully"})
}

func (test123Controller *Test123Controller) UpdateTest123(context *gin.Context) {
	// validate input
	var input models.Test123
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger test123 update
	if _, err := test123Controller.test123Service.UpdateTest123(id, &input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Test123 updated successfully"})
}

func (test123Controller *Test123Controller) FetchTest123(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger test123 fetching
	test123, err := test123Controller.test123Service.GetTest123(id)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, test123)
}

func (test123Controller *Test123Controller) DeleteTest123(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger test123 deletion
	if err := test123Controller.test123Service.DeleteTest123(id); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Test123 deleted successfully",
	})
}

func (test123Controller *Test123Controller) ListTest123s(context *gin.Context) {
	// trigger all test123s fetching
	test123s, err := test123Controller.test123Service.ListTest123s()
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, test123s)
}

func (*Test123Controller) PatchTest123(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PATCH",
	})
}

func (*Test123Controller) OptionsTest123(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "OPTIONS",
	})
}

func (*Test123Controller) HeadTest123(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "HEAD",
	})
}
