package controllers

import (
	"fasthttpweb/area"
	view "fasthttpweb/area/admin/views/loginfo"
	"fasthttpweb/model"
	"math"
	"strconv"
)

type LoginfoController struct {
	*area.BaseController
}

func init() {
	c := &LoginfoController{BaseController: &area.BaseController{}}
	c.RegistRoutes(areaName, c)
	c.RegistRoute("get", "/admin/loginfo/index", "/admin/loginfo")
}

func (c *LoginfoController) Index() {
	bpd := area.NewBasePageData(areaName, "admin-loginfo-title", "admin-loginfo-kwd", c)
	bp := area.NewBasePage(c.Ctx, bpd)
	currentPage := c.Ctx.FormValue("page")
	currentPageVal := 1
	if currentPage != nil && len(currentPage) > 0 {
		currentPageVal, _ = strconv.Atoi(string(currentPage))
		if currentPageVal < 1 {
			currentPageVal = 1
		}
	}
	bpd.Data["page"] = currentPageVal

	loginfo := &model.LogInfo{}
	var loginfoes []model.LogInfo
	pageSize := 5
	err := loginfo.DbListCommandSession(uint(currentPageVal-1), uint(pageSize), nil).Desc("CreateTime").Find(&loginfoes)
	if err != nil {
		bpd.Data["error"] = err.Error()
	} else {
		bpd.Data["Model"] = loginfoes
		total, _ := loginfo.DbCommandSession(nil).Count(loginfo)
		totalNum := float64(total)
		bpd.Data["Total"] = int(math.Ceil(totalNum / float64(pageSize)))
	}
	ip := &view.IndexPage{bp}
	c.View(ip, "text/html")
}
