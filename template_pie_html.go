package gochart

// see the resource of : http://www.freecdn.cn/libs/highcharts/

var TemplatePieHtml = `{{define "T"}}
<!DOCTYPE HTML>
<html>
    <head>
	    <meta http-equiv="refresh" content="1">
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
        <title>Gochart - {{.ChartType}}</title>

        <script type="text/javascript" src="/js/jquery-1.8.3.min.js"></script>
        <script type="text/javascript">
        $(function () {
            $('#container').highcharts({
                chart: {
                    //type: 'pie',
                    type: '{{.ChartType}}',
                    plotBackgroundColor: null,
                    plotBorderWidth: null,
                    plotShadow: false,
                    animation: false
                },
                title: {
                    // text: 'Browser market shares at a specific website, 2014'
                    text: '{{.Title}}',
                },
                subtitle: {
                    // text: 'Source: somewebsite.com',
                    text: '{{.SubTitle}}',
                },
                tooltip: {
                    pointFormat: '{series.name}: <b>{point.percentage:.1f}%</b>'
                },
                plotOptions: {
                    pie: {
                        allowPointSelect: true,
                        cursor: 'pointer',
                        dataLabels: {
                            enabled: true,
                            format: '<b>{point.name}</b>: {point.percentage:.1f} %',
                            style: {
                                color: (Highcharts.theme && Highcharts.theme.contrastTextColor) || 'black'
                            }
                        }
                    },
                    series: {
                        animation: false
                    }
                },
                credits : {
                    enabled: false
                },
                series: [{
                    // name: 'Browser share',
                    name : '{{.SeriesName}}',
                    data: 
                    	{{.DataArray}}
                    /*
                    data: 
                    [
                        ['Firefox',   45.0],
                        ['IE',       26.8],
                        ['Chrome',  12.8],
                        ['Safari',    8.5],
                        ['Opera',     6.2],
                        ['Others',   0.7]
                    ]
                    */
                }]
            });
        });
		</script>
    </head>
    <body>
    
    <script type="text/javascript" src="/js/highcharts.js"></script>
    <script type="text/javascript" src="/js/exporting.js"></script>

    <div id="container" style="min-width: 310px; height: {{.Height}}px; max-width: 600px; margin: 0 auto"></div>

    </body>
</html>
{{end}}
`
