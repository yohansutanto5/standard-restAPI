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

type UserProfileHandler struct {
	CRUDHandler
	userProfileService service.UserProfileService
}

func NewUserProfileHandler(db *db.DataStore) UserProfileHandler {
	svc := service.NewUserProfileService(db)
	h := UserProfileHandler{
		userProfileService: svc,
	}
	return h
}

func (h *UserProfileHandler) GetList(c *gin.Context) {
	// Call Service
	result, err := h.userProfileService.GetList()

	// Construct Response
	if err != nil {
		err.GenerateReponse(util.GetTransactionID(c))
		c.JSON(err.Status, err.Response)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func (h *UserProfileHandler) Insert(c *gin.Context) {
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
	err := h.userProfileService.Insert(newUserProfile)

	// Construct Response
	if err != nil {
		err.GenerateReponse(util.GetTransactionID(c))
		c.JSON(err.Status, err.Response)
	} else {
		c.JSON(http.StatusCreated, constanta.SuccessMessage)
	}
}
