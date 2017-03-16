package model

import (
	"testing"
)

func TestAdd(t *testing.T) {
	log := &LogInfo{IP: "127.0.0.1", UserAgent: "FF"}

	r, err := log.Add()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(r)
	}
}
