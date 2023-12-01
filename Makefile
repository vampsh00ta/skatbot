POSTGRESQL_URL:='postgresql://skatbot:skatbot@localhost:5432/skatbot?sslmode=disable'
PROD_URL = 'postgresql://postgres:Ae6BcFaBc6eCgB6A5be1EaADbDbGgada@roundhouse.proxy.rlwy.net:53113/railway?sslmode=disable'

migrate:
	migrate create -ext sql  -dir ./migrations -seq $(SEQ)


migrate-up:
	migrate -database $(POSTGRESQL_URL) -path ./migrations up
migrate-up-prod:
	migrate -database $(PROD_URL) -path ./migrations up