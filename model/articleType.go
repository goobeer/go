package model

type ArticleType struct {
	*BaseModel `xorm:"extends"`
	Name       string
	Ord        int
}
