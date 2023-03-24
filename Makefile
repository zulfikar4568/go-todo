migration-init:
	atlas migrate hash \
		--dir "file://database/migrations?format=golang-migrate" \
		--dir-format "atlas"

migration-plan:
	atlas migrate diff $(name) \
		--config "file://database/atlas-config.hcl" \
		--env default \
		--var url=$(POSTGRES_DEFAULT_URL) \
		--var dev_url=$(POSTGRES_DEFAULT_TEST_URL)
		
migration-apply:
	atlas migrate apply \
		--config "file://database/atlas-config.hcl" \
		--env default \
		--var url=$(POSTGRES_DEFAULT_URL) \
		--var dev_url=$(POSTGRES_DEFAULT_TEST_URL)

migration-clean:
	atlas schema clean \
		--config "file://database/atlas-config.hcl" \
		--env default \
		--var url=$(POSTGRES_DEFAULT_URL) \
		--var dev_url=$(POSTGRES_DEFAULT_TEST_URL)

migration-rehash:
	atlas migrate hash --dir "file://database/migrations"

test:
	go test ./...

lint:
	golangci-lint run