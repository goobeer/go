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
	var res []*LogInfo
	err := log.GetList(&res, 0, 10, nil, nil)
	if err != nil {
		t.Error(err)
	} else {
		for _, v := range res {
			t.Log(v)
		}
	}
}

func TestLogUpdate(t *testing.T) {
	log := &LogInfo{BaseModel: &BaseModel{ID: 3}, Msg: "guang~duang", Url: "http://www.g.com"}
	r, err := log.UpdateByID(log, "Msg,Url", "")
	if err != nil {
		t.Error(err)
	} else {
		t.Log("success:", r)
	}
}
