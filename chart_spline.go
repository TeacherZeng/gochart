package gochart

type ChartSpline struct {
	ChartBase
}

func (this *ChartSpline) Template() string {
	return TemplateSplineHtml
}
