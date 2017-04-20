package ueditor_go

//	"strconv"

type UeditorListFileManager struct {
	*UEditorHandler
}

func (u *UeditorListFileManager) Process() {
	//	start := u.Ctx.FormValue("start")
	//	startV, sizeV := 0, 0

	//	if start != nil && len(start) > 0 {
	//		startStr := string(start)
	//		startV, _ = strconv.Atoi(startStr)
	//	}
	//	sizev := u.Ctx.FormValue("size")
	//	if sizev != nil && len(sizev) > 0 {
	//		sizevStr := string(sizev)
	//		sizeV, err := strconv.Atoi(sizevStr)
	//		if err != nil {
	//			sizeV = int(config.ImageManagerListSize)
	//		}
	//	} else {

	//	}
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
