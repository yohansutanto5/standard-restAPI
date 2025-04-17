package test

import (
	"app/pkg/error"
	newerr "app/pkg/error"
	"app/pkg/util"
	"bytes"
	"errors"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SourceStruct struct {
	SourceField1 int          `convert:"TargetField1"`
	SourceField2 string       `convert:"TargetField2"`
	NestedStruct NestedStruct `convert:"NestedTargetStruct"`
}

type TargetStruct struct {
	TargetField1 int
	TargetField2 string
	NestedStruct NestedTargetStruct
}

type NestedStruct struct {
	NestedField1 int `convert:"NestedTargetField1"`
	NestedField2 string
}

type NestedTargetStruct struct {
	NestedTargetField1 int
	NestedTargetField2 string
}

func TestStructConverterToMap(t *testing.T) {
	sourceInstance := SourceStruct{
		SourceField1: 42,
		SourceField2: "Hello, World!",
		NestedStruct: NestedStruct{
			NestedField1: 24,
			NestedField2: "Nested Hello!",
		},
	}
	fieldsToInclude := []string{"TargetField1", "TargetField2", "NestedTargetStruct", "NestedTargetField1"}
	mapped := util.ConvertStructToMap(sourceInstance, fieldsToInclude)
	assert.Equal(t, sourceInstance.SourceField1, mapped["TargetField1"])
}

func TestError(t *testing.T) {
	err := &newerr.Error{}
	e := errors.New("Error 1062 (23000): Duplicate entry 'abc567' for key 'users.username'")
	// e := errors.New("Error 1146 (42S02): Table 'devsync.uats' doesn't exist")

	err = error.ParseMysqlError(e)
	assert.Equal(t, err.Status, 409)
	// assert.Equal(t,"APP-DB-1146",err.Code)
}

func TestExec(t *testing.T) {
	a := exec.Command("pwd")
	var stdout bytes.Buffer
	a.Stdout = &stdout
	err := a.Run()
	t.Logf(stdout.String())
	assert.Nil(t, err)
}
