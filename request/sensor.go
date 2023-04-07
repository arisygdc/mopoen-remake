package request

type PostSensor struct {
	Tipe   string `json:"tipe" binding:"required,min=3"`
	Satuan string `json:"satuan" binding:"required,min=1"`
}

type DeleteSensor struct {
	Id int32 `json:"sensor_id" binding:"required,min=1"`
}

type GetSensor struct {
	Id int32 `uri:"id"`
}
