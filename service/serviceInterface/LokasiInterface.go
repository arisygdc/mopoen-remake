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
	CreateKecamatan(ctx context.Context, kabupaten_id int32, kecamatan string) error
	// param: context, kecamatan id int32, desa name string
	CreateDesa(ctx context.Context, kecamatan_id int32, desa string) error
	// param: context, provinsi id int32
	DeleteProvinsi(ctx context.Context, provinsi_id int32) (string, error)
	// param: context, kabupaten id int32
	DeleteKabupaten(ctx context.Context, kabupaten_id int32) (string, error)
	// param: context, kecamatan id int32
	DeleteKecamatan(ctx context.Context, kecamatan_id int32) (string, error)
	// param: context, desa id int32
	DeleteDesa(ctx context.Context, desa_id int32) (string, error)
	// param: context, tipe string, depends int32
	// depends is enum of provinsi, kabupaten, kecamatan, or desa
	GetLokasiBy(ctx context.Context, tipe string, depends int32) ([]servicemodel.Lokasi, error)
}
