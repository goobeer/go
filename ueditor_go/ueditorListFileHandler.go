package ueditor_go

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

type UeditorListFileRespData struct {
	State string         `json:"state"`
	List  []ListFileItem `json:"list"`
	Start int            `json:"start"`
	Size  int            `json:"size"`
	Total int            `json:"total"`
}

type ListFileItem struct {
	Url string `json:"url"`
}

type UeditorListFileManager struct {
	*UEditorHandler
	PathToList       string
	SearchExtensions []string
}

func (u *UeditorListFileManager) Process() {
	start := u.Ctx.FormValue("start")
	startV, sizeV := 0, 0

	if start != nil && len(start) > 0 {
		startStr := string(start)
		startV, _ = strconv.Atoi(startStr)
	}

	TD := &UeditorListFileRespData{}
	TD.Start = startV
	sizev := u.Ctx.FormValue("size")
	if sizev != nil && len(sizev) > 0 {
		sizevStr := string(sizev)
		var err error
		sizeV, err = strconv.Atoi(sizevStr)
		if err != nil {
			TD.State = u.getStateString(InvalidParam)
			return
		}
	} else {
		sizeV = int(config.ImageManagerListSize)
	}
	TD.Size = sizeV
	buildingList := make([]ListFileItem, 0)
	fmt.Println(u.PathToList)
	err := filepath.Walk(u.PathToList, func(pth string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}

		if f.IsDir() {
			return nil
		}

		for _, v := range u.SearchExtensions {
			if strings.HasSuffix(v, strings.ToLower(path.Ext(f.Name()))) {
				buildingList = append(buildingList, ListFileItem{Url: strings.TrimLeft(strings.Replace(pth, "\\", "/", -1), u.PathToList)})
			}
		}

		return nil
	})

	if err != nil {
		TD.State = u.getStateString(AuthorizError)
	} else {
		TD.State = u.getStateString(Success)
		TD.Total = len(buildingList)
		TD.List = buildingList
	}

	u.WriteJson(TD)
}

func (u *UeditorListFileManager) getStateString(state int) (stateStr string) {
	switch state {
	case Success:
		stateStr = "SUCCESS"
	case InvalidParam:
		stateStr = "参数不正确"
	case PathNotFound:
		stateStr = "路径不存在"
	case AuthorizError:
		stateStr = "文件系统权限不足"
	case IOError:
		stateStr = "文件系统读取错误"
	default:
		stateStr = "未知错误"
	}
	return stateStr
}
