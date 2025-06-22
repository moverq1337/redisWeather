package models

type Weather struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}
