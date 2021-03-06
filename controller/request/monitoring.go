package request

type PostDaftarMonitoring struct {
	TipeSensor  int32  `json:"tipe_sensor" binding:"required,min=1"`
	Location_id int32  `json:"lokasi_id" binding:"required,min=1"`
	Nama        string `json:"nama" binding:"required,min=4"`
	Keterangan  string `json:"keterangan"`
}

type PostMonitoringValue struct {
	KodeMonitoring string  `json:"kode_monitoring" binding:"required,min=36,max=36"`
	Value          float64 `json:"value" binding:"required,min=1"`
}

type GetMonitoringTerdaftar struct {
	Lokasi int32 `json:"lokasi_id" binding:"required,min=1"`
}
