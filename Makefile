compose-build:
	docker-compose build

compose-up:
	docker-compose up

open:
	docker exec -it real-time-forum_db_1 psql -U postgres
