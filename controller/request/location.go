package request

type PostProvinsi struct {
	Nama string `json:"nama"`
}

type PostKabupaten struct {
	Provinsi_id int32  `json:"provinsi_id"`
	Nama        string `json:"nama"`
}

type PostKecamatan struct {
	Kabupaten_id int32  `json:"kabupaten_id"`
	Nama         string `json:"nama"`
}

type PostDesa struct {
	Kecamatan_id int32  `json:"kecamatan_id"`
	Nama         string `json:"nama"`
}

type DeleteLocationType struct {
	Type string `form:"type"`
}

type DeleteLocation struct {
	Id int32 `json:"id"`
}
