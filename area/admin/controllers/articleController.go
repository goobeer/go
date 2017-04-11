package controllers

import (
	"fasthttpweb/area"
)

type ArticleController struct {
	*area.BaseController
}

func init() {
	c := ArticleController{BaseController: &area.BaseController{}}
	c.RegistRoutes(areaName, c)
}

func (c *ArticleController) Index() {

}
