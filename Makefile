build:
	docker-compose build todo-app

run:
	docker-compose up todo-app

test:
	go test -v ./...

migrate:
	migrate -path ./schema -database 'postgres://postgres:mypassword@0.0.0.0:5436/todo-app?sslmode=disable' up

swag:
	swag init -g cmd/main.go

