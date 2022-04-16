package request

type UriParamTipeLocation struct {
	Tipe string `uri:"tipe" binding:"required,min=4"`
}

type DeleteLocation struct {
	Id int32 `json:"id" binding:"required,min=1"`
}

type PostProvinsi struct {
	Nama string `json:"nama" binding:"required,min=4"`
}

type PostKabupaten struct {
	Provinsi_id int32  `json:"provinsi_id" binding:"required,min=1"`
	Nama        string `json:"nama" binding:"required,min=4"`
}

type PostKecamatan struct {
	Kabupaten_id int32  `json:"kabupaten_id" binding:"required,min=1"`
	Nama         string `json:"nama" binding:"required,min=4"`
}

type PostDesa struct {
	Kecamatan_id int32  `json:"kecamatan_id" binding:"required,min=1"`
	Nama         string `json:"nama" binding:"required,min=4"`
}
