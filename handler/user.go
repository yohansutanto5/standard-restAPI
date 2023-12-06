package handler

import (
	"app/constanta"
	"app/model"
	"app/pkg/util"
	"app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context, User service.UserService) {
	// Call Service
	result, err := User.GetList()

	// Construct Response
	if err != nil {
		err.GenerateReponse(util.GetTransactionID(c))
		c.JSON(err.Status, err.Response)
	} else {
		// Construct DTO out
		var response []model.GetUserOut
		for _, user := range result {
			response = append(response, user.ConstructGetUserOut())
		}
		c.JSON(http.StatusOK, response)
	}
}

func InsertUser(c *gin.Context, User service.UserService) {
	// Cast data from request
	var data model.AddUserIn
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Construct User Model with the request data
	newUser := &model.User{}
	newUser.PopulateFromDTOInput(data)
	// Call create service
	err := User.Insert(newUser)

	// Construct Response
	if err != nil {
		err.GenerateReponse(util.GetTransactionID(c))
		c.JSON(err.Status, err.Response)
	} else {
		c.JSON(http.StatusCreated, constanta.SuccessMessage)
	}
}

func DeleteUser(c *gin.Context, User service.UserService) {
	// To do parsing data here
	id := 1

	// Call service func
	err := User.DeleteByID(id)

	// Construct Response
	if err != nil {
		err.GenerateReponse(util.GetTransactionID(c))
		c.JSON(err.Status, err.Response)
	} else {
		c.JSON(http.StatusNoContent, constanta.SuccessMessage)
	}
}

func UpdateUser(c *gin.Context, User service.UserService) {
	var data model.AddUserIn
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Construct User Model with the request data
	newUser := &model.User{}
	newUser.PopulateFromDTOInput(data)

	// Call Service
	err := User.Update(newUser)

	// Construct Response
	if err != nil {
		err.GenerateReponse(util.GetTransactionID(c))
		c.JSON(err.Status, err.Response)
	} else {
		c.JSON(http.StatusCreated, constanta.SuccessMessage)
	}
}
