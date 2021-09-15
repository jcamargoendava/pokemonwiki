package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	masterModel "github.com/jcamargoendava/pokemonwiki/models"
	"github.com/jcamargoendava/pokemonwiki/repository"
	"github.com/jcamargoendava/pokemonwiki/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetMaster(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "param id is required"})
		return
	}
	ctx, _ := c.MustGet("ctx").(context.Context)
	masterRepo := repository.NewMaster("master")
	masterService := services.NewMaster(masterRepo)
	foundMaster, err := masterService.GetMaster(ctx, id)
	if err != nil {
		fmt.Errorf("Error trying to get a master")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error ocurred"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": foundMaster})
}

func CreateMaster(c *gin.Context) {
	var master masterModel.Master
	if err := c.ShouldBindJSON(&master); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx, _ := c.MustGet("ctx").(context.Context)
	masterRepo := repository.NewMaster("master")
	masterService := services.NewMaster(masterRepo)
	master.ID = primitive.NewObjectID()
	createdMaster, err := masterService.SaveMaster(ctx, &master)
	if err != nil {
		fmt.Errorf("Error trying to create a master")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": createdMaster})
}

func UpdateMaster(c *gin.Context) {
	var master masterModel.Master
	if err := c.ShouldBindJSON(&master); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "param id is required"})
		return
	}
	ctx, _ := c.MustGet("ctx").(context.Context)
	masterRepo := repository.NewMaster("master")
	masterService := services.NewMaster(masterRepo)
	foundMaster, err := masterService.UpdateMaster(ctx, id, &master)
	if err != nil {
		fmt.Errorf("Error trying to get a master")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error ocurred"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": foundMaster})
}

func DeleteMaster(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "param id is required"})
		return
	}
	ctx, _ := c.MustGet("ctx").(context.Context)
	masterRepo := repository.NewMaster("master")
	masterService := services.NewMaster(masterRepo)
	if err := masterService.DeleteMaster(ctx, id); err != nil {
		fmt.Errorf("Error trying to get a master")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "An error ocurred"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "deleted"})
}
