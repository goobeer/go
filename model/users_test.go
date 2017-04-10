package model

import (
	"testing"
)

func TestUserAdd(t *testing.T) {
	u1 := &Users{Name: "ms", Pwd: "123456", Used: true}
	u2 := &Users{Name: "ws", Pwd: "123456", Used: true}
	r, err := u1.Add(u1, u2)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(r)
	}
}

func TestUserDelete(t *testing.T) {
	u1 := &Users{}

	r, err := u1.Delete("id>?", 1)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(r)
	}
}

func TestUserGet(t *testing.T) {
	u1 := &Users{}

	r, err := u1.GetList(0, 10, "createtime<?", "2015")

	if err != nil {
		t.Error(err)
	} else {
		for _, v := range r {
			t.Log("success:", v)
		}
	}
}

func TestUserUpdate(t *testing.T) {
	u1 := &Users{ID: 19, Used: false, Name: "wxbq"}
	r, err := u1.UpdateByID("Used,CreateTime", "Name")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(r)
	}
}
