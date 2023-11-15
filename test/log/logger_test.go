package testlogger

import (
	"app/constanta"
	"app/pkg/log"
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestInfo(t *testing.T) {
	log.Info(1, "dasda", constanta.StatusOK, constanta.CodeOK)

}
