POSTGRESQL_URL:='postgresql://skatbot:skatbot@localhost:5432/skatbot?sslmode=disable'
PROD_URL = 'postgresql://postgres:ab1C42Ca4dEF4E35-AG22gGc4aB1A4*c@roundhouse.proxy.rlwy.net:31519/railway?sslmode=disable'

migrate:
	migrate create -ext sql  -dir ./migrations -seq $(SEQ)


migrate-up:
	migrate -database $(POSTGRESQL_URL) -path ./migrations up
migrate-up-prod:
	migrate -database $(PROD_URL) -path ./migrations up