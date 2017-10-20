package gochart

import (
	"github.com/bitly/go-simplejson"
	"github.com/golang/glog"
	"net/http"
	"os"
	"text/template"
	"time"
)

type ChartServer struct {
	charts map[string]IChartInner
}

func (this *ChartServer) AddChart(chartname string, chart IChartInner, savedata bool) {
	if this.charts == nil {
		this.charts = make(map[string]IChartInner)
	}
	this.charts[chartname] = chart
	if savedata {
		chart.GoSaveData(chartname)
	}
}

func (this *ChartServer) ListenAndServe(addr string) error {
	http.HandleFunc("/", this.handler)
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {})
	http.HandleFunc("/js/", this.js)
	return http.ListenAndServe(addr, nil)
}

func (this *ChartServer) handler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	chartname := values.Get("query")

	if _, ok := this.charts[chartname]; !ok {
		glog.Errorln("no find the chart, chartname =", chartname)
		return
	}

	now := time.Now().Unix()

	chart := this.charts[chartname]
	datas := chart.Update(now)
	chart.SaveData(datas)
	json := simplejson.New()
	json.Set("DataArray", datas)
	b, _ := json.Get("DataArray").Encode()
	chart.Build(string(b))

	if t, err := template.New("foo").Parse(chart.Template()); err != nil {
		w.Write([]byte(err.Error()))
	} else {
		if err = t.ExecuteTemplate(w, "T", chart.Data()); err != nil {
			w.Write([]byte(err.Error()))
		}
	}
}

func (this *ChartServer) js(w http.ResponseWriter, r *http.Request) {
	wd, err := os.Getwd()
	if err != nil {
		glog.Errorln(err)
		return
	}
	http.FileServer(http.Dir(wd)).ServeHTTP(w, r)
}
