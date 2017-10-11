package gochart

type ChartPie struct {
	ChartBase
}

func (this *ChartPie) Build() {
	this.ChartBase.BuildBase()
}

func (this *ChartPie) Template() string {
	return TemplatePieHtml
}
