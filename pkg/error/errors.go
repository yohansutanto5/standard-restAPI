package error

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

var mappingStatus = map[int][]string{
	http.StatusBadRequest:   {"asd", "sdasd", "lol"},
	http.StatusForbidden:    {"asdas", "bibibi", "laooa"},
	http.StatusConflict:     {"asdas", "bibibi", "laooa"},
	http.StatusUnauthorized: {"asdas", "bibibi", "laooa"},
}

// Error Code
const (
	CodeDebug            = "SYS-DBG"
	CodeWarningRecover   = "SYS-WARN-RECOV"
	CodeOK               = "SUCCESS"
	CodeErrorDBDuplicate = "APP-DB-1062"
	CodeErrorServiceUser = "APP-SRV-01"
)

// Custom Status

const (
	StatusDebug = 999
)

type Error struct {
	error
	Code     string
	Message  string
	Status   int
	Response ErrorResponse
	Details  any
}

type ErrorResponse struct {
	TransactionID int
	Code          string
	Message       string
	Details       any
}

func (e *Error) String() string {
	return fmt.Sprintf("Error %s: %s", e.Code, e.Message)
}

func (e *Error) GenerateReponse(transcID int) *gin.Error {
	e.Response.TransactionID = transcID
	e.Response.Code = e.Code
	e.Response.Message = e.Message
	e.Response.Details = e.Details

	return &gin.Error{
		Err:  e.error,
		Type: gin.ErrorTypePrivate,
		Meta: e,
	}
}
func New(err error) *Error {
	return &Error{
		error: err,
	}
}

func ParseMysqlError(err error) *Error {
	if err == nil {
		return nil
	}
	e := &Error{}

	re := regexp.MustCompile(`Error (.+) \((.+)\): (.+)`)
	match := re.FindStringSubmatch(err.Error())
	if len(match) >= 4 {
		e.Message = match[3]
		e.Code = "APP-DB-" + match[1]
	} else {
		e.Code = "APP-DB-UNKNOWN"
		e.Message = err.Error()
	}
	// handle the HTTP status
	switch e.Code {
	// Duplicate Entry
	case CodeErrorDBDuplicate:
		e.Status = http.StatusConflict
	case "12412":
		e.Status = http.StatusUnprocessableEntity
	default:
		e.Status = http.StatusInternalServerError
	}
	return e
}
