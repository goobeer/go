package ueditor_go

import (
	"fasthttpweb/common"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func UeditorPathFormatter(originFileName, pathFormat string) string {
	if len(pathFormat) == 0 {
		pathFormat = "{filename}{rand:6}"
	}
	invalidPattern, _ := regexp.Compile("[\\\\\\/\\:\\*\\?\\042\\<\\>\\|]")
	originFileName = invalidPattern.ReplaceAllString(originFileName, "")

	extension := path.Ext(originFileName)
	filename := strings.TrimRight(originFileName, extension)
	pathFormat = strings.Replace(pathFormat, "{filename}", filename, -1)
	regPathFormat, _ := regexp.Compile("\\{rand(\\:?)(\\d+)\\}")

	pathSlices := regPathFormat.FindStringSubmatch(pathFormat)
	if len(pathSlices) > 0 {
		digit := 6
		digit, _ = strconv.Atoi(pathSlices[len(pathSlices)-1])
		randStr, _ := common.CreateRandom(0, digit)
		pathFormat = strings.Replace(pathFormat, regPathFormat.FindString(pathFormat), randStr, -1)
	}

	pathFormat = strings.Replace(pathFormat, "{time}", strconv.FormatInt(time.Now().UnixNano(), 10), 1)
	pathFormat = strings.Replace(pathFormat, "{yyyy}", time.Now().Format("2006"), 1)
	pathFormat = strings.Replace(pathFormat, "{yy}", time.Now().Format("06"), 1)
	pathFormat = strings.Replace(pathFormat, "{mm}", time.Now().Format("01"), 1)
	pathFormat = strings.Replace(pathFormat, "{dd}", time.Now().Format("02"), 1)
	pathFormat = strings.Replace(pathFormat, "{hh}", time.Now().Format("15"), 1)
	pathFormat = strings.Replace(pathFormat, "{ii}", time.Now().Format("04"), 1)
	pathFormat = strings.Replace(pathFormat, "{ss}", time.Now().Format("05"), 1)
	pathFormat += extension
	return pathFormat
}
