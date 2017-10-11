package gochart

type ChartTime struct {
	ChartBase
}

func (this *ChartTime) Build() {
	this.ChartBase.BuildBase()
}

func (this *ChartTime) Template() string {
	return TemplateTimeHtml
}
