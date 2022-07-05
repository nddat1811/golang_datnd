all:
	docker compose up -d
	
build:
	docker build -t my-golang -f Dockerfile .
server:
	go run main.go
migrateup:
	migrate -path migration -database "mysql://root:nguyenducdat18112001@tcp(localhost:3308)/golang_api?x-no-lock=true" -verbose up

migratedown:
	migrate -path migration -database "mysql://dat:123456@tcp(localhost:3306)/golang_api?x-no-lock=true" -verbose down

.PHONY: all build migrateup migratedown
clean:
	docker compose down
#	clean docker compose 

	docker rmi $$(docker images -a -q) -f
#	remove all images #-f auto force
#	docker rm $$(docker ps -a -f status=exited -q)
#	remove all container

