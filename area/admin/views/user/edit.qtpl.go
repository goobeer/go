// This file is automatically generated by qtc from "edit.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line admin\views\user\edit.qtpl:1
package user

//line admin\views\user\edit.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line admin\views\user\edit.qtpl:2
import (
	"fasthttpweb/area"
	"fasthttpweb/model"
)

//line admin\views\user\edit.qtpl:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line admin\views\user\edit.qtpl:9
type EditPage struct {
	*area.BasePage
}

//line admin\views\user\edit.qtpl:14
func (p *EditPage) StreamRenderTitle(qw422016 *qt422016.Writer) {
	//line admin\views\user\edit.qtpl:14
	p.BasePage.StreamRenderTitle(qw422016)
	//line admin\views\user\edit.qtpl:14
	qw422016.N().S(` `)
//line admin\views\user\edit.qtpl:14
}

//line admin\views\user\edit.qtpl:14
func (p *EditPage) WriteRenderTitle(qq422016 qtio422016.Writer) {
	//line admin\views\user\edit.qtpl:14
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line admin\views\user\edit.qtpl:14
	p.StreamRenderTitle(qw422016)
	//line admin\views\user\edit.qtpl:14
	qt422016.ReleaseWriter(qw422016)
//line admin\views\user\edit.qtpl:14
}

//line admin\views\user\edit.qtpl:14
func (p *EditPage) RenderTitle() string {
	//line admin\views\user\edit.qtpl:14
	qb422016 := qt422016.AcquireByteBuffer()
	//line admin\views\user\edit.qtpl:14
	p.WriteRenderTitle(qb422016)
	//line admin\views\user\edit.qtpl:14
	qs422016 := string(qb422016.B)
	//line admin\views\user\edit.qtpl:14
	qt422016.ReleaseByteBuffer(qb422016)
	//line admin\views\user\edit.qtpl:14
	return qs422016
//line admin\views\user\edit.qtpl:14
}

//line admin\views\user\edit.qtpl:15
func (p *EditPage) StreamRenderKwd(qw422016 *qt422016.Writer) {
	//line admin\views\user\edit.qtpl:15
	p.BasePage.StreamRenderKwd(qw422016)
	//line admin\views\user\edit.qtpl:15
	qw422016.N().S(` `)
//line admin\views\user\edit.qtpl:15
}

//line admin\views\user\edit.qtpl:15
func (p *EditPage) WriteRenderKwd(qq422016 qtio422016.Writer) {
	//line admin\views\user\edit.qtpl:15
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line admin\views\user\edit.qtpl:15
	p.StreamRenderKwd(qw422016)
	//line admin\views\user\edit.qtpl:15
	qt422016.ReleaseWriter(qw422016)
//line admin\views\user\edit.qtpl:15
}

//line admin\views\user\edit.qtpl:15
func (p *EditPage) RenderKwd() string {
	//line admin\views\user\edit.qtpl:15
	qb422016 := qt422016.AcquireByteBuffer()
	//line admin\views\user\edit.qtpl:15
	p.WriteRenderKwd(qb422016)
	//line admin\views\user\edit.qtpl:15
	qs422016 := string(qb422016.B)
	//line admin\views\user\edit.qtpl:15
	qt422016.ReleaseByteBuffer(qb422016)
	//line admin\views\user\edit.qtpl:15
	return qs422016
//line admin\views\user\edit.qtpl:15
}

//line admin\views\user\edit.qtpl:17
func (p *EditPage) StreamRenderCss(qw422016 *qt422016.Writer) {
//line admin\views\user\edit.qtpl:17
}

//line admin\views\user\edit.qtpl:17
func (p *EditPage) WriteRenderCss(qq422016 qtio422016.Writer) {
	//line admin\views\user\edit.qtpl:17
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line admin\views\user\edit.qtpl:17
	p.StreamRenderCss(qw422016)
	//line admin\views\user\edit.qtpl:17
	qt422016.ReleaseWriter(qw422016)
//line admin\views\user\edit.qtpl:17
}

//line admin\views\user\edit.qtpl:17
func (p *EditPage) RenderCss() string {
	//line admin\views\user\edit.qtpl:17
	qb422016 := qt422016.AcquireByteBuffer()
	//line admin\views\user\edit.qtpl:17
	p.WriteRenderCss(qb422016)
	//line admin\views\user\edit.qtpl:17
	qs422016 := string(qb422016.B)
	//line admin\views\user\edit.qtpl:17
	qt422016.ReleaseByteBuffer(qb422016)
	//line admin\views\user\edit.qtpl:17
	return qs422016
//line admin\views\user\edit.qtpl:17
}

//line admin\views\user\edit.qtpl:19
func (p *EditPage) StreamRenderScript(qw422016 *qt422016.Writer) {
//line admin\views\user\edit.qtpl:19
}

//line admin\views\user\edit.qtpl:19
func (p *EditPage) WriteRenderScript(qq422016 qtio422016.Writer) {
	//line admin\views\user\edit.qtpl:19
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line admin\views\user\edit.qtpl:19
	p.StreamRenderScript(qw422016)
	//line admin\views\user\edit.qtpl:19
	qt422016.ReleaseWriter(qw422016)
//line admin\views\user\edit.qtpl:19
}

//line admin\views\user\edit.qtpl:19
func (p *EditPage) RenderScript() string {
	//line admin\views\user\edit.qtpl:19
	qb422016 := qt422016.AcquireByteBuffer()
	//line admin\views\user\edit.qtpl:19
	p.WriteRenderScript(qb422016)
	//line admin\views\user\edit.qtpl:19
	qs422016 := string(qb422016.B)
	//line admin\views\user\edit.qtpl:19
	qt422016.ReleaseByteBuffer(qb422016)
	//line admin\views\user\edit.qtpl:19
	return qs422016
//line admin\views\user\edit.qtpl:19
}

//line admin\views\user\edit.qtpl:21
func (p *EditPage) StreamRenderBody(qw422016 *qt422016.Writer) {
	//line admin\views\user\edit.qtpl:21
	qw422016.N().S(`
	`)
	//line admin\views\user\edit.qtpl:23
	modelVal := p.BasePage.BPD.Data["Model"]
	var user *model.Users
	if modelVal != nil {
		user = modelVal.(*model.Users)
	}

	//line admin\views\user\edit.qtpl:28
	qw422016.N().S(`
	<div class="text-center">
		`)
	//line admin\views\user\edit.qtpl:30
	if user != nil {
		//line admin\views\user\edit.qtpl:30
		qw422016.N().S(`
			<form class="form-horizontal" action="`)
		//line admin\views\user\edit.qtpl:31
		qw422016.E().S(area.Url(p.BasePage, "edit", nil))
		//line admin\views\user\edit.qtpl:31
		qw422016.N().S(`" method="POST">
			<input type="hidden" name="ID" value="`)
		//line admin\views\user\edit.qtpl:32
		qw422016.N().D(int(user.ID))
		//line admin\views\user\edit.qtpl:32
		qw422016.N().S(`" />
			<div class="form-group">
				<label class="control-label col-md-2">密码：</label>
				<div class="col-md-10">
					<input class="form-control" id="Pwd" name="Pwd"  />
				</div>
			</div>
			<div class="form-group">
				<label class="control-label col-md-2">重复输入密码：</label>
				<div class="col-md-10">
					<input class="form-control" id="rePwd" />
				</div>
			</div>
			<div class="form-group">
				<div class="col-md-offset-2 col-md-10">
					<button type="submit" class="btn btn-default">添加</button>
				</div>
			</div>
			</form>
		`)
		//line admin\views\user\edit.qtpl:51
	}
	//line admin\views\user\edit.qtpl:51
	qw422016.N().S(`
	</div>
`)
//line admin\views\user\edit.qtpl:53
}

//line admin\views\user\edit.qtpl:53
func (p *EditPage) WriteRenderBody(qq422016 qtio422016.Writer) {
	//line admin\views\user\edit.qtpl:53
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line admin\views\user\edit.qtpl:53
	p.StreamRenderBody(qw422016)
	//line admin\views\user\edit.qtpl:53
	qt422016.ReleaseWriter(qw422016)
//line admin\views\user\edit.qtpl:53
}

//line admin\views\user\edit.qtpl:53
func (p *EditPage) RenderBody() string {
	//line admin\views\user\edit.qtpl:53
	qb422016 := qt422016.AcquireByteBuffer()
	//line admin\views\user\edit.qtpl:53
	p.WriteRenderBody(qb422016)
	//line admin\views\user\edit.qtpl:53
	qs422016 := string(qb422016.B)
	//line admin\views\user\edit.qtpl:53
	qt422016.ReleaseByteBuffer(qb422016)
	//line admin\views\user\edit.qtpl:53
	return qs422016
//line admin\views\user\edit.qtpl:53
}
