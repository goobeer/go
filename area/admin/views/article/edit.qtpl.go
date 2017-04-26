// This file is automatically generated by qtc from "edit.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line admin\views\article\edit.qtpl:1
package article

//line admin\views\article\edit.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line admin\views\article\edit.qtpl:2
import (
	"fasthttpweb/area"
	"fasthttpweb/model"
)

//line admin\views\article\edit.qtpl:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line admin\views\article\edit.qtpl:9
type EditPage struct {
	*area.BasePage
}

//line admin\views\article\edit.qtpl:14
func (p *EditPage) StreamRenderTitle(qw422016 *qt422016.Writer) {
	//line admin\views\article\edit.qtpl:14
	p.BasePage.StreamRenderTitle(qw422016)
	//line admin\views\article\edit.qtpl:14
	qw422016.N().S(` `)
//line admin\views\article\edit.qtpl:14
}

//line admin\views\article\edit.qtpl:14
func (p *EditPage) WriteRenderTitle(qq422016 qtio422016.Writer) {
	//line admin\views\article\edit.qtpl:14
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line admin\views\article\edit.qtpl:14
	p.StreamRenderTitle(qw422016)
	//line admin\views\article\edit.qtpl:14
	qt422016.ReleaseWriter(qw422016)
//line admin\views\article\edit.qtpl:14
}

//line admin\views\article\edit.qtpl:14
func (p *EditPage) RenderTitle() string {
	//line admin\views\article\edit.qtpl:14
	qb422016 := qt422016.AcquireByteBuffer()
	//line admin\views\article\edit.qtpl:14
	p.WriteRenderTitle(qb422016)
	//line admin\views\article\edit.qtpl:14
	qs422016 := string(qb422016.B)
	//line admin\views\article\edit.qtpl:14
	qt422016.ReleaseByteBuffer(qb422016)
	//line admin\views\article\edit.qtpl:14
	return qs422016
//line admin\views\article\edit.qtpl:14
}

//line admin\views\article\edit.qtpl:15
func (p *EditPage) StreamRenderKwd(qw422016 *qt422016.Writer) {
	//line admin\views\article\edit.qtpl:15
	p.BasePage.StreamRenderKwd(qw422016)
	//line admin\views\article\edit.qtpl:15
	qw422016.N().S(` `)
//line admin\views\article\edit.qtpl:15
}

//line admin\views\article\edit.qtpl:15
func (p *EditPage) WriteRenderKwd(qq422016 qtio422016.Writer) {
	//line admin\views\article\edit.qtpl:15
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line admin\views\article\edit.qtpl:15
	p.StreamRenderKwd(qw422016)
	//line admin\views\article\edit.qtpl:15
	qt422016.ReleaseWriter(qw422016)
//line admin\views\article\edit.qtpl:15
}

//line admin\views\article\edit.qtpl:15
func (p *EditPage) RenderKwd() string {
	//line admin\views\article\edit.qtpl:15
	qb422016 := qt422016.AcquireByteBuffer()
	//line admin\views\article\edit.qtpl:15
	p.WriteRenderKwd(qb422016)
	//line admin\views\article\edit.qtpl:15
	qs422016 := string(qb422016.B)
	//line admin\views\article\edit.qtpl:15
	qt422016.ReleaseByteBuffer(qb422016)
	//line admin\views\article\edit.qtpl:15
	return qs422016
//line admin\views\article\edit.qtpl:15
}

//line admin\views\article\edit.qtpl:17
func (p *EditPage) StreamRenderCss(qw422016 *qt422016.Writer) {
//line admin\views\article\edit.qtpl:17
}

//line admin\views\article\edit.qtpl:17
func (p *EditPage) WriteRenderCss(qq422016 qtio422016.Writer) {
	//line admin\views\article\edit.qtpl:17
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line admin\views\article\edit.qtpl:17
	p.StreamRenderCss(qw422016)
	//line admin\views\article\edit.qtpl:17
	qt422016.ReleaseWriter(qw422016)
//line admin\views\article\edit.qtpl:17
}

//line admin\views\article\edit.qtpl:17
func (p *EditPage) RenderCss() string {
	//line admin\views\article\edit.qtpl:17
	qb422016 := qt422016.AcquireByteBuffer()
	//line admin\views\article\edit.qtpl:17
	p.WriteRenderCss(qb422016)
	//line admin\views\article\edit.qtpl:17
	qs422016 := string(qb422016.B)
	//line admin\views\article\edit.qtpl:17
	qt422016.ReleaseByteBuffer(qb422016)
	//line admin\views\article\edit.qtpl:17
	return qs422016
//line admin\views\article\edit.qtpl:17
}

//line admin\views\article\edit.qtpl:19
func (p *EditPage) StreamRenderScript(qw422016 *qt422016.Writer) {
	//line admin\views\article\edit.qtpl:19
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
//line admin\views\article\edit.qtpl:28
}

//line admin\views\article\edit.qtpl:28
func (p *EditPage) WriteRenderScript(qq422016 qtio422016.Writer) {
	//line admin\views\article\edit.qtpl:28
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line admin\views\article\edit.qtpl:28
	p.StreamRenderScript(qw422016)
	//line admin\views\article\edit.qtpl:28
	qt422016.ReleaseWriter(qw422016)
//line admin\views\article\edit.qtpl:28
}

//line admin\views\article\edit.qtpl:28
func (p *EditPage) RenderScript() string {
	//line admin\views\article\edit.qtpl:28
	qb422016 := qt422016.AcquireByteBuffer()
	//line admin\views\article\edit.qtpl:28
	p.WriteRenderScript(qb422016)
	//line admin\views\article\edit.qtpl:28
	qs422016 := string(qb422016.B)
	//line admin\views\article\edit.qtpl:28
	qt422016.ReleaseByteBuffer(qb422016)
	//line admin\views\article\edit.qtpl:28
	return qs422016
//line admin\views\article\edit.qtpl:28
}

//line admin\views\article\edit.qtpl:31
func (p *EditPage) StreamRenderBody(qw422016 *qt422016.Writer) {
	//line admin\views\article\edit.qtpl:31
	qw422016.N().S(`
	`)
	//line admin\views\article\edit.qtpl:33
	art := p.BasePage.BPD.Data["Model"]
	var artVal *model.Article
	if art != nil {
		artVal = art.(*model.Article)
	}

	//line admin\views\article\edit.qtpl:38
	qw422016.N().S(`
	`)
	//line admin\views\article\edit.qtpl:39
	if artVal != nil {
		//line admin\views\article\edit.qtpl:39
		qw422016.N().S(`
	<div class="text-center">
		<form class="form-horizontal" action="`)
		//line admin\views\article\edit.qtpl:41
		qw422016.E().S(area.Url(p.BasePage, "edit", nil))
		//line admin\views\article\edit.qtpl:41
		qw422016.N().S(`" method="POST">
			<input type="hidden" name="ID" value="`)
		//line admin\views\article\edit.qtpl:42
		qw422016.N().D(int(artVal.ID))
		//line admin\views\article\edit.qtpl:42
		qw422016.N().S(`" />
			<div class="form-group">
				<label class="control-label col-md-2">标题：</label>
				<div class="col-md-10">
					<input class="form-control" id="Title" name="Title" value="`)
		//line admin\views\article\edit.qtpl:46
		qw422016.E().S(artVal.Title)
		//line admin\views\article\edit.qtpl:46
		qw422016.N().S(`"  />
				</div>
			</div>
			<div class="form-group">
				<label class="control-label col-md-2">内容：</label>
				<div class="col-md-10">
					<textarea name="Content" id="Content">`)
		//line admin\views\article\edit.qtpl:52
		qw422016.N().S(artVal.Content)
		//line admin\views\article\edit.qtpl:52
		qw422016.N().S(`</textarea>
				</div>
			</div>
			<div class="form-group">
				<div class="col-md-offset-2 col-md-10">
					<button type="submit" class="btn btn-default">保存</button>
				</div>
			</div>
		</form>
	</div>
	`)
		//line admin\views\article\edit.qtpl:62
	}
	//line admin\views\article\edit.qtpl:62
	qw422016.N().S(`
`)
//line admin\views\article\edit.qtpl:63
}

//line admin\views\article\edit.qtpl:63
func (p *EditPage) WriteRenderBody(qq422016 qtio422016.Writer) {
	//line admin\views\article\edit.qtpl:63
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line admin\views\article\edit.qtpl:63
	p.StreamRenderBody(qw422016)
	//line admin\views\article\edit.qtpl:63
	qt422016.ReleaseWriter(qw422016)
//line admin\views\article\edit.qtpl:63
}

//line admin\views\article\edit.qtpl:63
func (p *EditPage) RenderBody() string {
	//line admin\views\article\edit.qtpl:63
	qb422016 := qt422016.AcquireByteBuffer()
	//line admin\views\article\edit.qtpl:63
	p.WriteRenderBody(qb422016)
	//line admin\views\article\edit.qtpl:63
	qs422016 := string(qb422016.B)
	//line admin\views\article\edit.qtpl:63
	qt422016.ReleaseByteBuffer(qb422016)
	//line admin\views\article\edit.qtpl:63
	return qs422016
//line admin\views\article\edit.qtpl:63
}
