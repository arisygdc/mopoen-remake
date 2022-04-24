package pgservice

import (
	"context"
	"database/sql"
	"mopoen-remake/database/postgres"
	"mopoen-remake/service/servicemodel"

	_ "github.com/lib/pq"
)

type postgre struct {
	SQLConn *sql.DB
	Queries *postgres.Queries
}

func NewPostgres(driver string, source string) (postgre, error) {
	SQLConn, err := sql.Open(driver, source)
	if err != nil {
		return postgre{}, err
	}

	pgConn := postgre{
		SQLConn: SQLConn,
		Queries: postgres.New(SQLConn),
	}
	return pgConn, nil
}

func (db postgre) CreateTipeSensor(ctx context.Context, tipe string, satuan string) error {
	param := postgres.InsertTipeSensorParams{Tipe: tipe, Satuan: satuan}
	return db.Queries.InsertTipeSensor(ctx, param)

}

func (db postgre) GetTipeSensor(ctx context.Context, id int32) (servicemodel.TipeSensor, error) {
	sensor, err := db.Queries.GetTipeSensor(ctx, id)
	return servicemodel.TipeSensor(sensor), err
}

func (db postgre) GetAllTipeSensor(ctx context.Context) ([]servicemodel.TipeSensor, error) {
	sensors, err := db.Queries.GetTipeSensors(ctx)
	if err != nil {
		return nil, err
	}

	sensorsConverted := make([]servicemodel.TipeSensor, len(sensors))
	for i, v := range sensors {
		sensorsConverted[i] = servicemodel.TipeSensor(v)
	}

	return sensorsConverted, nil
}

func (db postgre) GetAllProvinsi(ctx context.Context) ([]servicemodel.Provinsi, error) {
	prov, err := db.Queries.GetAllProvinsi(ctx)
	if err != nil {
		return nil, err
	}

	provinsiConverted := make([]servicemodel.Provinsi, len(prov))
	for i, v := range prov {
		provinsiConverted[i] = servicemodel.Provinsi(v)
	}

	return provinsiConverted, nil
}

func (db postgre) GetAllKabupaten(ctx context.Context) ([]servicemodel.Kabupaten, error) {
	kab, err := db.Queries.GetAllKabupaten(ctx)
	if err != nil {
		return nil, err
	}

	kabupatenConverted := make([]servicemodel.Kabupaten, len(kab))
	for i, v := range kab {
		kabupatenConverted[i] = servicemodel.Kabupaten(v)
	}

	return kabupatenConverted, nil
}

func (db postgre) GetAllKecamatan(ctx context.Context) ([]servicemodel.Kecamatan, error) {
	kec, err := db.Queries.GetAllKecamatan(ctx)
	if err != nil {
		return nil, err
	}

	kecamatanConverted := make([]servicemodel.Kecamatan, len(kec))
	for i, v := range kec {
		kecamatanConverted[i] = servicemodel.Kecamatan(v)
	}

	return kecamatanConverted, nil
}

func (db postgre) GetAllDesa(ctx context.Context) ([]servicemodel.Desa, error) {
	des, err := db.Queries.GetAllDesa(ctx)
	if err != nil {
		return nil, err
	}

	DesaConverted := make([]servicemodel.Desa, len(des))
	for i, v := range des {
		DesaConverted[i] = servicemodel.Desa(v)
	}

	return DesaConverted, nil
}

func (db postgre) DeleteTipeSensor(ctx context.Context, id int32) error {
	return db.Queries.DeleteTipeSensor(ctx, id)
}

func (db postgre) CreateProvinsi(ctx context.Context, provinsi string) error {
	return db.Queries.CreateProvinsi(ctx, provinsi)
}

func (db postgre) CreateKabupaten(ctx context.Context, provinsi_id int32, kabupaten string) error {
	param := postgres.CreateKabupatenParams{ProvinsiID: provinsi_id, Nama: kabupaten}
	return db.Queries.CreateKabupaten(ctx, param)
}

func (db postgre) CreateKecamatan(ctx context.Context, kabupaten_id int32, kecamatan string) error {
	param := postgres.CreateKecamatanParams{KabupatenID: kabupaten_id, Nama: kecamatan}
	return db.Queries.CreateKecamatan(ctx, param)
}

func (db postgre) CreateDesa(ctx context.Context, kecamatan_id int32, desa string) error {
	param := postgres.CreateDesaParams{KecamatanID: kecamatan_id, Nama: desa}
	return db.Queries.CreateDesa(ctx, param)
}

func (db postgre) DeleteProvinsi(ctx context.Context, provinsi_id int32) (string, error) {
	return db.Queries.DeleteProvinsi(ctx, provinsi_id)
}

func (db postgre) DeleteKabupaten(ctx context.Context, kabupaten_id int32) (string, error) {
	return db.Queries.DeleteKabupaten(ctx, kabupaten_id)
}

func (db postgre) DeleteKecamatan(ctx context.Context, kecamatan_id int32) (string, error) {
	return db.Queries.DeleteKecamatan(ctx, kecamatan_id)
}

func (db postgre) DeleteDesa(ctx context.Context, desa_id int32) (string, error) {
	return db.Queries.DeleteDesa(ctx, desa_id)
}

func (db postgre) GetLokasiBy(ctx context.Context, tipe string, depends int32) ([]servicemodel.Lokasi, error) {
	var locationBy []servicemodel.Lokasi
	switch tipe {
	case "kabupaten":
		val, Qerr := db.Queries.GetKabupatenBy(ctx, depends)
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
		val, Qerr := db.Queries.GetKecamatanBy(ctx, depends)
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
		val, Qerr := db.Queries.GetDesaBy(ctx, depends)
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
