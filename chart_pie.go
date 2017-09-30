package gochart

type ChartPie struct {
	ChartBase
}

func (this *ChartPie) Template() string {
	return TemplatePieHtml
}
