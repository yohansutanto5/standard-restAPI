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
	UserService service.UserService
}

func NewUserHandler(db *db.DataStore) UserHandler {
	svc := service.NewUserService(db)
	h := UserHandler{
		UserService: svc,
	}
	return h
}



func (h *UserHandler) GetList(c *gin.Context) {
	// Call Service
	result, err := h.UserService.GetList()

	// Construct Response
	if err != nil {
		c.Errors = append(c.Errors, err.GenerateReponse(util.GetTransactionID(c)))
		c.JSON(err.Status, err.Response)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func (h *UserHandler) Insert(c *gin.Context) {
	// Cast data from request
	var data model.UserInput
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Construct User Model with the request data
	newUser := &model.User{}
	newUser.PopulateFromDTOInput(data)
	// Call create service
	err := h.UserService.Insert(newUser)

	// Construct Response
	if err != nil {
		err.GenerateReponse(util.GetTransactionID(c))
		c.JSON(err.Status, err.Response)
	} else {
		c.JSON(http.StatusCreated, constanta.SuccessMessage)
	}
}