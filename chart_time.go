package gochart

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	_ "github.com/golang/glog"
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

func (this *ChartTime) Load(filename string) (bool, []interface{}) {
	ok, root := this.LoadBase(filename)
	if !ok {
		return false, nil
	}
	this.TickInterval, _ = root.Get("TickInterval").String()
	this.TickLabelStep, _ = root.Get("TickLabelStep").String()
	this.PlotLinesY, _ = root.Get("PlotLinesY").String()
	tmpv, _ := strconv.Atoi(this.TickInterval)
	this.TickUnit = tmpv / this.RefreshTime

	outdatas, _ := root.Get("DataArray").String()
	outdatas = fmt.Sprintf("{\"DataArray\":%s}", outdatas)
	json, _ := simplejson.NewJson([]byte(outdatas))
	arrays, _ := json.Get("DataArray").Array()

	datas := make([]interface{}, 0)
	endtime := this.beginTime + int64(this.TickUnit*this.SampleNum*this.RefreshTime)
	for _, val := range arrays {
		temp := val.(map[string]interface{})
		json := simplejson.New()
		json.Set("name", temp["name"])
		json.Set("data", temp["data"])
		json.Set("pointInterval", this.RefreshTime*this.TickUnit)
		json.Set("pointStart", this.beginTime)
		json.Set("pointEnd", endtime)
		datas = append(datas, json)
	}
	return true, datas
}
