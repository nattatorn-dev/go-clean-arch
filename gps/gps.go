package gps

import "time"

type gps struct {
	ID string `json:"id" validate:"required"`
	Name string `json:"name"`

	BoxID string `json:"BoxID"`
	DataStatus int8 `json:"DataStatus"`

	SendTime time.Time `json:"SendTime"`
	GpsTime time.Time `json:"GpsTime"`
	// 0=valid, 1=invalid
	GpsStatus int8 `json:"GpsStatus"`
	Latitude float32 `json:"Latitude"`
	Longitude float32 `json:"Longitude"`
	// km/h
	Speed float32 `json:"Speed"`
	// 0-360
	Direction int8 `json:"Direction"`
	// 0=on, 1=off
	EngineStatus int8 `json:"EngineStatus"`
	// 0=on, 1=off
	SOS int8 `json:"SOS"`
}