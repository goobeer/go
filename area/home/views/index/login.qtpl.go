// This file is automatically generated by qtc from "login.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line home\views\index\login.qtpl:1
package index

//line home\views\index\login.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

// Main page template. Implements BasePage methods.
//

//line home\views\index\login.qtpl:3
import "fasthttpweb/area"

//line home\views\index\login.qtpl:5
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line home\views\index\login.qtpl:6
type LoginPage struct {
	*area.BasePage
}

//line home\views\index\login.qtpl:11
func (p *LoginPage) StreamRenderTitle(qw422016 *qt422016.Writer) {
	//line home\views\index\login.qtpl:11
	p.BasePage.StreamRenderTitle(qw422016)
//line home\views\index\login.qtpl:11
}

//line home\views\index\login.qtpl:11
func (p *LoginPage) WriteRenderTitle(qq422016 qtio422016.Writer) {
	//line home\views\index\login.qtpl:11
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line home\views\index\login.qtpl:11
	p.StreamRenderTitle(qw422016)
	//line home\views\index\login.qtpl:11
	qt422016.ReleaseWriter(qw422016)
//line home\views\index\login.qtpl:11
}

//line home\views\index\login.qtpl:11
func (p *LoginPage) RenderTitle() string {
	//line home\views\index\login.qtpl:11
	qb422016 := qt422016.AcquireByteBuffer()
	//line home\views\index\login.qtpl:11
	p.WriteRenderTitle(qb422016)
	//line home\views\index\login.qtpl:11
	qs422016 := string(qb422016.B)
	//line home\views\index\login.qtpl:11
	qt422016.ReleaseByteBuffer(qb422016)
	//line home\views\index\login.qtpl:11
	return qs422016
//line home\views\index\login.qtpl:11
}

//line home\views\index\login.qtpl:12
func (p *LoginPage) StreamRenderKwd(qw422016 *qt422016.Writer) {
//line home\views\index\login.qtpl:12
p.BasePage.StreamRenderKwd(qw422016) }

//line home\views\index\login.qtpl:12
//line home\views\index\login.qtpl:12
func (p *LoginPage) WriteRenderKwd(qq422016 qtio422016.Writer) {
	//line home\views\index\login.qtpl:12
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line home\views\index\login.qtpl:12
	p.StreamRenderKwd(qw422016)
	//line home\views\index\login.qtpl:12
	qt422016.ReleaseWriter(qw422016)
//line home\views\index\login.qtpl:12
}

//line home\views\index\login.qtpl:12
func (p *LoginPage) RenderKwd() string {
	//line home\views\index\login.qtpl:12
	qb422016 := qt422016.AcquireByteBuffer()
	//line home\views\index\login.qtpl:12
	p.WriteRenderKwd(qb422016)
	//line home\views\index\login.qtpl:12
	qs422016 := string(qb422016.B)
	//line home\views\index\login.qtpl:12
	qt422016.ReleaseByteBuffer(qb422016)
	//line home\views\index\login.qtpl:12
	return qs422016
//line home\views\index\login.qtpl:12
}

//line home\views\index\login.qtpl:14
func (p *LoginPage) StreamRenderCss(qw422016 *qt422016.Writer) {
	//line home\views\index\login.qtpl:14
	qw422016.N().S(`
`)
//line home\views\index\login.qtpl:15
}

//line home\views\index\login.qtpl:15
func (p *LoginPage) WriteRenderCss(qq422016 qtio422016.Writer) {
	//line home\views\index\login.qtpl:15
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line home\views\index\login.qtpl:15
	p.StreamRenderCss(qw422016)
	//line home\views\index\login.qtpl:15
	qt422016.ReleaseWriter(qw422016)
//line home\views\index\login.qtpl:15
}

//line home\views\index\login.qtpl:15
func (p *LoginPage) RenderCss() string {
	//line home\views\index\login.qtpl:15
	qb422016 := qt422016.AcquireByteBuffer()
	//line home\views\index\login.qtpl:15
	p.WriteRenderCss(qb422016)
	//line home\views\index\login.qtpl:15
	qs422016 := string(qb422016.B)
	//line home\views\index\login.qtpl:15
	qt422016.ReleaseByteBuffer(qb422016)
	//line home\views\index\login.qtpl:15
	return qs422016
//line home\views\index\login.qtpl:15
}

//line home\views\index\login.qtpl:17
func (p *LoginPage) StreamRenderScript(qw422016 *qt422016.Writer) {
	//line home\views\index\login.qtpl:17
	qw422016.N().S(`
`)
//line home\views\index\login.qtpl:18
}

//line home\views\index\login.qtpl:18
func (p *LoginPage) WriteRenderScript(qq422016 qtio422016.Writer) {
	//line home\views\index\login.qtpl:18
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line home\views\index\login.qtpl:18
	p.StreamRenderScript(qw422016)
	//line home\views\index\login.qtpl:18
	qt422016.ReleaseWriter(qw422016)
//line home\views\index\login.qtpl:18
}

//line home\views\index\login.qtpl:18
func (p *LoginPage) RenderScript() string {
	//line home\views\index\login.qtpl:18
	qb422016 := qt422016.AcquireByteBuffer()
	//line home\views\index\login.qtpl:18
	p.WriteRenderScript(qb422016)
	//line home\views\index\login.qtpl:18
	qs422016 := string(qb422016.B)
	//line home\views\index\login.qtpl:18
	qt422016.ReleaseByteBuffer(qb422016)
	//line home\views\index\login.qtpl:18
	return qs422016
//line home\views\index\login.qtpl:18
}

//line home\views\index\login.qtpl:20
func (p *LoginPage) StreamRenderBody(qw422016 *qt422016.Writer) {
	//line home\views\index\login.qtpl:20
	qw422016.N().S(`
	<div class="text-danger">
	`)
	//line home\views\index\login.qtpl:22
	if p.BPD.Data["ErrMsg"] != nil {
		//line home\views\index\login.qtpl:22
		qw422016.N().S(`
		`)
		//line home\views\index\login.qtpl:23
		qw422016.E().S((p.BPD.Data["ErrMsg"]).(string))
		//line home\views\index\login.qtpl:23
		qw422016.N().S(`
	`)
		//line home\views\index\login.qtpl:24
	}
	//line home\views\index\login.qtpl:24
	qw422016.N().S(`
	</div>
	<div class="text-center">
		<form class="form-horizontal" action="/home/index/login" method="POST">
			<div class="form-group">
				<label class="control-label col-md-2">用户名:</label>
				<div class="col-md-8">
					<input class="form-control" id="uname" name="uname" />
				</div>
			</div>
			<div class="form-group">
				<label class="control-label col-md-2">密码:</label>
				<div class="col-md-8">
					<input type="password" class="form-control" id="pwd" name="pwd" />
				</div>
			</div>
			<div class="form-group">
				<label class="control-label col-md-2">验证码:</label>
				<div class="col-md-8">
					<input class="form-control" id="vc" name="vcode" />
				</div>
				<div class="col-md-2">
					<img id="vcode" src="/api/verifyimg" title="点击切换验证码" />
				</div>
			</div>
			<div class="form-group">
				<div class="col-md-offset-2 col-md-10">
					<button type="submit" class="btn btn-default">登录</button>
				</div>
			</div>
		</form>
	</div>
`)
//line home\views\index\login.qtpl:56
}

//line home\views\index\login.qtpl:56
func (p *LoginPage) WriteRenderBody(qq422016 qtio422016.Writer) {
	//line home\views\index\login.qtpl:56
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line home\views\index\login.qtpl:56
	p.StreamRenderBody(qw422016)
	//line home\views\index\login.qtpl:56
	qt422016.ReleaseWriter(qw422016)
//line home\views\index\login.qtpl:56
}

//line home\views\index\login.qtpl:56
func (p *LoginPage) RenderBody() string {
	//line home\views\index\login.qtpl:56
	qb422016 := qt422016.AcquireByteBuffer()
	//line home\views\index\login.qtpl:56
	p.WriteRenderBody(qb422016)
	//line home\views\index\login.qtpl:56
	qs422016 := string(qb422016.B)
	//line home\views\index\login.qtpl:56
	qt422016.ReleaseByteBuffer(qb422016)
	//line home\views\index\login.qtpl:56
	return qs422016
//line home\views\index\login.qtpl:56
}
