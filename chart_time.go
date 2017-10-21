package gochart

import (
	"github.com/bitly/go-simplejson"
	"strconv"
)

type ChartTime struct {
	ChartBase
	TickInterval  string
	TickLabelStep string
	PlotLinesY    string
	TickUnit      int
}

func (this *ChartTime) Init() {
	this.ChartBase.InitBase()
	this.chartClassType = CCT_TIME
	if this.TickLabelStep == "" {
		this.TickLabelStep = "60"
	}
	this.chartArgs["TickLabelStep"] = this.TickLabelStep
	this.chartArgs["PlotLinesY"] = this.PlotLinesY
	if this.TickUnit == 0 {
		this.TickUnit = 1000
	}
	this.chartArgs["TickInterval"] = strconv.Itoa(this.RefreshTime * this.TickUnit)
}

func (this *ChartTime) Template() string {
	return TemplateTimeHtml
}

func (this *ChartTime) AddData(newDatas map[string][]interface{}, UTCTime int64) []interface{} {
	endtime := 1000 * int(8*60*60+UTCTime)
	begintime := endtime - this.TickUnit*this.SampleNum*this.RefreshTime
	datas := make([]interface{}, 0)
	for k, v := range newDatas {
		if _, ok := this.chartData[k]; !ok {
			this.chartData[k] = make([]interface{}, this.SampleNum)
		}
		for _, tempv := range v {
			this.chartData[k] = append(this.chartData[k], tempv)
		}
		for len(this.chartData[k]) > this.SampleNum {
			this.chartData[k] = this.chartData[k][1:]
		}
		var json *simplejson.Json
		json = simplejson.New()
		json.Set("name", k)
		json.Set("data", this.chartData[k])
		json.Set("pointInterval", this.RefreshTime*this.TickUnit)
		json.Set("pointStart", begintime)
		json.Set("pointEnd", endtime)
		datas = append(datas, json)
	}
	return datas
}

func (this *ChartTime) Load(filename string) bool {
	if this.LoadBase(filename) == false {
		return false
	}

	return true
}
