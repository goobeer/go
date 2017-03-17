package model

import (
	"testing"
)

func TestUserAdd(t *testing.T) {
	u1 := &Users{Name: "zs", Pwd: "123456", Used: true}
	u2 := &Users{Name: "ls", Pwd: "123456", Used: true}
	r, err := u1.Add(u1, u2)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(r)
	}
}

func TestUserDelete(t *testing.T) {
	u1 := &Users{}

	r, err := u1.Delete("Name=?", "zs")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(r)
	}
}

func TestUserSlt(t *testing.T) {
	u1 := &Users{ID: 1}
	r, err := u1.GetList(1, 10, "id>?", u1)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(r)
	}
}
