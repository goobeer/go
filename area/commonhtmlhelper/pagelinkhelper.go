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
				i++

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
				pageItems += fmt.Sprintf(`<li class="active"><a>%d</a></li>`, v)
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
			var prevPartial, nextPartial, firstPartial, lastPartial, parasStr string
			if len(commonParaPartial) > 0 {
				parasStr = fmt.Sprintf("%s&%s=", commonParaPartial, pageKey)
			} else {
				parasStr = fmt.Sprintf("%s=", pageKey)
			}
			if prev > 1 {
				prevPartial = fmt.Sprintf(`<li>
      <a href="%s?%s%d" aria-label="Previous">
        <span aria-hidden="true">&lt;</span>
      </a>
    </li>`, baseUrl, parasStr, prev)
			}
			if next < totalNum && next > 0 {
				nextPartial = fmt.Sprintf(`<li>
      <a href="%s?%s%d" aria-label="Next">
        <span aria-hidden="true">&gt;</span>
      </a>
    </li>`, baseUrl, parasStr, next)
			}
			if currentIndex == 1 {
				firstPartial = `<li class="disabled">
      <a aria-label="First">
        <span aria-hidden="true">&laquo;</span>
      </a>
    </li>`
				lastPartial = fmt.Sprintf(`<li>
      <a href="%s?%s%d" aria-label="Last">
        <span aria-hidden="true">&raquo;</span>
      </a>
    </li>`, baseUrl, parasStr, totalNum)
			} else if currentIndex == totalNum {
				firstPartial = fmt.Sprintf(`<li>
      <a href="%s?%s%d" aria-label="First">
        <span aria-hidden="true">&laquo;</span>
      </a>
    </li>`, baseUrl, parasStr, 1)
				lastPartial = `<li class="disabled">
      <a aria-label="Last">
        <span aria-hidden="true">&raquo;</span>
      </a>
    </li>`
			} else {
				firstPartial = fmt.Sprintf(`<li>
      <a href="%s?%s%d" aria-label="First">
        <span aria-hidden="true">&laquo;</span>
      </a>
    </li>`, baseUrl, parasStr, 1)
				lastPartial = fmt.Sprintf(`<li>
      <a href="%s?%s%d" aria-label="Last">
        <span aria-hidden="true">&raquo;</span>
      </a>
    </li>`, baseUrl, parasStr, totalNum)
			}

			pageLinkHtml = fmt.Sprintf(`<nav aria-label="Page navigation"><ul class="pagination">%s%s%s%s%s</ul></nav>`, firstPartial, prevPartial, pageItems, nextPartial, lastPartial)
		}
	}

	return pageLinkHtml
}
