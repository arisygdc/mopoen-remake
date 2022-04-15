package pgrepo

import (
	"context"
	"database/sql"
	"mopoen-remake/config"
	"mopoen-remake/database/postgres"
	"mopoen-remake/service/servicemodel"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type postgre struct {
	SQLConn *sql.DB
	Queries *postgres.Queries
}

func NewPostgres(env config.Environment) (pgConn postgre, err error) {
	SQLConn, err := sql.Open(env.DBDriver, env.DBSource)
	if err != nil {
		return
	}

	pgConn = postgre{
		SQLConn: SQLConn,
		Queries: postgres.New(SQLConn),
	}
	return
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
	var sensorsConverted []servicemodel.TipeSensor
	for i, v := range sensors {
		sensorsConverted[i] = servicemodel.TipeSensor(v)
	}
	return sensorsConverted, err
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

func (db postgre) DeleteProvinsi(ctx context.Context, provinsi_id int32) error {
	return db.Queries.DeleteProvinsi(ctx, provinsi_id)
}

func (db postgre) DeleteKabupaten(ctx context.Context, kabupaten_id int32) error {
	return db.Queries.DeleteKabupaten(ctx, kabupaten_id)
}

func (db postgre) DeleteKecamatan(ctx context.Context, kecamatan_id int32) error {
	return db.Queries.DeleteKecamatan(ctx, kecamatan_id)
}

func (db postgre) DeleteDesa(ctx context.Context, desa_id int32) error {
	return db.Queries.DeleteDesa(ctx, desa_id)
}

func (db postgre) DaftarMonitoring(ctx context.Context, daftarMonitoringParam servicemodel.DaftarMonitoring) error {
	param := postgres.CreateMonitoringTerdaftarParams{
		ID:           uuid.New(),
		TipeSensorID: daftarMonitoringParam.TipeSensor,
		LokasiID:     daftarMonitoringParam.Location_id,
		Nama:         daftarMonitoringParam.Nama,
		Keterangan:   daftarMonitoringParam.Keterangan,
	}

	return db.Queries.CreateMonitoringTerdaftar(ctx, param)
}

func (db postgre) CreateMonitoringValue(ctx context.Context, monitoring_id uuid.UUID, value float64) error {
	param := postgres.CreateMonitoringValueParams{
		MonitoringTerdaftar: monitoring_id,
		Value:               value,
	}

	return db.Queries.CreateMonitoringValue(ctx, param)
}
