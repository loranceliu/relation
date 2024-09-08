package vo

type ChartCommonVO struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ChartBarVO struct {
	Name []string `json:"name"`
	Data []string `json:"data"`
}

type LineChartCommonVO struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Time  string `json:"time"`
}

type SeriesLineChart struct {
	Name  string   `json:"name"`
	Type  string   `json:"type"`
	Stack string   `json:"stack"`
	Data  []string `json:"data"`
}

type LineChartDataVO struct {
	TypeData   []string          `json:"typeData"`
	CateData   []string          `json:"cateData"`
	SeriesData []SeriesLineChart `json:"seriesData"`
}
