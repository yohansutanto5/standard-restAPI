package system_test

import (
	"app/handler"
	"app/model"
	"app/pkg/error"
	"app/pkg/log"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSystemHealth(t *testing.T) {

	h := handler.NewSystemHandler(database)
	h.GetSystemHealth(ctx)
	assert.Equal(t, ctx.Writer.Status(), 200)
}

func TestGinError(t *testing.T) {
	transactionID := 12312312
	err := error.Error{
		Message:  "21",
		Code:     "ggg",
		Status:   88,
		Response: error.ErrorResponse{},
	}
	aa := err.GenerateReponse(transactionID) // gin error
	assert.NotNil(t, aa.Meta)
	ctx.Errors = append(ctx.Errors, aa)

	responseLog := model.CustomLog{
		Agent:         ctx.Request.UserAgent(),
		Method:        ctx.Request.Method,
		ClientIp:      ctx.ClientIP(),
		Path:          ctx.Request.URL.Path,
		TransactionID: transactionID,
		Status:        ctx.Writer.Status(),
	}
	assert.Equal(t, 1, len(ctx.Errors))
	v := ctx.Errors.Last() //gin errr
	assert.NotNil(t, v)
	fmt.Println(v.Meta)
	e, ok := v.Meta.(*error.Error)
	assert.Equal(t, true, ok)
	responseLog.Code = e.Response.Code
	responseLog.Message = e.Response.Message
	responseLog.Data = e.Response.Details
	log.Error(responseLog)
}
