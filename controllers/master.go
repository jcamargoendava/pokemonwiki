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

// GetMaster godoc
// @Description Gets a Master using the id from the path parameter
// @Accept  json
// @Produce  json
// @Header 200 {string} Token "qwerty"
// @Success 200 {object} masterModel.Master
// @Failure 404 "not found"
// @Param id path string true "Master ID"
// @Failure 500 "Internal Server Error"
// @Tags master
// @Router /master/{id} [get]
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

// CreateMaster godoc
// @Description Create a new Master based on the information in the body of the post request
// @Accept  json
// @Produce  json
// @Header 200 {string} Token "qwerty"
// @Success 200 {object} masterModel.Master
// @Failure 404 "not found"
// @Failure 500 "Internal Server Error"
// @Param Body body dto.master_body true "Request Body"
// @Tags master
// @Router /master [post]
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

// UpdateMaster godoc
// @Description Updates a Master based on the id found in the path parameter with the information found in the body of the put request
// @Accept  json
// @Produce  json
// @Header 200 {string} Token "qwerty"
// @Success 200 {object} masterModel.Master
// @Failure 404 "not found"
// @Failure 500 "Internal Server Error"
// @Param id path string true "ID"
// @Param Body body dto.master_body true "Request Body"
// @Tags master
// @Router /master/{id} [put]
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

// DeleteMaster godoc
// @Description Deletes a Master based on the id it finds in the path parameter
// @Accept  json
// @Produce  json
// @Header 200 {string} Token "qwerty"
// @Success 200 {object} masterModel.Master
// @Failure 400 "not found"
// @Failure 404 "not found"
// @Failure 500 "Internal Server Error"
// @Param id path string true "ID"
// @Tags master
// @Router /master/{id} [delete]
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

// GetMasters godoc
// @Description Returns a list of all the Masters inside the mongoDB database
// @Accept  json
// @Produce  json
// @Header 200 {string} Token "qwerty"
// @Success 200 {object} dto.Response_200_get_masters
// @Failure 404 "not found"
// @Failure 500 "Internal Server Error"
// @Tags master
// @Router /masters/ [get]
// func GetMasters(c *gin.Context) {
// 	ctx, _ := c.MustGet("ctx").(context.Context)
// 	limit := c.Query("limit")
// 	if len(limit) == 0 {
// 		limit = "100"
// 	}
// 	offset := c.Query("offset")
// 	if len(offset) == 0 {
// 		offset = "0"
// 	}

// 	offsetInt, _ := strconv.Atoi(offset)
// 	limitInt, _ := strconv.Atoi(limit)
// 	masterRepo := repository.NewMaster("master")
// 	masterService := services.NewMaster(masterRepo)
// 	mastersFound := masterService.RetrieveAllMasters(ctx, offsetInt, limitInt)
// 	c.JSON(http.StatusOK, gin.H{"data": mastersFound})
// }
