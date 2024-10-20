package main

type Weather struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
	Forecast Forecast `json:"forecast"`
}

type Location struct {
	Name string `json:"name"`

	Country string `json:"country"`
}

type Current struct {
	Condition Condition `json:"condition"`
	TempF     float64   `json:"temp_f"`
}

type Condition struct {
	Text string `json:"text"`
}

type Forecast struct {
	Forecastday []ForecastDay `json:"forecastday"`
}

type ForecastDay struct {
	Hour []Hour `json:"hour"`
}

type Hour struct {
	Condition    Condition `json:"condition"`
	TimeEpoch    int64     `json:"time_epoch"`
	TempF        float64   `json:"temp_f"`
	ChanceOfRain float64   `json:"chance_of_rain"`
}
