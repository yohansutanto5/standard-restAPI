package log

import "testing"

func TestMain(m *testing.M) {
	m.Run()
}

func TestInfo(t *testing.T) {
	Info(1, "dasda", nil)

}
