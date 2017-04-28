package commonhtmlhelper

import (
	"fmt"
)

func GeneratePageLink(currentIndex, showCount, totalNum int, pageKey string, commonData map[string]interface{}) string {
	var pageLinkHtml string
	var indexs []int
	if currentIndex <= totalNum && showCount <= totalNum {
		if showCount <= totalNum {
			maxIndex := currentIndex + showCount/2
			if maxIndex <= totalNum {
				for start := currentIndex - showCount/2; start <= maxIndex; start++ {
					indexs = append(indexs, start)
				}
			} else {
				for start := totalNum - showCount; start <= totalNum; start++ {
					indexs = append(indexs, start)
				}
			}
		} else {
			for start := 1; start < totalNum; start++ {
				indexs = append(indexs, start)
			}
		}
	}

	if len(indexs) > 0 {
		if len(pageKey) == 0 {
			pageKey = "page"
		}

		var commonParaPartial string

		if commonData != nil && len(commonData) > 0 {
			i := 0
			for k, v := range commonData {
				if i == 0 {
					i++
				}
				joinFlag := ""
				if i != 1 {
					joinFlag = "&"
				}
				commonParaPartial += fmt.Sprintf("%s%s=%v", joinFlag, k, v)
			}
			commonParaPartial = "?" + commonParaPartial
		}

		var pageItems string

		for _, v := range indexs {
			pageItems += fmt.Sprintf(`<li><a href="?%s&%s=%d">%d</a></li>`, commonParaPartial, pageKey, v)
		}
		if len(indexs) == 1 {
			pageLinkHtml = fmt.Sprintf(`<nav aria-label="Page navigation">
  <ul class="pagination">
    %s
  </ul>
</nav>`, pageItems)
		} else {
			pageLinkHtml = fmt.Sprintf(`<nav aria-label="Page navigation">
  <ul class="pagination">
    <li>
      <a href="#" aria-label="Previous">
        <span aria-hidden="true">&laquo;</span>
      </a>
    </li>
    %s
    <li>
      <a href="#" aria-label="Next">
        <span aria-hidden="true">&raquo;</span>
      </a>
    </li>
  </ul>
</nav>`, pageItems)
		}
	}

	return pageLinkHtml
}

func linkItemHelper(i, totalNum int) string {
	var itemStr string
	if i == 1 {
		itemStr = `<li>
      <a href="#" aria-label="Previous">
        <span aria-hidden="true">&laquo;</span>
      </a>
    </li>`
	} else if i == totalNum {
		itemStr = `<li>
      <a href="#" aria-label="Next">
        <span aria-hidden="true">&raquo;</span>
      </a>
    </li>`
	} else {
		itemStr = fmt.Sprintf(`<li><a href="#">%d</a></li>`, i)
	}
	return itemStr
}
