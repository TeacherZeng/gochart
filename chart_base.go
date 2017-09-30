package gochart

type IChart interface {
	Update()
}

type IChartInner interface {
	IChart
	Template() string
	Build()
	Data() map[string]string
}

type ChartBase struct {
	ChartType    string
	Title        string
	SubTitle     string
	YAxisText    string
	XAxisNumbers string
	ValueSuffix  string
	Height       string
	DataArray    string
	chartArgs    map[string]string
}

func (this *ChartBase) Build() {
	if this.chartArgs == nil {
		this.chartArgs = make(map[string]string)
		this.chartArgs["ChartType"] = this.ChartType
		this.chartArgs["Title"] = this.Title
		this.chartArgs["SubTitle"] = this.SubTitle
		this.chartArgs["YAxisText"] = this.YAxisText
		this.chartArgs["XAxisNumbers"] = this.XAxisNumbers
		this.chartArgs["ValueSuffix"] = this.ValueSuffix
		this.chartArgs["Height"] = this.Height
	}
	this.chartArgs["DataArray"] = this.DataArray
}

func (this *ChartBase) Data() map[string]string {
	return this.chartArgs
}
