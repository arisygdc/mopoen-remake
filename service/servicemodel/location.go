package servicemodel

type Provinsi struct {
	ID   int32  `json:"id"`
	Nama string `json:"nama"`
}

type Kabupaten struct {
	ID         int32  `json:"id"`
	ProvinsiID int32  `json:"provinsi_id"`
	Nama       string `json:"nama"`
}

type Kecamatan struct {
	ID          int32  `json:"id"`
	KabupatenID int32  `json:"kabupaten_id"`
	Nama        string `json:"nama"`
}

type Desa struct {
	ID          int32  `json:"id"`
	KecamatanID int32  `json:"kecamatan_id"`
	Nama        string `json:"nama"`
}

type LocationDepends struct {
	ID   int32  `json:"id"`
	Nama string `json:"nama"`
}
