package testtemplate

import (
	"app/internal/template"
	"testing"
)

func TestGetListStudent(t *testing.T) {
	template.GetStudent(ctx, dbg)
	if ctx.Writer.Status() != 200 {
		t.Fail()
	}
}
