package model

import (
	"testing"
)

func TestArticleAdd(t *testing.T) {
	var arts []*Article
	arts = append(arts, &Article{Title: "党中央", Content: "不错哟", Used: true})
	arts = append(arts, &Article{Title: "中央党", Content: "不错哟", Used: true})
	art := new(Article)
	r, err := art.Add(arts...)
	if err != nil {
		t.Error(err)
	} else {
		t.Log("Article added success:", r)
	}
}

func TestArticleDelete(t *testing.T) {
	art := &Article{}
	r, err := art.Delete(art, "id>=?", 4)
	if err != nil {
		t.Error(err)
	} else {
		t.Log("Article delete success:", r)
	}
}

func TestArticleUpdateByID(t *testing.T) {
	art := &Article{BaseModel: &BaseModel{ID: 3}, Title: "我打啊"}
	r, err := art.UpdateByID(art, "Title,CreateTime", "")
	if err != nil {
		t.Error(err)
	} else {
		t.Log("Article update success:", r)
	}
}
