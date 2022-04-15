package request

import "github.com/google/uuid"

type PostDaftarMonitoring struct {
	TipeSensor  int32  `json:"tipe_sensor"`
	Location_id int32  `json:"location_id"`
	Nama        string `json:"nama"`
	Keterangan  string `json:"keterangan"`
}

type PostMonitoringValue struct {
	KodeMonitoring uuid.UUID `json:"kode_monitoring"`
	Value          float64   `json:"value"`
}
