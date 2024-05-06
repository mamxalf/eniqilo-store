include .env

MIGRATION_STEP=1
DB_CONN=postgres://$(DB_USERNAME:"%"=%):$(DB_PASSWORD:"%"=%)@$(DB_HOST:"%"=%):$(DB_PORT:"%"=%)/$(DB_NAME:"%"=%)?$(DB_PARAMS:"%"=%)

dev: generate
	go run github.com/cosmtrek/air

run: generate
	go run .

lint-prepare:
	@echo "Installing golangci-lint"
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s latest

lint:
	./bin/golangci-lint run --config .golangci.yml ./... --fix

generate:
	go generate ./...

migrate_create:
	@read -p "migration name (do not use space): " NAME \
  	&& migrate create -ext sql -dir ./db/migrations $${NAME}

migrate_up:
	@migrate -path ./db/migrations -database "$(DB_CONN)" up $(MIGRATION_STEP)

migrate_down:
	@migrate -path ./db/migrations -database "$(DB_CONN)" down $(MIGRATION_STEP)

migrate_version:
	@migrate -path ./db/migrations -database "$(DB_CONN)" version

migrate_drop:
	@migrate -path ./db/migrations -database "$(DB_CONN)" drop

migrate_force:
	@read -p "please enter the migration version (the migration filename prefix): " VERSION \
  	&& migrate -path ./db/migrations -database "$(DB_CONN)" force $${VERSION}

docker_dev:
	@docker compose up -d

docker_dev_off:
	@docker compose down