package gochart

type ChartSpline struct {
	ChartBase
}

func (this *ChartSpline) Build() {
	this.ChartBase.BuildBase()
}

func (this *ChartSpline) Template() string {
	return TemplateSplineHtml
}
