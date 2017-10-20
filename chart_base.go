package gochart

import (
	"github.com/bitly/go-simplejson"
	"github.com/golang/glog"
	"runtime/debug"
	"strconv"
	"sync"
	"time"
)

type IChart interface {
	Update(now int64) map[string][]interface{}
}

type IChartInner interface {
	IChart
	Init()
	Template() string
	Build(dataArray string)
	Data() map[string]string
	AddData(map[string][]interface{}, int64) []interface{}
	GetSaveData() []interface{}

	// save data
	GoSaveData(filename string)
	IsEnableSaveData() bool
	SaveData(datas map[string][]interface{})
}

type ChartBase struct {
	ChartType    string
	Title        string
	SubTitle     string
	YAxisText    string
	XAxisNumbers string
	ValueSuffix  string
	SeriesName   string
	RefreshTime  int
	SampleNum    int
	chartArgs    map[string]string
	m            sync.RWMutex

	//chart data
	chartData map[string][]interface{}

	// save data
	filename     string
	saveData     map[string][]interface{}
	chanSaveData chan map[string][]interface{}
	beginTime    int64
}

func (this *ChartBase) InitBase() {
	this.chartArgs = make(map[string]string)
	this.chartArgs["ChartType"] = this.ChartType
	this.chartArgs["Title"] = this.Title
	this.chartArgs["SubTitle"] = this.SubTitle
	this.chartArgs["YAxisText"] = this.YAxisText
	this.chartArgs["XAxisNumbers"] = this.XAxisNumbers
	this.chartArgs["ValueSuffix"] = this.ValueSuffix
	this.chartArgs["SeriesName"] = this.SeriesName
	if this.RefreshTime == 0 {
		this.RefreshTime = 60
	}
	this.chartArgs["RefreshTime"] = strconv.Itoa(this.RefreshTime)

	this.chartData = make(map[string][]interface{})
}

func (this *ChartBase) Build(dataArray string) {
	this.m.Lock()
	this.chartArgs["DataArray"] = dataArray
	this.m.Unlock()
}

func (this *ChartBase) Data() map[string]string {
	return this.chartArgs
}

func (this *ChartBase) GoSaveData(filename string) {
	this.filename = filename
	this.chanSaveData = make(chan map[string][]interface{}, 1)
	this.saveData = make(map[string][]interface{})

	go func() {
		defer func() {
			if err := recover(); err != nil {
				glog.Errorln("[异常] ", err, "\n", string(debug.Stack()))
			}
		}()

		newDataFlag := false
		tick := time.NewTicker(30 * time.Second)
		for {
			select {
			case datas := <-this.chanSaveData:
				if len(datas) > 0 {

					if this.beginTime == 0 {
						this.beginTime = time.Now().Unix()
					}

					newDataFlag = true
					for k, v := range datas {
						if _, ok := this.saveData[k]; !ok {
							this.saveData[k] = make([]interface{}, 0)
						}
						for _, tempv := range v {
							this.saveData[k] = append(this.saveData[k], tempv)
						}
					}
				}
			case <-tick.C:
				if newDataFlag {
					newDataFlag = false

					root := simplejson.New()
					this.m.RLock()
					for key, val := range this.chartArgs {
						root.Set(key, val)
					}
					this.m.RUnlock()

					root.Set("beginTime", this.beginTime)

					//					outdatas := this.GetSaveData()
					//					json := simplejson.New()
					//					json.Set("DataArray", outdatas)
					//					b, _ := json.Get("DataArray").Encode()
					//					string(b)

				}
			}
		}
	}()
}

func (this *ChartBase) IsEnableSaveData() bool {
	return this.filename != ""
}

func (this *ChartBase) SaveData(datas map[string][]interface{}) {
	if this.chanSaveData != nil {
		this.chanSaveData <- datas
	}
}
