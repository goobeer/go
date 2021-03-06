// This file is automatically generated by qtc from "detail.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line admin\views\article\detail.qtpl:1
package article

//line admin\views\article\detail.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line admin\views\article\detail.qtpl:2
import (
	"fasthttpweb/area"
	"fasthttpweb/model"
)

//line admin\views\article\detail.qtpl:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line admin\views\article\detail.qtpl:9
type DetailPage struct {
	*area.BasePage
}

//line admin\views\article\detail.qtpl:14
func (p *DetailPage) StreamRenderTitle(qw422016 *qt422016.Writer) {
	//line admin\views\article\detail.qtpl:14
	p.BasePage.StreamRenderTitle(qw422016)
	//line admin\views\article\detail.qtpl:14
	qw422016.N().S(` `)
//line admin\views\article\detail.qtpl:14
}

//line admin\views\article\detail.qtpl:14
func (p *DetailPage) WriteRenderTitle(qq422016 qtio422016.Writer) {
	//line admin\views\article\detail.qtpl:14
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line admin\views\article\detail.qtpl:14
	p.StreamRenderTitle(qw422016)
	//line admin\views\article\detail.qtpl:14
	qt422016.ReleaseWriter(qw422016)
//line admin\views\article\detail.qtpl:14
}

//line admin\views\article\detail.qtpl:14
func (p *DetailPage) RenderTitle() string {
	//line admin\views\article\detail.qtpl:14
	qb422016 := qt422016.AcquireByteBuffer()
	//line admin\views\article\detail.qtpl:14
	p.WriteRenderTitle(qb422016)
	//line admin\views\article\detail.qtpl:14
	qs422016 := string(qb422016.B)
	//line admin\views\article\detail.qtpl:14
	qt422016.ReleaseByteBuffer(qb422016)
	//line admin\views\article\detail.qtpl:14
	return qs422016
//line admin\views\article\detail.qtpl:14
}

//line admin\views\article\detail.qtpl:15
func (p *DetailPage) StreamRenderKwd(qw422016 *qt422016.Writer) {
	//line admin\views\article\detail.qtpl:15
	p.BasePage.StreamRenderKwd(qw422016)
	//line admin\views\article\detail.qtpl:15
	qw422016.N().S(` `)
//line admin\views\article\detail.qtpl:15
}

//line admin\views\article\detail.qtpl:15
func (p *DetailPage) WriteRenderKwd(qq422016 qtio422016.Writer) {
	//line admin\views\article\detail.qtpl:15
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line admin\views\article\detail.qtpl:15
	p.StreamRenderKwd(qw422016)
	//line admin\views\article\detail.qtpl:15
	qt422016.ReleaseWriter(qw422016)
//line admin\views\article\detail.qtpl:15
}

//line admin\views\article\detail.qtpl:15
func (p *DetailPage) RenderKwd() string {
	//line admin\views\article\detail.qtpl:15
	qb422016 := qt422016.AcquireByteBuffer()
	//line admin\views\article\detail.qtpl:15
	p.WriteRenderKwd(qb422016)
	//line admin\views\article\detail.qtpl:15
	qs422016 := string(qb422016.B)
	//line admin\views\article\detail.qtpl:15
	qt422016.ReleaseByteBuffer(qb422016)
	//line admin\views\article\detail.qtpl:15
	return qs422016
//line admin\views\article\detail.qtpl:15
}

//line admin\views\article\detail.qtpl:17
func (p *DetailPage) StreamRenderCss(qw422016 *qt422016.Writer) {
//line admin\views\article\detail.qtpl:17
}

//line admin\views\article\detail.qtpl:17
func (p *DetailPage) WriteRenderCss(qq422016 qtio422016.Writer) {
	//line admin\views\article\detail.qtpl:17
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line admin\views\article\detail.qtpl:17
	p.StreamRenderCss(qw422016)
	//line admin\views\article\detail.qtpl:17
	qt422016.ReleaseWriter(qw422016)
//line admin\views\article\detail.qtpl:17
}

//line admin\views\article\detail.qtpl:17
func (p *DetailPage) RenderCss() string {
	//line admin\views\article\detail.qtpl:17
	qb422016 := qt422016.AcquireByteBuffer()
	//line admin\views\article\detail.qtpl:17
	p.WriteRenderCss(qb422016)
	//line admin\views\article\detail.qtpl:17
	qs422016 := string(qb422016.B)
	//line admin\views\article\detail.qtpl:17
	qt422016.ReleaseByteBuffer(qb422016)
	//line admin\views\article\detail.qtpl:17
	return qs422016
//line admin\views\article\detail.qtpl:17
}

//line admin\views\article\detail.qtpl:19
func (p *DetailPage) StreamRenderScript(qw422016 *qt422016.Writer) {
	//line admin\views\article\detail.qtpl:19
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
//line admin\views\article\detail.qtpl:28
}

//line admin\views\article\detail.qtpl:28
func (p *DetailPage) WriteRenderScript(qq422016 qtio422016.Writer) {
	//line admin\views\article\detail.qtpl:28
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line admin\views\article\detail.qtpl:28
	p.StreamRenderScript(qw422016)
	//line admin\views\article\detail.qtpl:28
	qt422016.ReleaseWriter(qw422016)
//line admin\views\article\detail.qtpl:28
}

//line admin\views\article\detail.qtpl:28
func (p *DetailPage) RenderScript() string {
	//line admin\views\article\detail.qtpl:28
	qb422016 := qt422016.AcquireByteBuffer()
	//line admin\views\article\detail.qtpl:28
	p.WriteRenderScript(qb422016)
	//line admin\views\article\detail.qtpl:28
	qs422016 := string(qb422016.B)
	//line admin\views\article\detail.qtpl:28
	qt422016.ReleaseByteBuffer(qb422016)
	//line admin\views\article\detail.qtpl:28
	return qs422016
//line admin\views\article\detail.qtpl:28
}

//line admin\views\article\detail.qtpl:31
func (p *DetailPage) StreamRenderBody(qw422016 *qt422016.Writer) {
	//line admin\views\article\detail.qtpl:31
	qw422016.N().S(`
	<div class="text-center">
		`)
	//line admin\views\article\detail.qtpl:34
	err := p.BasePage.BPD.Data["error"]
	art := p.BasePage.BPD.Data["Model"]
	var artVal *model.Article
	if art != nil {
		artVal = art.(*model.Article)
	}

	//line admin\views\article\detail.qtpl:40
	qw422016.N().S(`
		`)
	//line admin\views\article\detail.qtpl:41
	if err != nil {
		//line admin\views\article\detail.qtpl:41
		qw422016.N().S(`<div class="text-danger">`)
		//line admin\views\article\detail.qtpl:41
		qw422016.E().S(err.(string))
		//line admin\views\article\detail.qtpl:41
		qw422016.N().S(`</div>`)
		//line admin\views\article\detail.qtpl:41
	}
	//line admin\views\article\detail.qtpl:41
	qw422016.N().S(`
		`)
	//line admin\views\article\detail.qtpl:42
	if artVal != nil {
		//line admin\views\article\detail.qtpl:42
		qw422016.N().S(`
			<div>
				标题:`)
		//line admin\views\article\detail.qtpl:44
		qw422016.E().S(artVal.Title)
		//line admin\views\article\detail.qtpl:44
		qw422016.N().S(`
			</div>
			<div>
				创建时间:`)
		//line admin\views\article\detail.qtpl:47
		qw422016.E().S(artVal.CreateTime.Format("2006-01-02 15:04:05"))
		//line admin\views\article\detail.qtpl:47
		qw422016.N().S(`
			</div>
			<div>
				启用:`)
		//line admin\views\article\detail.qtpl:50
		if artVal.Used {
			//line admin\views\article\detail.qtpl:50
			qw422016.N().S(`是`)
			//line admin\views\article\detail.qtpl:50
		} else {
			//line admin\views\article\detail.qtpl:50
			qw422016.N().S(`否`)
			//line admin\views\article\detail.qtpl:50
		}
		//line admin\views\article\detail.qtpl:50
		qw422016.N().S(`
			</div>
			<div>
				`)
		//line admin\views\article\detail.qtpl:53
		qw422016.N().S(artVal.Content)
		//line admin\views\article\detail.qtpl:53
		qw422016.N().S(`
			</div>
		`)
		//line admin\views\article\detail.qtpl:55
	}
	//line admin\views\article\detail.qtpl:55
	qw422016.N().S(`
	</div>
`)
//line admin\views\article\detail.qtpl:57
}

//line admin\views\article\detail.qtpl:57
func (p *DetailPage) WriteRenderBody(qq422016 qtio422016.Writer) {
	//line admin\views\article\detail.qtpl:57
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line admin\views\article\detail.qtpl:57
	p.StreamRenderBody(qw422016)
	//line admin\views\article\detail.qtpl:57
	qt422016.ReleaseWriter(qw422016)
//line admin\views\article\detail.qtpl:57
}

//line admin\views\article\detail.qtpl:57
func (p *DetailPage) RenderBody() string {
	//line admin\views\article\detail.qtpl:57
	qb422016 := qt422016.AcquireByteBuffer()
	//line admin\views\article\detail.qtpl:57
	p.WriteRenderBody(qb422016)
	//line admin\views\article\detail.qtpl:57
	qs422016 := string(qb422016.B)
	//line admin\views\article\detail.qtpl:57
	qt422016.ReleaseByteBuffer(qb422016)
	//line admin\views\article\detail.qtpl:57
	return qs422016
//line admin\views\article\detail.qtpl:57
}
