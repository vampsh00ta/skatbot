POSTGRESQL_URL:='postgresql://skatbot:skatbot@localhost:5432/skatbot?sslmode=disable'

migrate:
	migrate create -ext sql  -dir ./migrations -seq $(SEQ)


migrate-up:
	migrate -database $(POSTGRESQL_URL) -path ./migrations up