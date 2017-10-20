package gochart

import (
	"github.com/golang/glog"
	"runtime/debug"
	"time"
)

type IChart interface {
	Update(now int64) []interface{}
}

type IChartInner interface {
	IChart
	Template() string
	Build(dataArray string)
	Data() map[string]string

	// save data
	GoSaveData(filename string)
	IsEnableSaveData() bool
	SaveData(datas []interface{})
}

type ChartBase struct {
	ChartType    string
	Title        string
	SubTitle     string
	YAxisText    string
	XAxisNumbers string
	ValueSuffix  string
	SeriesName   string
	RefreshTime  string
	chartArgs    map[string]string

	// save data
	filename     string
	saveData     []interface{}
	chanSaveData chan []interface{}
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
		if this.RefreshTime == "" {
			this.RefreshTime = "60"
		}
		this.chartArgs["RefreshTime"] = this.RefreshTime
	}
	this.chartArgs["DataArray"] = dataArray
}

func (this *ChartBase) Data() map[string]string {
	return this.chartArgs
}

func (this *ChartBase) GoSaveData(filename string) {
	this.filename = filename
	this.chanSaveData = make(chan []interface{}, 1)
	this.saveData = make([]interface{}, 0)

	go func() {
		defer func() {
			if err := recover(); err != nil {
				glog.Errorln("[异常] ", err, "\n", string(debug.Stack()))
			}
		}()

		newDataFlag := false
		tick := time.NewTicker(30 * time.Second)
		select {
		case datas := <-this.chanSaveData:
			if len(datas) > 0 {
				newDataFlag = true

				data := datas[len(datas)-1]
				this.saveData = append(this.saveData, data)

				//			datas := make([]interface{}, 0)
				//			var json *simplejson.Json
				//			json = simplejson.New()
				//			json.Set("name", name)
				//			json.Set("data", data)
				//			json.Set("pointInterval", refreshtime*this.TickUnit)
				//			json.Set("pointStart", begintime)
				//			json.Set("pointEnd", endtime)
				//			datas = append(datas, json)
			}
		case <-tick.C:
			if newDataFlag {
				newDataFlag = false
			}
		}
	}()
}

func (this *ChartBase) IsEnableSaveData() bool {
	return this.filename != ""
}

func (this *ChartBase) SaveData(datas []interface{}) {
	this.chanSaveData <- datas
}
