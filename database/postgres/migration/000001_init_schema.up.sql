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
    nama VARCHAR(70) NOT NULL,
    keterangan VARCHAR(200) NOT NULL
);

CREATE TABLE monitoring_data (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    monitoring_terdaftar uuid NOT NULL,
    value INT NOT NULL
);