create migration file
go run ./migrate/create/createMigration.go

<!-- run migrations using go
go run ./migrate/run/runMigration.go -->

run migration using CLI
migrate -path "migrate/migrations" -database "mysql://root:root@tcp(localhost:3306)/scim" up
