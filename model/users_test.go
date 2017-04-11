package model

import (
	"testing"
)

func TestUserAdd(t *testing.T) {
	u1 := &Users{Name: "mawu", Pwd: "123456", Used: true}
	u2 := &Users{Name: "zhangfei", Pwd: "123456", Used: true}
	r, err := u1.Add(u1, u2)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(r)
	}
}

func TestUserDelete(t *testing.T) {
	u1 := &Users{}

	r, err := u1.Delete(u1, "id>?", 1)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(r)
	}
}

func TestUserGet(t *testing.T) {
	u1 := &Users{}
	var r []*Users
	err := u1.GetList(&r, 0, 10, "id>?", 18)

	if err != nil {
		t.Error(err)
	} else {
		for _, v := range r {
			t.Log("success:", v)
		}
	}
}

func TestUserUpdate(t *testing.T) {
	u1 := &Users{BaseModel: &BaseModel{ID: 20}, Used: false, Name: "JJ"}
	r, err := u1.UpdateByID(u1, "Used,CreateTime", "Name")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(r)
	}
}
