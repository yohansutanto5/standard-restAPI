package handler

import (
	"app/constanta"
	"app/db"
	"app/model"
	"app/pkg/log"
	"app/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSystemHealth(c *gin.Context, ds *db.DataStore) {
	// Variable
	var redis bool = true
	var database_primary bool = true
	var database_secondary bool = true

	redis = ds.Redis != nil && func() bool {
		_, err := ds.Redis.Ping().Result()
		return err == nil
	}()

	database_secondary = ds.Db != nil && func() bool {
		sqlDB, err := ds.Db.DB()
		return err == nil && sqlDB.Ping() == nil
	}()

	database_primary = ds.Db != nil && func() bool {
		sqlDB, err := ds.Db.DB()
		return err == nil && sqlDB.Ping() == nil
	}()

	log.Debug(util.GetTransactionID(c), "Debug System", nil)
	result := map[string]bool{
		"redis":              redis,
		"database_primary":   database_primary,
		"database_secondary": database_secondary,
	}

	// c.JSON(http.StatusOK, result)
	errorResponse := model.ErrorResponse{
		TransactionID: util.GetTransactionID(c),
		Message:       constanta.InternalServerErrorMessage,
		Code:          constanta.CodeErrorService,
		Details:       result,
	}
	c.JSON(http.StatusInternalServerError, errorResponse)
}
