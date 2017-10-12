package gochart

type ChartTime struct {
	ChartBase
	YMax string
}

func (this *ChartTime) Build() {
	this.ChartBase.BuildBase()
	this.chartArgs["YMax"] = this.YMax
}

func (this *ChartTime) Template() string {
	return TemplateTimeHtml
}
