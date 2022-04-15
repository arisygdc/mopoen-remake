package request

type PostSensor struct {
	Tipe   string `json:"tipe"`
	Satuan string `json:"satuan"`
}

type DeleteSensor struct {
	Id int32 `json:"sensor_id"`
}
