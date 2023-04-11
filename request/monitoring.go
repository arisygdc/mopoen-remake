package request

type PostDaftarMonitoring struct {
	TipeSensor_id int32  `json:"tipe_sensor_id" binding:"required,min=1"`
	Location_id   int32  `json:"lokasi_id" binding:"required,min=1"`
	Nama          string `json:"nama" binding:"required,min=4"`
	Keterangan    string `json:"keterangan"`
}

type PostMonitoringValue struct {
	KodeMonitoring string  `json:"kode_monitoring" binding:"required,min=36,max=36"`
	Value          float64 `json:"value"`
}

type GetMonitoringTerdaftar struct {
	Lokasi int32 `json:"lokasi_id" binding:"required,min=1"`
}

type GetUUID struct {
	ID string `uri:"uuid" binding:"required,min=36,max=36"`
}
