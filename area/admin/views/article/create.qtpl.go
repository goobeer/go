// This file is automatically generated by qtc from "create.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line admin\views\article\create.qtpl:1
package article

//line admin\views\article\create.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line admin\views\article\create.qtpl:2
import (
	"fasthttpweb/area"
)

//line admin\views\article\create.qtpl:7
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line admin\views\article\create.qtpl:8
type CreatePage struct {
	*area.BasePage
}

//line admin\views\article\create.qtpl:13
func (p *CreatePage) StreamRenderTitle(qw422016 *qt422016.Writer) {
	//line admin\views\article\create.qtpl:13
	p.BasePage.StreamRenderTitle(qw422016)
	//line admin\views\article\create.qtpl:13
	qw422016.N().S(` `)
//line admin\views\article\create.qtpl:13
}

//line admin\views\article\create.qtpl:13
func (p *CreatePage) WriteRenderTitle(qq422016 qtio422016.Writer) {
	//line admin\views\article\create.qtpl:13
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line admin\views\article\create.qtpl:13
	p.StreamRenderTitle(qw422016)
	//line admin\views\article\create.qtpl:13
	qt422016.ReleaseWriter(qw422016)
//line admin\views\article\create.qtpl:13
}

//line admin\views\article\create.qtpl:13
func (p *CreatePage) RenderTitle() string {
	//line admin\views\article\create.qtpl:13
	qb422016 := qt422016.AcquireByteBuffer()
	//line admin\views\article\create.qtpl:13
	p.WriteRenderTitle(qb422016)
	//line admin\views\article\create.qtpl:13
	qs422016 := string(qb422016.B)
	//line admin\views\article\create.qtpl:13
	qt422016.ReleaseByteBuffer(qb422016)
	//line admin\views\article\create.qtpl:13
	return qs422016
//line admin\views\article\create.qtpl:13
}

//line admin\views\article\create.qtpl:14
func (p *CreatePage) StreamRenderKwd(qw422016 *qt422016.Writer) {
	//line admin\views\article\create.qtpl:14
	p.BasePage.StreamRenderKwd(qw422016)
	//line admin\views\article\create.qtpl:14
	qw422016.N().S(` `)
//line admin\views\article\create.qtpl:14
}

//line admin\views\article\create.qtpl:14
func (p *CreatePage) WriteRenderKwd(qq422016 qtio422016.Writer) {
	//line admin\views\article\create.qtpl:14
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line admin\views\article\create.qtpl:14
	p.StreamRenderKwd(qw422016)
	//line admin\views\article\create.qtpl:14
	qt422016.ReleaseWriter(qw422016)
//line admin\views\article\create.qtpl:14
}

//line admin\views\article\create.qtpl:14
func (p *CreatePage) RenderKwd() string {
	//line admin\views\article\create.qtpl:14
	qb422016 := qt422016.AcquireByteBuffer()
	//line admin\views\article\create.qtpl:14
	p.WriteRenderKwd(qb422016)
	//line admin\views\article\create.qtpl:14
	qs422016 := string(qb422016.B)
	//line admin\views\article\create.qtpl:14
	qt422016.ReleaseByteBuffer(qb422016)
	//line admin\views\article\create.qtpl:14
	return qs422016
//line admin\views\article\create.qtpl:14
}

//line admin\views\article\create.qtpl:16
func (p *CreatePage) StreamRenderCss(qw422016 *qt422016.Writer) {
//line admin\views\article\create.qtpl:16
}

//line admin\views\article\create.qtpl:16
func (p *CreatePage) WriteRenderCss(qq422016 qtio422016.Writer) {
	//line admin\views\article\create.qtpl:16
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line admin\views\article\create.qtpl:16
	p.StreamRenderCss(qw422016)
	//line admin\views\article\create.qtpl:16
	qt422016.ReleaseWriter(qw422016)
//line admin\views\article\create.qtpl:16
}

//line admin\views\article\create.qtpl:16
func (p *CreatePage) RenderCss() string {
	//line admin\views\article\create.qtpl:16
	qb422016 := qt422016.AcquireByteBuffer()
	//line admin\views\article\create.qtpl:16
	p.WriteRenderCss(qb422016)
	//line admin\views\article\create.qtpl:16
	qs422016 := string(qb422016.B)
	//line admin\views\article\create.qtpl:16
	qt422016.ReleaseByteBuffer(qb422016)
	//line admin\views\article\create.qtpl:16
	return qs422016
//line admin\views\article\create.qtpl:16
}

//line admin\views\article\create.qtpl:18
func (p *CreatePage) StreamRenderScript(qw422016 *qt422016.Writer) {
	//line admin\views\article\create.qtpl:18
	qw422016.N().S(`
<script type="text/javascript" src="/public/ueditor/ueditor.config.js"></script>
<script type="text/javascript" src="/public/ueditor/ueditor.all.js"></script>
<script type="text/javascript">
	$(function(){
		var editor = new baidu.editor.ui.Editor({initialFrameWidth: 784,initialFrameHeight: 400});
    	editor.render('Content');
	});
</script>
`)
//line admin\views\article\create.qtpl:27
}

//line admin\views\article\create.qtpl:27
func (p *CreatePage) WriteRenderScript(qq422016 qtio422016.Writer) {
	//line admin\views\article\create.qtpl:27
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line admin\views\article\create.qtpl:27
	p.StreamRenderScript(qw422016)
	//line admin\views\article\create.qtpl:27
	qt422016.ReleaseWriter(qw422016)
//line admin\views\article\create.qtpl:27
}

//line admin\views\article\create.qtpl:27
func (p *CreatePage) RenderScript() string {
	//line admin\views\article\create.qtpl:27
	qb422016 := qt422016.AcquireByteBuffer()
	//line admin\views\article\create.qtpl:27
	p.WriteRenderScript(qb422016)
	//line admin\views\article\create.qtpl:27
	qs422016 := string(qb422016.B)
	//line admin\views\article\create.qtpl:27
	qt422016.ReleaseByteBuffer(qb422016)
	//line admin\views\article\create.qtpl:27
	return qs422016
//line admin\views\article\create.qtpl:27
}

//line admin\views\article\create.qtpl:30
func (p *CreatePage) StreamRenderBody(qw422016 *qt422016.Writer) {
	//line admin\views\article\create.qtpl:30
	qw422016.N().S(`
	<div class="text-center">
		<form class="form-horizontal" action="`)
	//line admin\views\article\create.qtpl:32
	qw422016.E().S(area.Url(p.BasePage, "create", nil))
	//line admin\views\article\create.qtpl:32
	qw422016.N().S(`" method="POST">
			<div class="form-group">
				<label class="control-label col-md-2">标题：</label>
				<div class="col-md-10">
					<input class="form-control" id="Title" name="Title"  />
				</div>
			</div>
			<div class="form-group">
				<label class="control-label col-md-2">内容：</label>
				<div class="col-md-10">
					<textarea name="Content" id="Content"></textarea>
				</div>
			</div>
			<div class="form-group">
				<div class="col-md-offset-2 col-md-10">
					<button type="submit" class="btn btn-default">添加</button>
				</div>
			</div>
		</form>
	</div>
`)
//line admin\views\article\create.qtpl:52
}

//line admin\views\article\create.qtpl:52
func (p *CreatePage) WriteRenderBody(qq422016 qtio422016.Writer) {
	//line admin\views\article\create.qtpl:52
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line admin\views\article\create.qtpl:52
	p.StreamRenderBody(qw422016)
	//line admin\views\article\create.qtpl:52
	qt422016.ReleaseWriter(qw422016)
//line admin\views\article\create.qtpl:52
}

//line admin\views\article\create.qtpl:52
func (p *CreatePage) RenderBody() string {
	//line admin\views\article\create.qtpl:52
	qb422016 := qt422016.AcquireByteBuffer()
	//line admin\views\article\create.qtpl:52
	p.WriteRenderBody(qb422016)
	//line admin\views\article\create.qtpl:52
	qs422016 := string(qb422016.B)
	//line admin\views\article\create.qtpl:52
	qt422016.ReleaseByteBuffer(qb422016)
	//line admin\views\article\create.qtpl:52
	return qs422016
//line admin\views\article\create.qtpl:52
}
