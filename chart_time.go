package gochart

type ChartTime struct {
	ChartBase
	YMax         string
	TickInterval string
}

func (this *ChartTime) Build() {
	this.ChartBase.BuildBase()
	this.chartArgs["YMax"] = this.YMax
	this.chartArgs["TickInterval"] = this.TickInterval
}

func (this *ChartTime) Template() string {
	return TemplateTimeHtml
}
