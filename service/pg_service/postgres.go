package pgrepo

import (
	"context"
	"database/sql"
	"mopoen-remake/config"
	"mopoen-remake/database/postgres"
	"mopoen-remake/service/servicemodel"

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

func (db postgre) DeleteTipeSensor(ctx context.Context, id int32) error {
	return db.Queries.DeleteTipeSensor(ctx, id)
}

func (db postgre) GetAllTipeSensor(ctx context.Context) ([]servicemodel.TipeSensor, error) {
	sensors, err := db.Queries.GetTipeSensors(ctx)
	var sensorsConverted []servicemodel.TipeSensor
	for i, v := range sensors {
		sensorsConverted[i] = servicemodel.TipeSensor(v)
	}
	return sensorsConverted, err
}

func (db postgre) GetTipeSensor(ctx context.Context, id int32) (servicemodel.TipeSensor, error) {
	sensor, err := db.Queries.GetTipeSensor(ctx, id)
	return servicemodel.TipeSensor(sensor), err
}
