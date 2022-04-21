package dodo_open_go

import (
	"dodo-open-go/client"
	"dodo-open-go/version"
	"testing"
	"time"
)

func TestNewInstance(t *testing.T) {
	instance, err := NewInstance("", "", client.WithTimeout(time.Second*3))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", instance)
	t.Logf("%+v", instance.GetConfig())
}

func TestVersion(t *testing.T) {
	t.Log(version.Version())
}
