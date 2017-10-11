package gochart

type ChartTime struct {
	ChartBase
	TimeFormat string
}

func (this *ChartTime) Build() {
	this.ChartBase.BuildBase()
	this.chartArgs["TimeFormat"] = this.TimeFormat
}

func (this *ChartTime) Template() string {
	return TemplateTimeHtml
}
