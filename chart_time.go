package gochart

import (
	"github.com/bitly/go-simplejson"
	"strconv"
)

type ChartTime struct {
	ChartBase
	YMax          string
	TickInterval  string
	TickLabelStep string
	PlotLinesY    string
	TickFormat    string
	TickUnit      int
}

func (this *ChartTime) Build(dataArray string) {
	if this.chartArgs == nil {
		this.ChartBase.BuildBase(dataArray)
		this.chartArgs["YMax"] = this.YMax
		if this.TickLabelStep == "" {
			this.TickLabelStep = "60"
		}
		this.chartArgs["TickLabelStep"] = this.TickLabelStep
		this.chartArgs["PlotLinesY"] = this.PlotLinesY
		if this.TickUnit == 0 {
			this.TickUnit = 1000
		}
		v, _ := strconv.Atoi(this.RefreshTime)
		this.chartArgs["TickInterval"] = strconv.Itoa(v * this.TickUnit)
	} else {
		this.ChartBase.BuildBase(dataArray)
	}
}

func (this *ChartTime) Template() string {
	return TemplateTimeHtml
}

func (this *ChartTime) AddData(name string, data interface{}, UTCTime int64, samplenum, refreshtime int) *simplejson.Json {
	endtime := 1000 * int(8*60*60+UTCTime)
	begintime := endtime - this.TickUnit*samplenum*refreshtime
	var json *simplejson.Json
	json = simplejson.New()
	json.Set("name", name)
	json.Set("data", data)
	json.Set("pointInterval", refreshtime*this.TickUnit)
	json.Set("pointStart", begintime)
	json.Set("pointEnd", endtime)
	return json
}
