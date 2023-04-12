# MONITORING POTENSI ENERGI REMAKE
This project is remake version of [mopoen](https://github.com/arisygdc/mopoen-remake), trying to explore my knowlege about golang, postgresql, and coding style.


## API Documentation

### Post sensing value
/api/sensor/value <br>
Request body
```JSON
// Response 201 with empty json
{
    "kode_monitoring": "d7e6ec83-1549-46bf-bdc0-0f7f1d3e23c5",
    "value": 23
}
```

### Get lokasi request
/api/v1/lokasi/:tipe <br />
tipe can be provinsi | kabupaten | kecamatan | desa
```JSON
// Provinsi
{
    "data": [
        {
            "id": 1,
            "nama": "jawa timur"
        }
    ]
}
// Kabupaten
{
    "data": [
        {
            "id": 1,
            "provinsi_id": 1,
            "nama": "batu"
        },
        {
            "id": 2,
            "provinsi_id": 1,
            "nama": "malang"
        }
    ]
}
```

### Get lokasi parent
/api/v1/lokasi/parent
```JSON
{
    "data": [
        {
            "id": 1,
            "nama": "Torongrejo, junrejo, batu, jawa timur"
        }
    ]
}
```

### Get tipe sensor request
/api/v1/sensors
```JSON
{
    "data": [
        {
            "id": 1,
            "tipe": "angin",
            "satuan": "m/s"
        },
        {
            "id": 2,
            "tipe": "air",
            "satuan": "m/s"
        }
    ]
}
```

### Get tipe sensor by id request
/api/v1/sensor/:id <br>
Response 200 if success and 204 not found
```JSON
{
    "data": {
        "id": 1,
        "tipe": "angin",
        "satuan": "m/s"
    }
}
```

### Post daftar monitoring
/api/v1/monitoring/daftar
Request body
```JSON
{
    "tipe_sensor_id": 1,
    "lokasi_id": 1,
    "email": "somemail@mail.com",
    "author": "author",
    "nama": "analisa angin torong rejo",
    "keterangan": "untuk melakukan penelitian skripsi"
}
```
Response 201 if success
```JSON
{
    "message": "analisa angin torong rejo created"
}
```
Response 400 when validation error
```JSON
{
    "message": "invalid character"
}
```
### Get monitoring terdaftar
/api/v1/monitoring/terdaftar <br>
Response 200 <br>
query param: sensor_id, lokasi_id
```JSON
{
    "data": [
        {
            "id": "d7e6ec83-1549-46bf-bdc0-0f7f1d3e23c5",
            "tipe_sensor_id": 1,
            "tipe_sensor": "angin (m/s)",
            "nama": "analisa angin torong rejo",
            "keterangan": "untuk melakukan penelitian skripsi",
            "address":	"jawa timur, batu, junrejo, Torongrejo"
        },
        {
            "id": "d7b8caf5-f3f0-41d8-9506-4e10f947a742",
            "tipe_sensor_id": 2,
            "tipe_sensor": "air (m/s)",
            "nama": "analisa air torong rejo",
            "keterangan": "untuk melakukan penelitian skripsi",
            "address":	"jawa timur, batu, junrejo, Torongrejo"
        }
    ]
}
```

### Get monitoring terdaftar by id
/api/v1/monitoring/terdaftar/:id <br>
Response 200
```JSON
// example /api/v1/monitoring/terdaftar/d7e6ec83-1549-46bf-bdc0-0f7f1d3e23c5
{
    "tipe_sensor": 2,
    "lokasi_id": 1,
    "nama": "analisa air torong rejo",
    "keterangan": "untuk melakukan penelitian skripsi"
}
```

### Get value monitoring
monitoring/value/:uuid <br>
Response 200
```JSON
// example /api/v1/monitoring/value/d7e6ec83-1549-46bf-bdc0-0f7f1d3e23c5
{
    "data": [
        {
            "value": 23,
            "dibuat_pada": "2023-04-07T16:27:30.01939Z"
        },
        {
            "value": 22,
            "dibuat_pada": "2023-04-07T16:27:47.587209Z"
        },
        {
            "value": 23,
            "dibuat_pada": "2023-04-07T16:27:50.534785Z"
        },
        {
            "value": 24,
            "dibuat_pada": "2023-04-07T16:27:54.885099Z"
        },
        {
            "value": 23,
            "dibuat_pada": "2023-04-07T16:28:02.050358Z"
        }
    ]
}
```
## How to run
```bash
$ docker run -d --name mopoen-remake-db \
	-p 5432:5432 \
	-e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=qwer1234 \
	-e TZ=Asia/Jakarta -e PGTZ=Asia/Jakarta \
	-e POSTGRES_DB=mopoen \
    --network mopoen \
	postgres:12-alpine3.14

$ docker run -d --name mopoen-s1 -p 8081:8080 --network mopoen -e DATABASE_SOURCE=postgresql://postgres:qwer1234@mopoen-db-release:5432/mopoen?sslmode=disable bf27107e9ea4
```