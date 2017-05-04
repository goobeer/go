package filter

import (
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
)

type PerformanceMonitorFilter struct {
	*FilterStateResult
	time.Time
}

func (filter *PerformanceMonitorFilter) BeforeExecute(ctx *fasthttp.RequestCtx) *FilterStateResult {
	filter.Time = time.Now()
	fmt.Printf("[performance-begin]:%s(start-time:%s)\r\n", string(ctx.RequestURI()), filter.Time.Format("2016-01-02 15:04:05"))
	return filter.FilterStateResult
}

func (filter *PerformanceMonitorFilter) AfterExecute(ctx *fasthttp.RequestCtx) *FilterStateResult {
	fmt.Printf("[performance-end]:%s(end-time:%s,spend:%v)\r\n", string(ctx.RequestURI()), time.Now().Format("2016-01-02 15:04:05"), time.Now().Sub(filter.Time).Nanoseconds())
	return filter.FilterStateResult
}
