installpg:
	docker run -d --name mopoen-remake-db \
	-p 5432:5432 \
	-e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=qwer1234 \
	-e TZ=Asia/Jakarta -e PGTZ=Asia/Jakarta \
	-e POSTGRES_DB=mopoen \
	postgres:14.5-alpine3.16

uninstallpg:
	docker container rm mopoen-remake-db

startpg:
	docker start mopoen-remake-db

stoppg:
	docker stop mopoen-remake-db

execpg:
	docker exec -it mopoen-remake-db psql -U postgres -d mopoen

backuppg:
	docker exec -it mopoen-remake-db pg_dump -U postgres mopoen > backups/mopoen.sql

restorepg:
	docker exec -i mopoen-remake-db psql -U postgres mopoen < backups/mopoen.sql

createmigrate:
	migrate create -ext sql -dir repository/postgres/migration -seq init_schema

migrateup:
	migrate -path repository/postgres/migration/ -database "postgresql://postgres:qwer1234@localhost:5432/mopoen?sslmode=disable" -verbose up

migratedown:
	migrate -path repository/postgres/migration/ -database "postgresql://postgres:qwer1234@localhost:5432/mopoen?sslmode=disable" -verbose down

.PHONY: installpg uninstallpg startpg stoppg execpg backuppg restorepg createmigrate migrateup migratedown