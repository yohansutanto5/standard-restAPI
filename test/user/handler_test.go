package user_test

import (
	"app/model"
	"app/pkg/util"
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetList(t *testing.T) {
	userHandler.GetList(ctx)
	assert.Equal(t, 200, ctx.Writer.Status())
}

func TestInsert(t *testing.T) {
	// Instantiate test data
	var data model.UserInput
	data.FirstName = util.GenerateRandomString(5)
	data.LastName = util.GenerateRandomString(4)
	data.Email = "asdv@gma.com"
	data.Username = util.GenerateRandomString(5)

	// Inject test data to request context
	// Marshal data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		t.FailNow()
	}
	reader := strings.NewReader(string(jsonData))
	req, _ := http.NewRequest("GET", "/mockurl", reader)
	ctx.Request = req

	// Test the handler
	userHandler.Insert(ctx)
	assert.Equal(t, 200, ctx.Writer.Status())
}
