package controllers

import (
	"fasthttpweb/area"

	//	view "fasthttpweb/area/article/views/article"
)

type ArticleController struct {
	*area.BaseController
}

func init() {
	c := &ArticleController{BaseController: &area.BaseController{}}
	c.RegistRoutes(areaName, c)
}

func (c *ArticleController) Index() {

}

func (c *ArticleController) Detail() {

}
