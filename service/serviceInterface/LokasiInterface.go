package serviceinterface

import (
	"context"
	"mopoen-remake/service/servicemodel"
)

// depends is enum of provinsi, kabupaten, kecamatan, or desa
// provinsi is not depends on anything
type LokasiInterface interface {
	GetAllProvinsi(context.Context) ([]servicemodel.Provinsi, error)
	GetAllKabupaten(context.Context) ([]servicemodel.Kabupaten, error)
	GetAllKecamatan(context.Context) ([]servicemodel.Kecamatan, error)
	GetAllDesa(context.Context) ([]servicemodel.Desa, error)
	// param: context, tipe_lokasi servicemodel.LokasiType, nama string, depends ...int32
	CreateLokasi(context.Context, servicemodel.LokasiType, string, ...int32) error
	// param: context, tipe_lokasi servicemodel.LokasiType, id_lokasi int32
	// return: deleted_lokasi_name string, error
	DeleteLokasi(context.Context, servicemodel.LokasiType, int32) (string, error)
	// param: context, tipe_lokasi servicemodel.LokasiType, depends int32
	GetLokasiBy(context.Context, servicemodel.LokasiType, int32) ([]servicemodel.Lokasi, error)
}
