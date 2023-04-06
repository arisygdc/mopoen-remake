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

func (db LokasiService) GetAllProvinsi(ctx context.Context) ([]servicemodel.Provinsi, error) {
	prov, err := db.repo.GetAllProvinsi(ctx)
	if err != nil {
		return nil, err
	}

	provinsiConverted := make([]servicemodel.Provinsi, len(prov))
	for i, v := range prov {
		provinsiConverted[i] = servicemodel.Provinsi(v)
	}

	return provinsiConverted, nil
}

func (db LokasiService) GetAllKabupaten(ctx context.Context) ([]servicemodel.Kabupaten, error) {
	kab, err := db.repo.GetAllKabupaten(ctx)
	if err != nil {
		return nil, err
	}

	kabupatenConverted := make([]servicemodel.Kabupaten, len(kab))
	for i, v := range kab {
		kabupatenConverted[i] = servicemodel.Kabupaten(v)
	}

	return kabupatenConverted, nil
}

func (db LokasiService) GetAllKecamatan(ctx context.Context) ([]servicemodel.Kecamatan, error) {
	kec, err := db.repo.GetAllKecamatan(ctx)
	if err != nil {
		return nil, err
	}

	kecamatanConverted := make([]servicemodel.Kecamatan, len(kec))
	for i, v := range kec {
		kecamatanConverted[i] = servicemodel.Kecamatan(v)
	}

	return kecamatanConverted, nil
}

func (db LokasiService) GetAllDesa(ctx context.Context) ([]servicemodel.Desa, error) {
	des, err := db.repo.GetAllDesa(ctx)
	if err != nil {
		return nil, err
	}

	DesaConverted := make([]servicemodel.Desa, len(des))
	for i, v := range des {
		DesaConverted[i] = servicemodel.Desa(v)
	}

	return DesaConverted, nil
}

func (db LokasiService) CreateProvinsi(ctx context.Context, provinsi string) error {
	return db.repo.CreateProvinsi(ctx, provinsi)
}

func (db LokasiService) CreateKabupaten(ctx context.Context, provinsi_id int32, kabupaten string) error {
	param := postgres.CreateKabupatenParams{ProvinsiID: provinsi_id, Nama: kabupaten}
	return db.repo.CreateKabupaten(ctx, param)
}

func (db LokasiService) CreateKecamatan(ctx context.Context, kabupaten_id int32, kecamatan string) error {
	param := postgres.CreateKecamatanParams{KabupatenID: kabupaten_id, Nama: kecamatan}
	return db.repo.CreateKecamatan(ctx, param)
}

func (db LokasiService) CreateDesa(ctx context.Context, kecamatan_id int32, desa string) error {
	param := postgres.CreateDesaParams{KecamatanID: kecamatan_id, Nama: desa}
	return db.repo.CreateDesa(ctx, param)
}

func (db LokasiService) DeleteProvinsi(ctx context.Context, provinsi_id int32) (string, error) {
	return db.repo.DeleteProvinsi(ctx, provinsi_id)
}

func (db LokasiService) DeleteKabupaten(ctx context.Context, kabupaten_id int32) (string, error) {
	return db.repo.DeleteKabupaten(ctx, kabupaten_id)
}

func (db LokasiService) DeleteKecamatan(ctx context.Context, kecamatan_id int32) (string, error) {
	return db.repo.DeleteKecamatan(ctx, kecamatan_id)
}

func (db LokasiService) DeleteDesa(ctx context.Context, desa_id int32) (string, error) {
	return db.repo.DeleteDesa(ctx, desa_id)
}

func (db LokasiService) GetLokasiBy(ctx context.Context, tipe string, depends int32) ([]servicemodel.Lokasi, error) {
	var locationBy []servicemodel.Lokasi
	switch tipe {
	case "kabupaten":
		val, Qerr := db.repo.GetKabupatenBy(ctx, depends)
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

	case "kecamatan":
		val, Qerr := db.repo.GetKecamatanBy(ctx, depends)
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

	case "desa":
		val, Qerr := db.repo.GetDesaBy(ctx, depends)
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
