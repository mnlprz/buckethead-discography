postgres:
	sudo docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:alpine

createdb:
	sudo docker exec -it postgres createdb --username=root --owner=root buckethead

dropdb:
	docker exec -it postgres dropdb buckethead

.PHONY: postgres createdb dropdb