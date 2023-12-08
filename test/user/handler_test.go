package user_test

import (
	"app/model"
	"app/pkg/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetList(t *testing.T) {
	userHandler.GetList(ctx)
	assert.Equal(t, 200, ctx.Writer.Status())
}

func TestInsert(t *testing.T) {
	// Instantiate test data
	var data model.AddUserIn
	data.FirstName = util.GenerateRandomString(5)
	data.LastName = util.GenerateRandomString(4)

	// Inject test data to request context

	// Test the handler
	userHandler.Insert(ctx)
	assert.Equal(t, 200, ctx.Writer.Status())
}
