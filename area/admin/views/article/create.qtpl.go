// This file is automatically generated by qtc from "create.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line admin\views\article\create.qtpl:1
package article

//line admin\views\article\create.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

// Main page template. Implements BasePage methods.
//

//line admin\views\article\create.qtpl:4
import (
	"fasthttpweb/area"
)

//line admin\views\article\create.qtpl:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line admin\views\article\create.qtpl:10
type CreatePage struct {
	*area.BasePage
}

//line admin\views\article\create.qtpl:15
func (p *CreatePage) StreamRenderTitle(qw422016 *qt422016.Writer) {
	//line admin\views\article\create.qtpl:15
	p.BasePage.StreamRenderTitle(qw422016)
	//line admin\views\article\create.qtpl:15
	qw422016.N().S(` `)
//line admin\views\article\create.qtpl:15
}

//line admin\views\article\create.qtpl:15
func (p *CreatePage) WriteRenderTitle(qq422016 qtio422016.Writer) {
	//line admin\views\article\create.qtpl:15
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line admin\views\article\create.qtpl:15
	p.StreamRenderTitle(qw422016)
	//line admin\views\article\create.qtpl:15
	qt422016.ReleaseWriter(qw422016)
//line admin\views\article\create.qtpl:15
}

//line admin\views\article\create.qtpl:15
func (p *CreatePage) RenderTitle() string {
	//line admin\views\article\create.qtpl:15
	qb422016 := qt422016.AcquireByteBuffer()
	//line admin\views\article\create.qtpl:15
	p.WriteRenderTitle(qb422016)
	//line admin\views\article\create.qtpl:15
	qs422016 := string(qb422016.B)
	//line admin\views\article\create.qtpl:15
	qt422016.ReleaseByteBuffer(qb422016)
	//line admin\views\article\create.qtpl:15
	return qs422016
//line admin\views\article\create.qtpl:15
}

//line admin\views\article\create.qtpl:16
func (p *CreatePage) StreamRenderKwd(qw422016 *qt422016.Writer) {
	//line admin\views\article\create.qtpl:16
	p.BasePage.StreamRenderKwd(qw422016)
	//line admin\views\article\create.qtpl:16
	qw422016.N().S(` `)
//line admin\views\article\create.qtpl:16
}

//line admin\views\article\create.qtpl:16
func (p *CreatePage) WriteRenderKwd(qq422016 qtio422016.Writer) {
	//line admin\views\article\create.qtpl:16
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line admin\views\article\create.qtpl:16
	p.StreamRenderKwd(qw422016)
	//line admin\views\article\create.qtpl:16
	qt422016.ReleaseWriter(qw422016)
//line admin\views\article\create.qtpl:16
}

//line admin\views\article\create.qtpl:16
func (p *CreatePage) RenderKwd() string {
	//line admin\views\article\create.qtpl:16
	qb422016 := qt422016.AcquireByteBuffer()
	//line admin\views\article\create.qtpl:16
	p.WriteRenderKwd(qb422016)
	//line admin\views\article\create.qtpl:16
	qs422016 := string(qb422016.B)
	//line admin\views\article\create.qtpl:16
	qt422016.ReleaseByteBuffer(qb422016)
	//line admin\views\article\create.qtpl:16
	return qs422016
//line admin\views\article\create.qtpl:16
}

//line admin\views\article\create.qtpl:18
func (p *CreatePage) StreamRenderCss(qw422016 *qt422016.Writer) {
//line admin\views\article\create.qtpl:18
}

//line admin\views\article\create.qtpl:18
func (p *CreatePage) WriteRenderCss(qq422016 qtio422016.Writer) {
	//line admin\views\article\create.qtpl:18
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line admin\views\article\create.qtpl:18
	p.StreamRenderCss(qw422016)
	//line admin\views\article\create.qtpl:18
	qt422016.ReleaseWriter(qw422016)
//line admin\views\article\create.qtpl:18
}

//line admin\views\article\create.qtpl:18
func (p *CreatePage) RenderCss() string {
	//line admin\views\article\create.qtpl:18
	qb422016 := qt422016.AcquireByteBuffer()
	//line admin\views\article\create.qtpl:18
	p.WriteRenderCss(qb422016)
	//line admin\views\article\create.qtpl:18
	qs422016 := string(qb422016.B)
	//line admin\views\article\create.qtpl:18
	qt422016.ReleaseByteBuffer(qb422016)
	//line admin\views\article\create.qtpl:18
	return qs422016
//line admin\views\article\create.qtpl:18
}

//line admin\views\article\create.qtpl:20
func (p *CreatePage) StreamRenderScript(qw422016 *qt422016.Writer) {
	//line admin\views\article\create.qtpl:20
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
//line admin\views\article\create.qtpl:29
}

//line admin\views\article\create.qtpl:29
func (p *CreatePage) WriteRenderScript(qq422016 qtio422016.Writer) {
	//line admin\views\article\create.qtpl:29
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line admin\views\article\create.qtpl:29
	p.StreamRenderScript(qw422016)
	//line admin\views\article\create.qtpl:29
	qt422016.ReleaseWriter(qw422016)
//line admin\views\article\create.qtpl:29
}

//line admin\views\article\create.qtpl:29
func (p *CreatePage) RenderScript() string {
	//line admin\views\article\create.qtpl:29
	qb422016 := qt422016.AcquireByteBuffer()
	//line admin\views\article\create.qtpl:29
	p.WriteRenderScript(qb422016)
	//line admin\views\article\create.qtpl:29
	qs422016 := string(qb422016.B)
	//line admin\views\article\create.qtpl:29
	qt422016.ReleaseByteBuffer(qb422016)
	//line admin\views\article\create.qtpl:29
	return qs422016
//line admin\views\article\create.qtpl:29
}

//line admin\views\article\create.qtpl:32
func (p *CreatePage) StreamRenderBody(qw422016 *qt422016.Writer) {
	//line admin\views\article\create.qtpl:32
	qw422016.N().S(`
	<div class="text-center">
		<form class="form-horizontal" action="`)
	//line admin\views\article\create.qtpl:34
	qw422016.E().S(area.Url(p.BasePage, "create", nil))
	//line admin\views\article\create.qtpl:34
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
//line admin\views\article\create.qtpl:54
}

//line admin\views\article\create.qtpl:54
func (p *CreatePage) WriteRenderBody(qq422016 qtio422016.Writer) {
	//line admin\views\article\create.qtpl:54
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line admin\views\article\create.qtpl:54
	p.StreamRenderBody(qw422016)
	//line admin\views\article\create.qtpl:54
	qt422016.ReleaseWriter(qw422016)
//line admin\views\article\create.qtpl:54
}

//line admin\views\article\create.qtpl:54
func (p *CreatePage) RenderBody() string {
	//line admin\views\article\create.qtpl:54
	qb422016 := qt422016.AcquireByteBuffer()
	//line admin\views\article\create.qtpl:54
	p.WriteRenderBody(qb422016)
	//line admin\views\article\create.qtpl:54
	qs422016 := string(qb422016.B)
	//line admin\views\article\create.qtpl:54
	qt422016.ReleaseByteBuffer(qb422016)
	//line admin\views\article\create.qtpl:54
	return qs422016
//line admin\views\article\create.qtpl:54
}
