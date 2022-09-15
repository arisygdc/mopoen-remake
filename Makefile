installpg:
	docker run -d --name mopoen-remake-db \
	-p 5432:5432 \
	-e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=qwer1234 \
	-e TZ=Asia/Jakarta -e PGTZ=Asia/Jakarta \
	-e POSTGRES_DB=mopoen \
	postgres:12-alpine3.14

uninstallpg:
	docker container rm mopoen-remake-db

startpg:
	docker start mopoen-remake-db

stoppg:
	docker stop mopoen-remake-db

execpg:
	docker exec -it mopoen-remake-db psql -U postgres

createmigrate:
	migrate create -ext sql -dir database/postgres/migration -seq init_schema

migrateup:
	migrate -path database/postgres/migration/ -database "postgresql://postgres:qwer1234@localhost:5432/mopoen?sslmode=disable" -verbose up

migratedown:
	migrate -path database/postgres/migration/ -database "postgresql://postgres:qwer1234@localhost:5432/mopoen?sslmode=disable" -verbose down

.PHONY: installpg uninstallpg startpg stoppg execdb createmigrate migrateup migratedown