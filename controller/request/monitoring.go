package request

import "github.com/google/uuid"

type PostDaftarMonitoring struct {
	TipeSensor  int32  `json:"tipe_sensor" binding:"required,min=1"`
	Location_id int32  `json:"location_id" binding:"required,min=1"`
	Nama        string `json:"nama" binding:"required,min=4"`
	Keterangan  string `json:"keterangan"`
}

type PostMonitoringValue struct {
	KodeMonitoring uuid.UUID `json:"kode_monitoring" binding:"required,min=36,max=36"`
	Value          float64   `json:"value" binding:"required,min=1"`
}
