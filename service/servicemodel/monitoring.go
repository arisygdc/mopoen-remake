package servicemodel

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

var ErrWrongSecret = fmt.Errorf("wrong secret")

type DaftarMonitoring struct {
	TipeSensor_id int32  `json:"tipe_sensor_id"`
	Location_id   int32  `json:"location_id"`
	Email         string `json:"email"`
	Author        string `json:"author"`
	Nama          string `json:"nama"`
	Keterangan    string `json:"keterangan"`
}

type GetMonitoringTerdaftarFilterOptions struct {
	LokasiID     int32 `json:"lokasi_id"`
	TipeSensorID int32 `json:"tipe_sensor_id"`
}

type MonitoringTerdaftar struct {
	ID           uuid.UUID `json:"id"`
	TipeSensorID int32     `json:"tipe_sensor_id"`
	LokasiID     int32     `json:"lokasi_id"`
	Email        string    `json:"email"`
	Author       string    `json:"author"`
	Nama         string    `json:"nama"`
	Keterangan   string    `json:"keterangan"`
}

type DetailMonitoringTerdaftar struct {
	MonitoringID uuid.UUID `json:"monitoring_id"`
	TipeSensorID int32     `json:"tipe_sensor_id"`
	TipeSensor   string    `json:"tipe_sensor"`
	Nama         string    `json:"nama"`
	Keterangan   string    `json:"keterangan"`
	Address      string    `json:"alamat"`
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
	Overall   ResultMonitoring `json:"keseluruhan"`
	Morning   ResultMonitoring `json:"pagi"`
	Noon      ResultMonitoring `json:"siang"`
	Afternoon ResultMonitoring `json:"sore"`
	Night     ResultMonitoring `json:"malam"`
}
