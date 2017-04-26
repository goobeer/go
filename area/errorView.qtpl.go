// This file is automatically generated by qtc from "errorView.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line errorView.qtpl:1
package area

//line errorView.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line errorView.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line errorView.qtpl:2
type ErrorPage struct {
	*BasePage
}

//line errorView.qtpl:7
func (p *ErrorPage) StreamRenderTitle(qw422016 *qt422016.Writer) {
	//line errorView.qtpl:7
	p.BasePage.StreamRenderTitle(qw422016)
	//line errorView.qtpl:7
	qw422016.N().S(` `)
//line errorView.qtpl:7
}

//line errorView.qtpl:7
func (p *ErrorPage) WriteRenderTitle(qq422016 qtio422016.Writer) {
	//line errorView.qtpl:7
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line errorView.qtpl:7
	p.StreamRenderTitle(qw422016)
	//line errorView.qtpl:7
	qt422016.ReleaseWriter(qw422016)
//line errorView.qtpl:7
}

//line errorView.qtpl:7
func (p *ErrorPage) RenderTitle() string {
	//line errorView.qtpl:7
	qb422016 := qt422016.AcquireByteBuffer()
	//line errorView.qtpl:7
	p.WriteRenderTitle(qb422016)
	//line errorView.qtpl:7
	qs422016 := string(qb422016.B)
	//line errorView.qtpl:7
	qt422016.ReleaseByteBuffer(qb422016)
	//line errorView.qtpl:7
	return qs422016
//line errorView.qtpl:7
}

//line errorView.qtpl:8
func (p *ErrorPage) StreamRenderKwd(qw422016 *qt422016.Writer) {
	//line errorView.qtpl:8
	p.BasePage.StreamRenderKwd(qw422016)
	//line errorView.qtpl:8
	qw422016.N().S(` `)
//line errorView.qtpl:8
}

//line errorView.qtpl:8
func (p *ErrorPage) WriteRenderKwd(qq422016 qtio422016.Writer) {
	//line errorView.qtpl:8
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line errorView.qtpl:8
	p.StreamRenderKwd(qw422016)
	//line errorView.qtpl:8
	qt422016.ReleaseWriter(qw422016)
//line errorView.qtpl:8
}

//line errorView.qtpl:8
func (p *ErrorPage) RenderKwd() string {
	//line errorView.qtpl:8
	qb422016 := qt422016.AcquireByteBuffer()
	//line errorView.qtpl:8
	p.WriteRenderKwd(qb422016)
	//line errorView.qtpl:8
	qs422016 := string(qb422016.B)
	//line errorView.qtpl:8
	qt422016.ReleaseByteBuffer(qb422016)
	//line errorView.qtpl:8
	return qs422016
//line errorView.qtpl:8
}

//line errorView.qtpl:10
func (p *ErrorPage) StreamRenderCss(qw422016 *qt422016.Writer) {
//line errorView.qtpl:10
}

//line errorView.qtpl:10
func (p *ErrorPage) WriteRenderCss(qq422016 qtio422016.Writer) {
	//line errorView.qtpl:10
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line errorView.qtpl:10
	p.StreamRenderCss(qw422016)
	//line errorView.qtpl:10
	qt422016.ReleaseWriter(qw422016)
//line errorView.qtpl:10
}

//line errorView.qtpl:10
func (p *ErrorPage) RenderCss() string {
	//line errorView.qtpl:10
	qb422016 := qt422016.AcquireByteBuffer()
	//line errorView.qtpl:10
	p.WriteRenderCss(qb422016)
	//line errorView.qtpl:10
	qs422016 := string(qb422016.B)
	//line errorView.qtpl:10
	qt422016.ReleaseByteBuffer(qb422016)
	//line errorView.qtpl:10
	return qs422016
//line errorView.qtpl:10
}

//line errorView.qtpl:12
func (p *ErrorPage) StreamRenderScript(qw422016 *qt422016.Writer) {
//line errorView.qtpl:12
}

//line errorView.qtpl:12
func (p *ErrorPage) WriteRenderScript(qq422016 qtio422016.Writer) {
	//line errorView.qtpl:12
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line errorView.qtpl:12
	p.StreamRenderScript(qw422016)
	//line errorView.qtpl:12
	qt422016.ReleaseWriter(qw422016)
//line errorView.qtpl:12
}

//line errorView.qtpl:12
func (p *ErrorPage) RenderScript() string {
	//line errorView.qtpl:12
	qb422016 := qt422016.AcquireByteBuffer()
	//line errorView.qtpl:12
	p.WriteRenderScript(qb422016)
	//line errorView.qtpl:12
	qs422016 := string(qb422016.B)
	//line errorView.qtpl:12
	qt422016.ReleaseByteBuffer(qb422016)
	//line errorView.qtpl:12
	return qs422016
//line errorView.qtpl:12
}

//line errorView.qtpl:14
func (p *ErrorPage) StreamRenderBody(qw422016 *qt422016.Writer) {
	//line errorView.qtpl:14
	qw422016.N().S(`
	`)
	//line errorView.qtpl:16
	err := p.BasePage.BPD.Data["error"]

	//line errorView.qtpl:17
	qw422016.N().S(`
	<div class="text-center text-danger">`)
	//line errorView.qtpl:18
	qw422016.E().S(err.(string))
	//line errorView.qtpl:18
	qw422016.N().S(`</div>
`)
//line errorView.qtpl:19
}

//line errorView.qtpl:19
func (p *ErrorPage) WriteRenderBody(qq422016 qtio422016.Writer) {
	//line errorView.qtpl:19
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line errorView.qtpl:19
	p.StreamRenderBody(qw422016)
	//line errorView.qtpl:19
	qt422016.ReleaseWriter(qw422016)
//line errorView.qtpl:19
}

//line errorView.qtpl:19
func (p *ErrorPage) RenderBody() string {
	//line errorView.qtpl:19
	qb422016 := qt422016.AcquireByteBuffer()
	//line errorView.qtpl:19
	p.WriteRenderBody(qb422016)
	//line errorView.qtpl:19
	qs422016 := string(qb422016.B)
	//line errorView.qtpl:19
	qt422016.ReleaseByteBuffer(qb422016)
	//line errorView.qtpl:19
	return qs422016
//line errorView.qtpl:19
}