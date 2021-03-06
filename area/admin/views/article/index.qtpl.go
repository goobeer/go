// This file is automatically generated by qtc from "index.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line admin\views\article\index.qtpl:1
package article

//line admin\views\article\index.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

// Main page template. Implements BasePage methods.
//

//line admin\views\article\index.qtpl:4
import (
	"fasthttpweb/area"
	"fasthttpweb/area/commonhtmlhelper"
	"fasthttpweb/model"
)

//line admin\views\article\index.qtpl:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line admin\views\article\index.qtpl:12
type IndexPage struct {
	*area.BasePage
}

//line admin\views\article\index.qtpl:17
func (p *IndexPage) StreamRenderTitle(qw422016 *qt422016.Writer) {
	//line admin\views\article\index.qtpl:17
	p.BasePage.StreamRenderTitle(qw422016)
	//line admin\views\article\index.qtpl:17
	qw422016.N().S(` `)
//line admin\views\article\index.qtpl:17
}

//line admin\views\article\index.qtpl:17
func (p *IndexPage) WriteRenderTitle(qq422016 qtio422016.Writer) {
	//line admin\views\article\index.qtpl:17
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line admin\views\article\index.qtpl:17
	p.StreamRenderTitle(qw422016)
	//line admin\views\article\index.qtpl:17
	qt422016.ReleaseWriter(qw422016)
//line admin\views\article\index.qtpl:17
}

//line admin\views\article\index.qtpl:17
func (p *IndexPage) RenderTitle() string {
	//line admin\views\article\index.qtpl:17
	qb422016 := qt422016.AcquireByteBuffer()
	//line admin\views\article\index.qtpl:17
	p.WriteRenderTitle(qb422016)
	//line admin\views\article\index.qtpl:17
	qs422016 := string(qb422016.B)
	//line admin\views\article\index.qtpl:17
	qt422016.ReleaseByteBuffer(qb422016)
	//line admin\views\article\index.qtpl:17
	return qs422016
//line admin\views\article\index.qtpl:17
}

//line admin\views\article\index.qtpl:18
func (p *IndexPage) StreamRenderKwd(qw422016 *qt422016.Writer) {
	//line admin\views\article\index.qtpl:18
	p.BasePage.StreamRenderKwd(qw422016)
	//line admin\views\article\index.qtpl:18
	qw422016.N().S(` `)
//line admin\views\article\index.qtpl:18
}

//line admin\views\article\index.qtpl:18
func (p *IndexPage) WriteRenderKwd(qq422016 qtio422016.Writer) {
	//line admin\views\article\index.qtpl:18
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line admin\views\article\index.qtpl:18
	p.StreamRenderKwd(qw422016)
	//line admin\views\article\index.qtpl:18
	qt422016.ReleaseWriter(qw422016)
//line admin\views\article\index.qtpl:18
}

//line admin\views\article\index.qtpl:18
func (p *IndexPage) RenderKwd() string {
	//line admin\views\article\index.qtpl:18
	qb422016 := qt422016.AcquireByteBuffer()
	//line admin\views\article\index.qtpl:18
	p.WriteRenderKwd(qb422016)
	//line admin\views\article\index.qtpl:18
	qs422016 := string(qb422016.B)
	//line admin\views\article\index.qtpl:18
	qt422016.ReleaseByteBuffer(qb422016)
	//line admin\views\article\index.qtpl:18
	return qs422016
//line admin\views\article\index.qtpl:18
}

//line admin\views\article\index.qtpl:20
func (p *IndexPage) StreamRenderCss(qw422016 *qt422016.Writer) {
//line admin\views\article\index.qtpl:20
}

//line admin\views\article\index.qtpl:20
func (p *IndexPage) WriteRenderCss(qq422016 qtio422016.Writer) {
	//line admin\views\article\index.qtpl:20
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line admin\views\article\index.qtpl:20
	p.StreamRenderCss(qw422016)
	//line admin\views\article\index.qtpl:20
	qt422016.ReleaseWriter(qw422016)
//line admin\views\article\index.qtpl:20
}

//line admin\views\article\index.qtpl:20
func (p *IndexPage) RenderCss() string {
	//line admin\views\article\index.qtpl:20
	qb422016 := qt422016.AcquireByteBuffer()
	//line admin\views\article\index.qtpl:20
	p.WriteRenderCss(qb422016)
	//line admin\views\article\index.qtpl:20
	qs422016 := string(qb422016.B)
	//line admin\views\article\index.qtpl:20
	qt422016.ReleaseByteBuffer(qb422016)
	//line admin\views\article\index.qtpl:20
	return qs422016
//line admin\views\article\index.qtpl:20
}

//line admin\views\article\index.qtpl:22
func (p *IndexPage) StreamRenderScript(qw422016 *qt422016.Writer) {
//line admin\views\article\index.qtpl:22
}

//line admin\views\article\index.qtpl:22
func (p *IndexPage) WriteRenderScript(qq422016 qtio422016.Writer) {
	//line admin\views\article\index.qtpl:22
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line admin\views\article\index.qtpl:22
	p.StreamRenderScript(qw422016)
	//line admin\views\article\index.qtpl:22
	qt422016.ReleaseWriter(qw422016)
//line admin\views\article\index.qtpl:22
}

//line admin\views\article\index.qtpl:22
func (p *IndexPage) RenderScript() string {
	//line admin\views\article\index.qtpl:22
	qb422016 := qt422016.AcquireByteBuffer()
	//line admin\views\article\index.qtpl:22
	p.WriteRenderScript(qb422016)
	//line admin\views\article\index.qtpl:22
	qs422016 := string(qb422016.B)
	//line admin\views\article\index.qtpl:22
	qt422016.ReleaseByteBuffer(qb422016)
	//line admin\views\article\index.qtpl:22
	return qs422016
//line admin\views\article\index.qtpl:22
}

//line admin\views\article\index.qtpl:24
func (p *IndexPage) StreamRenderBody(qw422016 *qt422016.Writer) {
	//line admin\views\article\index.qtpl:24
	qw422016.N().S(`
	`)
	//line admin\views\article\index.qtpl:26
	modelVal := p.BasePage.BPD.Data["Model"]
	pageIndex := p.BPD.Data["page"].(int)
	total := p.BasePage.BPD.Data["Total"].(int)
	var arts []model.Article
	if modelVal != nil {
		arts = modelVal.([]model.Article)
	}

	//line admin\views\article\index.qtpl:33
	qw422016.N().S(`
	<div class="text-center">
		<a class="btn btn-default" href="`)
	//line admin\views\article\index.qtpl:35
	qw422016.E().S(area.Url(p.BasePage, "create", nil))
	//line admin\views\article\index.qtpl:35
	qw422016.N().S(`">添加</a>
		
		`)
	//line admin\views\article\index.qtpl:37
	if arts != nil && len(arts) > 0 {
		//line admin\views\article\index.qtpl:37
		qw422016.N().S(`
		<table class="table table-bordered table-hover">
		`)
		//line admin\views\article\index.qtpl:39
		for _, v := range arts {
			//line admin\views\article\index.qtpl:39
			qw422016.N().S(`
		`)
			//line admin\views\article\index.qtpl:41
			id := int(v.ID)
			data := make(map[string]interface{})
			data["id"] = id

			//line admin\views\article\index.qtpl:44
			qw422016.N().S(`
			<tr>
				<td><a href="`)
			//line admin\views\article\index.qtpl:46
			qw422016.E().S(area.Url(p.BasePage, "detail", data))
			//line admin\views\article\index.qtpl:46
			qw422016.N().S(`">`)
			//line admin\views\article\index.qtpl:46
			qw422016.E().S(v.Title)
			//line admin\views\article\index.qtpl:46
			qw422016.N().S(`</a></td>
				<td>`)
			//line admin\views\article\index.qtpl:47
			qw422016.E().S(v.CreateTime.Format("2006-01-02 15:04:05"))
			//line admin\views\article\index.qtpl:47
			qw422016.N().S(`</td>
				<td>`)
			//line admin\views\article\index.qtpl:48
			if v.Used {
				//line admin\views\article\index.qtpl:48
				qw422016.N().S(`是`)
				//line admin\views\article\index.qtpl:48
			} else {
				//line admin\views\article\index.qtpl:48
				qw422016.N().S(`否`)
				//line admin\views\article\index.qtpl:48
			}
			//line admin\views\article\index.qtpl:48
			qw422016.N().S(`</td>
				<td>
					<a href="`)
			//line admin\views\article\index.qtpl:50
			qw422016.E().S(area.Url(p.BasePage, "edit", data))
			//line admin\views\article\index.qtpl:50
			qw422016.N().S(`">编辑</a>
					<a href="`)
			//line admin\views\article\index.qtpl:51
			qw422016.E().S(area.Url(p.BasePage, "mdelete", data))
			//line admin\views\article\index.qtpl:51
			qw422016.N().S(`">删除</a>
				</td>
			</tr>
		`)
			//line admin\views\article\index.qtpl:54
		}
		//line admin\views\article\index.qtpl:54
		qw422016.N().S(`
		</table>
		`)
		//line admin\views\article\index.qtpl:56
		qw422016.N().S(commonhtmlhelper.GeneratePageLink("", pageIndex, 3, total, "page", nil))
		//line admin\views\article\index.qtpl:56
		qw422016.N().S(`
		`)
		//line admin\views\article\index.qtpl:57
	}
	//line admin\views\article\index.qtpl:57
	qw422016.N().S(`
	</div>
`)
//line admin\views\article\index.qtpl:59
}

//line admin\views\article\index.qtpl:59
func (p *IndexPage) WriteRenderBody(qq422016 qtio422016.Writer) {
	//line admin\views\article\index.qtpl:59
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line admin\views\article\index.qtpl:59
	p.StreamRenderBody(qw422016)
	//line admin\views\article\index.qtpl:59
	qt422016.ReleaseWriter(qw422016)
//line admin\views\article\index.qtpl:59
}

//line admin\views\article\index.qtpl:59
func (p *IndexPage) RenderBody() string {
	//line admin\views\article\index.qtpl:59
	qb422016 := qt422016.AcquireByteBuffer()
	//line admin\views\article\index.qtpl:59
	p.WriteRenderBody(qb422016)
	//line admin\views\article\index.qtpl:59
	qs422016 := string(qb422016.B)
	//line admin\views\article\index.qtpl:59
	qt422016.ReleaseByteBuffer(qb422016)
	//line admin\views\article\index.qtpl:59
	return qs422016
//line admin\views\article\index.qtpl:59
}
