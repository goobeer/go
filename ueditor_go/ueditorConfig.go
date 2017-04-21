package ueditor_go

import (
	"encoding/json"
	"os"
	"regexp"
)

const (
	noCache           = true
	ueditorConfigFile = "./public/ueditor/net/config.json"
)

type UeditorConfig struct {
	ImageActionName     string   `json:"imageActionName"`
	ImageFieldName      string   `json:"imageFieldName"`
	ImageMaxSize        int64    `json:"imageMaxSize"`
	ImageAllowFiles     []string `json:"imageAllowFiles"`
	ImageCompressEnable bool     `json:"imageCompressEnable"`
	ImageCompressBorder int64    `json:"imageCompressBorder"`
	ImageInsertAlign    string   `json:"imageInsertAlign"`
	ImageUrlPrefix      string   `json:"imageUrlPrefix"`
	ImagePathFormat     string   `json:"imagePathFormat"`

	ScrawlActionName      string   `json:"scrawlActionName"`
	ScrawlFieldName       string   `json:"scrawlFieldName"`
	ScrawlPathFormat      string   `json:"scrawlPathFormat"`
	ScrawlMaxSize         int64    `json:"scrawlMaxSize"`
	ScrawlUrlPrefix       string   `json:"scrawlUrlPrefix"`
	ScrawlInsertAlign     string   `json:"scrawlInsertAlign"`
	SnapscreenActionName  string   `json:"snapscreenActionName"`
	SnapscreenPathFormat  string   `json:"snapscreenPathFormat"`
	SnapscreenUrlPrefix   string   `json:"snapscreenUrlPrefix"`
	SnapscreenInsertAlign string   `json:"snapscreenInsertAlign"`
	CatcherLocalDomain    []string `json:"catcherLocalDomain"`
	CatcherActionName     string   `json:"catcherActionName"`
	CatcherFieldName      string   `json:"catcherFieldName"`
	CatcherPathFormat     string   `json:"catcherPathFormat"`
	CatcherUrlPrefix      string   `json:"catcherUrlPrefix"`
	CatcherMaxSize        int64    `json:"catcherMaxSize"`
	CatcherAllowFiles     []string `json:"catcherAllowFiles"`

	VideoActionName string   `json:"videoActionName"`
	VideoFieldName  string   `json:"videoFieldName"`
	VideoPathFormat string   `json:"videoPathFormat"`
	VideoUrlPrefix  string   `json:"videoUrlPrefix"`
	VideoMaxSize    int64    `json:"videoMaxSize"`
	VideoAllowFiles []string `json:"videoAllowFiles"`

	FileActionName string   `json:"fileActionName"`
	FileFieldName  string   `json:"fileFieldName"`
	FilePathFormat string   `json:"filePathFormat"`
	FileUrlPrefix  string   `json:"fileUrlPrefix"`
	FileMaxSize    int64    `json:"fileMaxSize"`
	FileAllowFiles []string `json:"fileAllowFiles"`

	ImageManagerActionName  string   `json:"imageManagerActionName"`
	ImageManagerListPath    string   `json:"imageManagerListPath"`
	ImageManagerListSize    int64    `json:"imageManagerListSize"`
	ImageManagerUrlPrefix   string   `json:"imageManagerUrlPrefix"`
	ImageManagerInsertAlign string   `json:"imageManagerInsertAlign"`
	ImageManagerAllowFiles  []string `json:"imageManagerAllowFiles"`

	FileManagerActionName string   `json:"fileManagerActionName"`
	FileManagerListPath   string   `json:"fileManagerListPath"`
	FileManagerUrlPrefix  string   `json:"fileManagerUrlPrefix"`
	FileManagerListSize   int64    `json:"fileManagerListSize"`
	FileManagerAllowFiles []string `json:"fileManagerAllowFiles"`
}

func BuildItems() (config UeditorConfig, err error) {
	var configFile *os.File
	configFile, err = os.Open(ueditorConfigFile)
	defer configFile.Close()

	if err != nil {
		return
	}
	var jsonDataBuf []byte
	buf := make([]byte, 1024*4)
	for {
		n, _ := configFile.Read(buf)

		if 0 == n {
			break
		}
		jsonDataBuf = append(jsonDataBuf, buf[:n]...)
	}

	reg, _ := regexp.Compile("/\\*.*\\*/")
	jsonDataBuf = reg.ReplaceAll(jsonDataBuf, []byte(""))
	err = json.Unmarshal(jsonDataBuf, &config)
	return
}
