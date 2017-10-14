package gochart

type IChart interface {
	Update(now int64) []interface{}
}

type IChartInner interface {
	IChart
	Template() string
	Build(dataArray string)
	Data() map[string]string
}

type ChartBase struct {
	ChartType    string
	Title        string
	SubTitle     string
	YAxisText    string
	XAxisNumbers string
	ValueSuffix  string
	SeriesName   string
	Height       string
	RefreshTime  string
	chartArgs    map[string]string
}

func (this *ChartBase) BuildBase(dataArray string) {
	if this.chartArgs == nil {
		this.chartArgs = make(map[string]string)
		this.chartArgs["ChartType"] = this.ChartType
		this.chartArgs["Title"] = this.Title
		this.chartArgs["SubTitle"] = this.SubTitle
		this.chartArgs["YAxisText"] = this.YAxisText
		this.chartArgs["XAxisNumbers"] = this.XAxisNumbers
		this.chartArgs["ValueSuffix"] = this.ValueSuffix
		this.chartArgs["SeriesName"] = this.SeriesName
		this.chartArgs["Height"] = this.Height
		r := "60"
		if this.RefreshTime != "" {
			r = this.RefreshTime
		}
		this.chartArgs["RefreshTime"] = r
	}
	this.chartArgs["DataArray"] = dataArray
}

func (this *ChartBase) Data() map[string]string {
	return this.chartArgs
}
