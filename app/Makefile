BINARY_NAMES="golang-migrate|migrate"
DATABASE=postgres://postgres:postgres@localhost:5432/test?sslmode=disabl

help: ## Displays the help for each command.
	@grep -E '^[a-zA-Z_-]+:.*## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

dbseq: ## Creates a new migration file with the given name. ex: name=create_users_table
	migrate create -ext sql -dir migrations $(name)

dbu: ## Migrates the database to the latest version.
	migrate -database="$(DATABASE)" -path=migrations -lock-timeout=20  -verbose up

dbd: ## Migrates the database down by first version.
	migrate -database="$(DATABASE)" -path=migrations -lock-timeout=20 -verbose down

dbdrop: ## Drops the database.
	migrate -database="$(DATABASE)" -path=migrations -lock-timeout=20 -verbose drop -f