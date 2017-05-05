package commonhtmlhelper

import (
	"fmt"
)

func GeneratePageLink(baseUrl string, currentIndex, showCount, totalNum int, pageKey string, commonData map[string]interface{}) string {
	var pageLinkHtml string
	var indexs []int
	if !(currentIndex >= 1 && currentIndex <= totalNum) {
		currentIndex = 1
	}
	if currentIndex <= totalNum && showCount <= totalNum {
		if showCount <= totalNum {
			maxIndex := currentIndex + showCount/2
			if maxIndex <= totalNum {
				start := currentIndex - showCount/2
				if start > 0 {
					for ; start <= maxIndex; start++ {
						indexs = append(indexs, start)
					}
				} else {
					start = 1
					for ; start <= showCount; start++ {
						indexs = append(indexs, start)
					}
				}
			} else {
				for start := totalNum - showCount; start <= totalNum; start++ {
					indexs = append(indexs, start)
				}
			}
		} else {
			for start := 1; start <= totalNum; start++ {
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
			var joinFlag string
			for k, v := range commonData {
				if i == 0 {
					i++
				}

				if i != 1 && len(joinFlag) == 0 {
					joinFlag = "&"
				}
				commonParaPartial += fmt.Sprintf("%s%s=%v", joinFlag, k, v)
			}
		}

		var pageItems string

		for _, v := range indexs {
			var parasStr string
			if len(commonParaPartial) > 0 {
				parasStr = fmt.Sprintf("%s&%s=%d", commonParaPartial, pageKey, v)
			} else {
				parasStr = fmt.Sprintf("%s=%d", pageKey, v)
			}
			if v == currentIndex {
				pageItems += fmt.Sprintf(`<li><a>%d</a></li>`, v)
			} else {
				pageItems += fmt.Sprintf(`<li><a href="%s?%s">%d</a></li>`, baseUrl, parasStr, v)
			}
		}

		if totalNum <= showCount {
			pageLinkHtml = fmt.Sprintf(`<nav aria-label="Page navigation">
  <ul class="pagination">
    %s
  </ul>
</nav>`, pageItems)
		} else {
			var prev, next int
			if currentIndex == 1 {
				next = currentIndex + 1
			} else if currentIndex == totalNum {
				prev = currentIndex - 1
			} else {
				prev = currentIndex - 1
				next = currentIndex + 1
			}
			var prevPartial, nextPartial string
			if prev > 1 {
				var parasStr string
				if len(commonParaPartial) > 0 {
					parasStr = fmt.Sprintf("%s&%s=%d", commonParaPartial, pageKey, prev)
				} else {
					parasStr = fmt.Sprintf("%s=%d", pageKey, prev)
				}
				prevPartial = fmt.Sprintf(`<li>
      <a href="%s?%s" aria-label="Previous">
        <span aria-hidden="true">&laquo;</span>
      </a>
    </li>`, baseUrl, parasStr)
			}
			if next < totalNum {
				var parasStr string
				if len(commonParaPartial) > 0 {
					parasStr = fmt.Sprintf("%s&%s=%d", commonParaPartial, pageKey, next)
				} else {
					parasStr = fmt.Sprintf("%s=%d", pageKey, next)
				}
				nextPartial = fmt.Sprintf(`<li>
      <a href="%s?%s" aria-label="Next">
        <span aria-hidden="true">&raquo;</span>
      </a>
    </li>`, baseUrl, parasStr)
			}
			pageLinkHtml = fmt.Sprintf(`<nav aria-label="Page navigation"><ul class="pagination">%s%s%s</ul></nav>`, prevPartial, pageItems, nextPartial)
		}
	}

	return pageLinkHtml
}
