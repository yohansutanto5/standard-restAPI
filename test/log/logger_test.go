package log

import (
	"app/pkg/log"
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestInfo(t *testing.T) {
	log.Info(1, "dasda", nil)

}
