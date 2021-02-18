package ini

import (
	"testing"
)

func TestINI(t *testing.T) {
	cfg := SetConfig("test.ini")
	port := cfg.GetValue("server", "port")
	t.Log(port)
}
