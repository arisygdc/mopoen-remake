CREATE TABLE tipe_sensor (
    id SERIAL NOT NULL PRIMARY KEY,
    tipe VARCHAR(50) NOT NULL,
    satuan VARCHAR(4) NOT NULL
);

CREATE TABLE provinsi (
    id SERIAL NOT NULL PRIMARY KEY,
    nama VARCHAR(50) NOT NULL
);

CREATE TABLE kabupaten (
    id SERIAL NOT NULL PRIMARY KEY,
    provinsi_id INT NOT NULL,
    nama VARCHAR(50) NOT NULL
);

CREATE TABLE kecamatan (
    id SERIAL NOT NULL PRIMARY KEY,
    kabupaten_id INT NOT NULL,
    nama VARCHAR(50) NOT NULL
);

CREATE TABLE desa (
    id SERIAL NOT NULL PRIMARY KEY,
    kecamatan_id INT NOT NULL,
    nama VARCHAR(50) NOT NULL
);

CREATE TABLE monitoring_terdaftar (
    id uuid NOT NULL PRIMARY KEY,
    tipe_sensor_id INT NOT NULL,
    lokasi_id INT NOT NULL,
    email VARCHAR(50) NOT NULL,
    author VARCHAR(50) NOT NULL,
    secret VARCHAR(33) NOT NULL,
    nama VARCHAR(70) NOT NULL,
    keterangan VARCHAR(200) NOT NULL
);

CREATE TABLE monitoring_data (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    monitoring_terdaftar uuid NOT NULL,
    value FLOAT NOT NULL,
    dibuat_pada TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Relation table
ALTER TABLE kabupaten ADD FOREIGN KEY (provinsi_id) REFERENCES provinsi(id);
ALTER TABLE kecamatan ADD FOREIGN KEY (kabupaten_id) REFERENCES kabupaten(id);
ALTER TABLE desa ADD FOREIGN KEY (kecamatan_id) REFERENCES kecamatan(id);

ALTER TABLE monitoring_terdaftar ADD FOREIGN KEY (tipe_sensor_id) REFERENCES tipe_sensor(id);
ALTER TABLE monitoring_terdaftar ADD FOREIGN KEY (lokasi_id) REFERENCES desa(id);

ALTER TABLE monitoring_data ADD FOREIGN KEY (monitoring_terdaftar) REFERENCES monitoring_terdaftar(id);

-- Index reference
CREATE INDEX ON kabupaten (provinsi_id);
CREATE INDEX ON kecamatan (kabupaten_id);
CREATE INDEX ON desa (kecamatan_id);

CREATE INDEX ON monitoring_terdaftar (tipe_sensor_id);
CREATE INDEX ON monitoring_terdaftar (lokasi_id);

CREATE INDEX ON monitoring_data (monitoring_terdaftar);