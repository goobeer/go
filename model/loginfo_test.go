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

func TestGetList(t *testing.T) {
	var log LogInfo
	res, err := log.GetList(0, 10)
	if err != nil {
		t.Error(err)
	} else {
		t.Log("###", res, len(res))
	}
}
