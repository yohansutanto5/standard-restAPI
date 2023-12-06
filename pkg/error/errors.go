package error

import (
	"app/model"
	"app/pkg/log"
	"fmt"
	"net/http"
	"regexp"
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
	CodeErrorDBDuplicate = "APP-DB-1093"
	CodeErrorServiceUser = "APP-SRV-01"
)

// Custom Status

const (
	StatusDebug = 999
)

type Error struct {
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

func (e *Error) GenerateReponse(transcID int) {
	e.Response.TransactionID = transcID
	e.Response.Code = e.Code
	e.Response.Message = e.Message
	e.Response.Details = e.Details
}

func (e *Error) LogError() {
	customLog := model.CustomLog{}
	log.Error(customLog)
}
func (e *Error) SetStatus() {
	if e.Code == "" {
		e.Status = http.StatusBadRequest
	}
}

func (e *Error) ParseMysqlError(err error) {
	if err == nil {
		return
	}

	re := regexp.MustCompile(`Error (\d+): (.+)`)
	match := re.FindStringSubmatch(err.Error())
	if len(match) == 3 {
		e.Message = match[2]
		e.Code = "APP-DB-" + match[1]
	}

	// handle the HTTP status
	switch match[1] {
	// Duplicate Entry
	case CodeErrorDBDuplicate:
		e.Status = http.StatusConflict
	case "12412":
		e.Status = http.StatusUnprocessableEntity
	default:
		e.Status = http.StatusInternalServerError
	}
}
