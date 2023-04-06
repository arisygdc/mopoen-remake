package servicemodel

type LokasiType string

const (
	LokProvinsi  LokasiType = "provinsi"
	LokKabupaten LokasiType = "kabupaten"
	LokKecamatan LokasiType = "kecamatan"
	LokDesa      LokasiType = "desa"
)

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

type Lokasi struct {
	ID   int32  `json:"id"`
	Nama string `json:"nama"`
}

type FetchLokasi struct {
	Desa      string `json:"desa"`
	Kecamatan string `json:"kecamatan"`
	Kabupaten string `json:"kabupaten"`
	Provinsi  string `json:"provinsi"`
}
