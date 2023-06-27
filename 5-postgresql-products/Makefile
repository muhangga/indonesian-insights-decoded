run:
	go run main.go

migrate-up:
	migrate -path db/migrations -database "postgres://postgres:root@localhost:5432/decoded-insight-test?sslmode=disable" up

migrate-down:
	migrate -path db/migrations -database "postgres://postgres:root@localhost:5432/decoded-insight-test?sslmode=disable" down