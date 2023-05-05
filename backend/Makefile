dev_db = "postgres://postgres:w0ntt3lly0u@localhost/tweemote?sslmode=disable"
test_db = "postgres://postgres:w0ntt3lly0u@localhost/tweemote_test?sslmode=disable"

.PHONY: all run install_migrate migrate rollback testdb local

all: local

local:
	govendor add +external
	govendor sync
	govendor install

prod:
	docker-compose build

server: local
	tweemote server

#
# DB
#

migrate:
	docker-compose build server
	docker-compose run --rm server tweemote migrate

rollback: install_migrate
	migrate -path="./migrations" -verbose -database="$(dev_db)" down 1

dropdb:
	docker-compose stop postgres
	docker-compose rm -f postgres
	docker volume rm -f tweemote_postgres
	docker-compose up -d postgres

seed:
	docker-compose exec postgres psql -c "INSERT INTO twitter_applications (name, consumer_key, consumer_secret) VALUES('qainstructor', 'redacted', 'redacted')" -U postgres tweemote

wait:
	echo "Waiting 15 seconds for PG to become ready"
	sleep 15

recreatedb: dropdb wait migrate

#
# Testing
#

testdb:
	docker-compose exec postgres psql -c "DROP DATABASE IF EXISTS tweemote_test" -U postgres tweemote
	docker-compose exec postgres psql -c "CREATE DATABASE tweemote_test WITH TEMPLATE tweemote" -U postgres tweemote
	migrate -path="./migrations" -verbose -database="$(test_db)" up
