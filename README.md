# MONITORING POTENSI ENERGI REMAKE
This project is remake version of [mopoen](https://github.com/arisygdc/mopoen-remake), trying to explore my knowlege about golang, postgresql, and coding style.


## API Documentation
Welcome to the API documentation for our sensing value endpoint. Here, you can learn how to interact with our API.

### POST sensing value
To post sensing values, you can make a POST request to the `/api/sensor/value` endpoint.

**Request Body**

The request body must be a JSON object with the following fields:
|Field Name     | Type  | Required|Description|
|---------------|-------|---------|-----------|
|kode_monitoring| string|	Yes   |	The ID of the monitoring.|
|value          | number|	Yes   |	The sensing value.|

**Responses**

- 201 Created if the sensing value submission was successful.
```JSON
{}
```
Upon successful submission, the API will return a 201 status code along with an empty JSON response.

### GET lokasi request
To request location data, you can make a GET request to the endpoint `/api/v1/lokasi/:tipe`.
<!-- list -->
- Type can be filled with values `provinsi`, `kabupaten`, `kecamatan`, or `desa`.
- Example: `/api/v1/lokasi/provinsi`
<!-- end of the list -->
Here are the examples of response for each type:

**Provinsi**
```JSON
{
    "data": [
        {
            "id": 1,
            "nama": "jawa timur"
        }
    ]
}
```
**Kabupaten**
```JSON

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

### GET lokasi parent
To request parent location data, you can make a GET request to the endpoint `/api/v1/lokasi/parent`.

This API generates the parent location data of the selected location. For example, if the selected location is a village, then this API will generate the data of the province, district, and sub-district from that village.

Here is an example response:
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

### GET Available Sensor Types Request
To request available sensor type data, you can make a GET request to the endpoint `/api/v1/sensors`.

This API generates the available sensor types and their units of measurement.

Here is an example response:
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

### GET Sensor Type by ID Request
To get the sensor type by ID, you can make a GET request to the endpoint `/api/v1/sensor/:id`. The `:id` parameter should be replaced with the ID of the desired sensor type.

If the sensor ID is available, the API will return a JSON response with a status code of 200, and the details of the sensor type. If the ID is not found, the API will return a status code of 204 with an empty response body.

Here is an example response for a valid sensor ID:
```JSON
{
    "data": {
        "id": 1,
        "tipe": "angin",
        "satuan": "m/s"
    }
}
```

### POST daftar monitoring
`/api/v1/monitoring/daftar` This endpoint registers a user for monitoring, and send information to email.

**Request Body**

The request body must be a JSON object with the following fields:
|Field Name     | Type | Required | Description |
|---------------|------|----------|-------------|
|tipe_sensor_id |uuid  | Yes| The ID of the sensor type.|
|lokasi_id      |uuid  | Yes| The ID of the location.|
|email	        |string| Yes| The user's email address.|
|author	        |string| Yes| The author of the monitoring.|
|nama	        |string| Yes| The name of the monitoring.|
|keterangan     |string| No | Additional information about the monitoring.|

**Responses**

- 201 Created if the monitoring registration was successful.
```JSON
{
    "message": "analisa angin torong rejo created"
}
```
- 400 Bad Request if the request body is missing required fields or contains invalid data.
```JSON
{
    "message": "invalid character"
}
```
### Get monitoring terdaftar
`/api/v1/monitoring/terdaftar`

This endpoint allows you to retrieve a list of registered monitoring.

**Query Parameters**
|Parameter Name|	Type|	Required|	Description|
|--------------|--------|-----------|--------------|
|sensor_id|	number|	No|	Filter the monitoring by sensor type ID.|
|lokasi_id|	number|	No|	Filter the monitoring by location ID.|
**Response**
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

### GET monitoring terdaftar by id
To retrieve information about a specific monitoring registration, you can make a GET request to the `/api/v1/monitoring/terdaftar/:id` endpoint, where `:id` should be replaced with the ID of the monitoring registration you want to retrieve. The API will respond with a 200 status code and return the details of the monitoring registration in JSON format, including the sensor type, location ID, name, and description.


Here's an example response for the `/api/v1/monitoring/terdaftar/d7e6ec83-1549-46bf-bdc0-0f7f1d3e23c5` endpoint:
```JSON
{
    "tipe_sensor": 2,
    "lokasi_id": 1,
    "nama": "analisa air torong rejo",
    "keterangan": "untuk melakukan penelitian skripsi"
}
```

### GET value monitoring
To retrieve all the monitoring values associated with a specific monitoring registration, you can make a GET request to the `/api/v1/monitoring/value/:uuid` endpoint, where `:uuid` should be replaced with the UUID of the monitoring registration. The API will respond with a 200 status code and return the sensing values in JSON format, including the value and timestamp.

Here's an example response for the `/api/v1/monitoring/value/d7e6ec83-1549-46bf-bdc0-0f7f1d3e23c5` endpoint:
```JSON
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
## How to run in your local machine
The following instructions explain how to run the Mopoen Remake application using `Docker`:

- Run a PostgreSQL database container with the following command:
```bash
$ docker run -d --name mopoen-remake-db \
	-p 5432:5432 \
	-e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=qwer1234 \
	-e TZ=Asia/Jakarta -e PGTZ=Asia/Jakarta \
	-e POSTGRES_DB=mopoen \
    --network mopoen \
	postgres:12-alpine3.14
```

Create a Docker container named **mopoen-remake-db** running PostgreSQL version 12. The container will be connected to a Docker network named **mopoen**. You can change the PostgreSQL username, password, and database name by modifying the environment variables.

- Run the Mopoen Remake application container with the following command:
```bash
$ docker run -d --name mopoen-s1 -p 8080:8080 --network mopoen -e DATABASE_SOURCE=postgresql://postgres:qwer1234@mopoen-db-release:5432/mopoen?sslmode=disable
```
Create a Docker container named **mopoen-s1** running the Mopoen Remake application. The container will be connected to the mopoen Docker network. The environment variable `DATABASE_SOURCE` specifies the PostgreSQL database connection string. You can modify this value to match your PostgreSQL database configuration. The container will expose the Mopoen Remake application on port **8080**, which you can access from your host machine.

That's it! You should now be able to access the Mopoen Remake API by using http://localhost:8080.