package serviceinterface

import (
	"context"
	"mopoen-remake/service/servicemodel"
)

type LokasiInterface interface {
	GetAllProvinsi(context.Context) ([]servicemodel.Provinsi, error)
	GetAllKabupaten(context.Context) ([]servicemodel.Kabupaten, error)
	GetAllKecamatan(context.Context) ([]servicemodel.Kecamatan, error)
	GetAllDesa(context.Context) ([]servicemodel.Desa, error)
	// param: context, provinsi name string
	CreateProvinsi(context.Context, string) error
	// param: context, provinsi id int32, kabupaten name string
	CreateKabupaten(context.Context, int32, string) error
	// param: context, kabupaten id int32, kecamatan name string
	CreateKecamatan(context.Context, int32, string) error
	// param: context, kecamatan id int32, desa name string
	CreateDesa(context.Context, int32, string) error
	// param: context, provinsi id int32
	DeleteProvinsi(context.Context, int32) (string, error)
	// param: context, kabupaten id int32
	DeleteKabupaten(context.Context, int32) (string, error)
	// param: context, kecamatan id int32
	DeleteKecamatan(context.Context, int32) (string, error)
	// param: context, desa id int32
	DeleteDesa(context.Context, int32) (string, error)
	// param: context, tipe string, depends int32
	// depends is enum of provinsi, kabupaten, kecamatan, or desa
	GetLokasiBy(context.Context, string, int32) ([]servicemodel.Lokasi, error)
}
