package service

import (
	"context"
	"mopoen-remake/repository"
	"mopoen-remake/repository/postgres"
	"mopoen-remake/service/servicemodel"
)

type LokasiService struct {
	repo repository.Repository
}

func NewLokasiService(repo repository.Repository) LokasiService {
	return LokasiService{repo: repo}
}

func (ls LokasiService) GetAllProvinsi(ctx context.Context) ([]servicemodel.Provinsi, error) {
	prov, err := ls.repo.GetAllProvinsi(ctx)
	if err != nil {
		return nil, err
	}

	provinsiConverted := make([]servicemodel.Provinsi, len(prov))
	for i, v := range prov {
		provinsiConverted[i] = servicemodel.Provinsi(v)
	}

	return provinsiConverted, nil
}

func (ls LokasiService) GetAllKabupaten(ctx context.Context) ([]servicemodel.Kabupaten, error) {
	kab, err := ls.repo.GetAllKabupaten(ctx)
	if err != nil {
		return nil, err
	}

	kabupatenConverted := make([]servicemodel.Kabupaten, len(kab))
	for i, v := range kab {
		kabupatenConverted[i] = servicemodel.Kabupaten(v)
	}

	return kabupatenConverted, nil
}

func (ls LokasiService) GetAllKecamatan(ctx context.Context) ([]servicemodel.Kecamatan, error) {
	kec, err := ls.repo.GetAllKecamatan(ctx)
	if err != nil {
		return nil, err
	}

	kecamatanConverted := make([]servicemodel.Kecamatan, len(kec))
	for i, v := range kec {
		kecamatanConverted[i] = servicemodel.Kecamatan(v)
	}

	return kecamatanConverted, nil
}

func (ls LokasiService) GetAllDesa(ctx context.Context) ([]servicemodel.Desa, error) {
	des, err := ls.repo.GetAllDesa(ctx)
	if err != nil {
		return nil, err
	}

	DesaConverted := make([]servicemodel.Desa, len(des))
	for i, v := range des {
		DesaConverted[i] = servicemodel.Desa(v)
	}

	return DesaConverted, nil
}

func (ls LokasiService) CreateProvinsi(ctx context.Context, provinsi string) error {
	return ls.repo.CreateProvinsi(ctx, provinsi)
}

func (ls LokasiService) CreateKabupaten(ctx context.Context, provinsi_id int32, kabupaten string) error {
	param := postgres.CreateKabupatenParams{ProvinsiID: provinsi_id, Nama: kabupaten}
	return ls.repo.CreateKabupaten(ctx, param)
}

func (ls LokasiService) CreateKecamatan(ctx context.Context, kabupaten_id int32, kecamatan string) error {
	param := postgres.CreateKecamatanParams{KabupatenID: kabupaten_id, Nama: kecamatan}
	return ls.repo.CreateKecamatan(ctx, param)
}

func (ls LokasiService) CreateDesa(ctx context.Context, kecamatan_id int32, desa string) error {
	param := postgres.CreateDesaParams{KecamatanID: kecamatan_id, Nama: desa}
	return ls.repo.CreateDesa(ctx, param)
}

func (ls LokasiService) DeleteProvinsi(ctx context.Context, provinsi_id int32) (string, error) {
	return ls.repo.DeleteProvinsi(ctx, provinsi_id)
}

func (ls LokasiService) DeleteKabupaten(ctx context.Context, kabupaten_id int32) (string, error) {
	return ls.repo.DeleteKabupaten(ctx, kabupaten_id)
}

func (ls LokasiService) DeleteKecamatan(ctx context.Context, kecamatan_id int32) (string, error) {
	return ls.repo.DeleteKecamatan(ctx, kecamatan_id)
}

func (ls LokasiService) DeleteDesa(ctx context.Context, desa_id int32) (string, error) {
	return ls.repo.DeleteDesa(ctx, desa_id)
}

func (ls LokasiService) GetLokasiBy(ctx context.Context, tipe string, depends int32) ([]servicemodel.Lokasi, error) {
	var locationBy []servicemodel.Lokasi
	switch tipe {
	case servicemodel.LokKabupaten:
		val, Qerr := ls.repo.GetKabupatenBy(ctx, depends)
		if Qerr != nil {
			return nil, Qerr
		}

		locationBy = make([]servicemodel.Lokasi, len(val))
		for i, val := range val {
			locationBy[i] = servicemodel.Lokasi{
				ID:   val.ID,
				Nama: val.Nama,
			}
		}

	case servicemodel.LokKecamatan:
		val, Qerr := ls.repo.GetKecamatanBy(ctx, depends)
		if Qerr != nil {
			return nil, Qerr
		}

		locationBy = make([]servicemodel.Lokasi, len(val))
		for i, val := range val {
			locationBy[i] = servicemodel.Lokasi{
				ID:   val.ID,
				Nama: val.Nama,
			}
		}

	case servicemodel.LokDesa:
		val, Qerr := ls.repo.GetDesaBy(ctx, depends)
		if Qerr != nil {
			return nil, Qerr
		}

		locationBy = make([]servicemodel.Lokasi, len(val))
		for i, val := range val {
			locationBy[i] = servicemodel.Lokasi{
				ID:   val.ID,
				Nama: val.Nama,
			}
		}

	}

	return locationBy, nil
}
