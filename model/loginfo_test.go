package model

import (
	"testing"
)

func TestLogAdd(t *testing.T) {
	log := &LogInfo{IP: "127.0.0.1", UserAgent: "FF"}

	r, err := log.Add()
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("success:%d", r)
	}
}

func TestLogGet(t *testing.T) {
	var log LogInfo
	res, err := log.GetList(0, 10, nil, nil)
	if err != nil {
		t.Error(err)
	} else {
		for _, v := range res {
			t.Log(v)
		}
	}
}

func TestLogUpdate(t *testing.T) {
	log := &LogInfo{ID: 3, Msg: "duang~duang", Url: "http://www.g.com"}
	r, err := log.UpdateByID("Msg,Url", "")
	if err != nil {
		t.Error(err)
	} else {
		t.Log("success:", r)
	}
}
