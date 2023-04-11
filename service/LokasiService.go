package service

import (
	"context"
	"fmt"
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

func (ls LokasiService) GetParentLokasi(ctx context.Context) ([]postgres.GetParentLokasiRow, error) {
	return ls.repo.GetParentLokasi(ctx)
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

func (ls LokasiService) CreateLokasi(ctx context.Context, tipe servicemodel.LokasiType, nama string, depends ...int32) error {
	if len(depends) > 0 {
		switch tipe {
		case servicemodel.LokKabupaten:
			return ls.repo.CreateKabupaten(ctx, postgres.CreateKabupatenParams{ProvinsiID: depends[0], Nama: nama})
		case servicemodel.LokKecamatan:
			return ls.repo.CreateKecamatan(ctx, postgres.CreateKecamatanParams{KabupatenID: depends[0], Nama: nama})
		case servicemodel.LokDesa:
			return ls.repo.CreateDesa(ctx, postgres.CreateDesaParams{KecamatanID: depends[0], Nama: nama})
		}
	}
	if tipe == servicemodel.LokProvinsi {
		return ls.repo.CreateProvinsi(ctx, nama)
	}
	return fmt.Errorf("invalid lokasi type")
}

func (ls LokasiService) DeleteLokasi(ctx context.Context, tipe servicemodel.LokasiType, id int32) (string, error) {
	switch tipe {
	case servicemodel.LokProvinsi:
		return ls.repo.DeleteProvinsi(ctx, id)
	case servicemodel.LokKabupaten:
		return ls.repo.DeleteKabupaten(ctx, id)
	case servicemodel.LokKecamatan:
		return ls.repo.DeleteKecamatan(ctx, id)
	case servicemodel.LokDesa:
		return ls.repo.DeleteDesa(ctx, id)
	default:
		return "", fmt.Errorf("tipe lokasi tidak valid")
	}
}

func (ls LokasiService) GetLokasiBy(ctx context.Context, tipe servicemodel.LokasiType, depends int32) ([]servicemodel.Lokasi, error) {
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
