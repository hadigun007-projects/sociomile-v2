# Backend commands
run:
	cd backend && go run cmd/api/main.go

build:
	cd backend && go build -o bin/api cmd/api/main.go

test:
	cd backend && go test ./...

# Docker commands
docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

docker-build:
	docker-compose build

docker-logs:
	docker-compose logs -f

docker-restart:
	docker-compose restart

# Database commands
db-create:
	docker-compose exec mysql mysql -u root -p${DB_PASSWORD:-rootpassword} -e "CREATE DATABASE IF NOT EXISTS ${DB_NAME:-sociomile};"

db-drop:
	docker-compose exec mysql mysql -u root -p${DB_PASSWORD:-rootpassword} -e "DROP DATABASE IF EXISTS ${DB_NAME:-sociomile};"

# Clean commands
clean:
	cd backend && rm -rf bin/
	docker-compose down -v

.PHONY: run build test docker-up docker-down docker-build docker-logs docker-restart db-create db-drop clean
