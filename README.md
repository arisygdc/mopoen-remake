# MONITORING POTENSI ENERGI REMAKE
This project is remake version of [mopoen](https://github.com/arisygdc/mopoen-remake), trying to explore my knowlege about golang, postgresql, and coding style.


## API Spec
### Post and Error response
```JSON
// object returns
{
    "messages": {

    }
}
```
And
```JSON
{
    "message": "string return"
}
```

### Get response

```JSON
// Object returns
{
    "data": {
        
    }
}
```
AND
```JSON
// Array returns
{
    "data": [

    ]
}
```
AND
```JSON
{
    "data": "string return"
}
```

docker run -d --name mopoen-remake-db \
	-p 5432:5432 \
	-e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=qwer1234 \
	-e TZ=Asia/Jakarta -e PGTZ=Asia/Jakarta \
	-e POSTGRES_DB=mopoen \
    --network mopoen \
	postgres:12-alpine3.14

docker run -d --name mopoen-s1 -p 8081:8080 --network mopoen -e DATABASE_SOURCE=postgresql://postgres:qwer1234@mopoen-db-release:5432/mopoen?sslmode=disable bf27107e9ea4