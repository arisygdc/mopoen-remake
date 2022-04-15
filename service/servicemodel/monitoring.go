package servicemodel

type DaftarMonitoring struct {
	TipeSensor  int32  `json:"tipe_sensor"`
	Location_id int32  `json:"location_id"`
	Nama        string `json:"nama"`
	Keterangan  string `json:"keterangan"`
}
