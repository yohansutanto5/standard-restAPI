package handler

import (
	"app/constanta"
	"app/model"
	"app/pkg/util"
	"app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserProfile(c *gin.Context, UserProfile service.UserProfileService) {
	// Call Service
	result, err := UserProfile.GetList()

	// Construct Response
	if err != nil {
		err.GenerateReponse(util.GetTransactionID(c))
		c.JSON(err.Status, err.Response)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func InsertUserProfile(c *gin.Context, UserProfile service.UserProfileService) {
	// Cast data from request
	var data model.AddUserProfileIn
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Construct UserProfile Model with the request data
	newUserProfile := &model.UserProfile{}
	newUserProfile.PopulateFromDTOInput(data)
	// Call create service
	err := UserProfile.Insert(newUserProfile)

	// Construct Response
	if err != nil {
		err.GenerateReponse(util.GetTransactionID(c))
		c.JSON(err.Status, err.Response)
	} else {
		c.JSON(http.StatusCreated, constanta.SuccessMessage)
	}
}

func DeleteUserProfile(c *gin.Context, UserProfile service.UserProfileService) {
	// To do parsing data here
	id := 1

	// Call service func
	err := UserProfile.DeleteByID(id)

	// Construct Response
	if err != nil {
		err.GenerateReponse(util.GetTransactionID(c))
		c.JSON(err.Status, err.Response)
	} else {
		c.JSON(http.StatusNoContent, constanta.SuccessMessage)
	}
}

func UpdateUserProfile(c *gin.Context, UserProfile service.UserProfileService) {
	var data model.AddUserProfileIn
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Construct UserProfile Model with the request data
	newUserProfile := &model.UserProfile{}
	newUserProfile.PopulateFromDTOInput(data)

	// Call Service
	err := UserProfile.Update(newUserProfile)

	// Construct Response
	if err != nil {
		err.GenerateReponse(util.GetTransactionID(c))
		c.JSON(err.Status, err.Response)
	} else {
		c.JSON(http.StatusCreated, constanta.SuccessMessage)
	}
}
