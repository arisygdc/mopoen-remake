INSERT INTO tipe_sensor (tipe, satuan) 
VALUES
('angin', 'm/s'),
('air', 'm3/s');

INSERT INTO provinsi (nama) 
VALUES 
('jawa timur');

INSERT INTO kabupaten (id, nama, provinsi_id) 
VALUES 
(DEFAULT, 'batu', (SELECT id FROM provinsi where nama='jawa timur')),
(DEFAULT, 'malang', (SELECT id FROM provinsi where nama='jawa timur'));

INSERT INTO kecamatan (nama, kabupaten_id)
VALUES
('junrejo', (SELECT id FROM kabupaten where nama='batu'));

INSERT INTO desa (nama, kecamatan_id)
VALUES
('Torongrejo', (SELECT id FROM kecamatan where nama='junrejo'));