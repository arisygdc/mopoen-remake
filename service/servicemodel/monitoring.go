package servicemodel

import (
	"time"

	"github.com/google/uuid"
)

type DaftarMonitoring struct {
	TipeSensor  int32  `json:"tipe_sensor"`
	Location_id int32  `json:"location_id"`
	Nama        string `json:"nama"`
	Keterangan  string `json:"keterangan"`
}

type MonitoringTerdaftar struct {
	ID           uuid.UUID `json:"id"`
	TipeSensorID int32     `json:"tipe_sensor_id"`
	LokasiID     int32     `json:"lokasi_id"`
	Nama         string    `json:"nama"`
	Keterangan   string    `json:"keterangan"`
}

type DetailMonitoringTerdaftar struct {
	ID         uuid.UUID   `json:"id"`
	TipeSensor TipeSensor  `json:"tipe_sensor"`
	LokasiID   FetchLokasi `json:"lokasi"`
	Nama       string      `json:"nama"`
	Keterangan string      `json:"keterangan"`
}

type MonitoringData struct {
	Value      float64   `json:"value"`
	DibuatPada time.Time `json:"dibuat_pada"`
}

type ResultMonitoring struct {
	Total   int64   `json:"total"`
	Average float64 `json:"average"`
}

type AnalisaMonitoring struct {
	Overall   ResultMonitoring `json:"overall"`
	Morning   ResultMonitoring `json:"morning"`
	Afternoon ResultMonitoring `json:"afternoon"`
	Noon      ResultMonitoring `json:"noon"`
	Night     ResultMonitoring `json:"night"`
	Midnight  ResultMonitoring `json:"midnight"`
}
