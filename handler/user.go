package handler

import (
	"app/constanta"
	"app/db"
	"app/model"
	"app/pkg/util"
	"app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	CRUDHandler
	userService service.UserService
}

func NewUserHandler(db *db.DataStore) UserHandler {
	svc := service.NewUserService(db)
	h := UserHandler{
		userService: svc,
	}
	return h
}

func (h *UserHandler) GetList(c *gin.Context) {
	// Call Service
	result, err := h.userService.GetList()

	// Construct Response
	if err != nil {
		c.Errors = append(c.Errors, err.GenerateReponse(util.GetTransactionID(c)))
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

func (h *UserHandler) Insert(c *gin.Context) {
	// Cast data from request
	var data model.AddUserIn
	if errx := c.ShouldBindJSON(&data); errx != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errx.Error()})
		return
	}
	// Construct User Model with the request data
	newUser := &model.User{}
	newUser.PopulateFromDTOInput(data)
	// Call create service
	err := h.userService.Insert(newUser)

	// Construct Response
	if err != nil {
		c.Errors = append(c.Errors, err.GenerateReponse(util.GetTransactionID(c)))
		c.JSON(err.Status, err.Response)
	} else {
		c.JSON(http.StatusCreated, constanta.SuccessMessage)
	}
}
