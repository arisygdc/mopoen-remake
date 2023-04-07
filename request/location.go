package request

type UriParamTipeLokasi struct {
	Tipe string `uri:"tipe" binding:"required,min=4"`
}

type UriParamLokasiDepends struct {
	Tipe    string `uri:"tipe" binding:"required,min=4"`
	Depends int32  `uri:"depends" binding:"required,min=1"`
}

type DeleteLokasi struct {
	Id int32 `json:"id" binding:"required,min=1"`
}

type PostNamaLokasi struct {
	Nama string `json:"nama" binding:"required,min=4"`
}
