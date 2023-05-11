package request

type PostDaftarMonitoring struct {
	TipeSensor_id int32  `json:"tipe_sensor_id" binding:"required,min=1"`
	Location_id   int32  `json:"lokasi_id" binding:"required,min=1"`
	Email         string `json:"email" binding:"required,email"`
	Author        string `json:"author" binding:"required,min=4"`
	Nama          string `json:"nama" binding:"required,min=4"`
	Keterangan    string `json:"keterangan"`
}

type PostMonitoringValue struct {
	KodeMonitoring string  `json:"id" binding:"required,min=36,max=36"`
	Secret         string  `json:"secret" binding:"required,min=32,max=33"`
	Value          float64 `json:"value"`
}

type GetMonitoringTerdaftar struct {
	Lokasi int32 `json:"lokasi_id" binding:"required,min=1"`
}

type GetUUID struct {
	ID string `uri:"uuid" binding:"required,min=36,max=36"`
}

type GetFile struct {
	FileName string `uri:"file" binding:"required,min=8,max=255"`
}
