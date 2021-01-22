run: test

mod:
	# This make rule requires Go 1.11+
	GO111MODULE=on go mod tidy
	GO111MODULE=on go mod vendor

test:
	docker-compose -f docker-compose.yml up --abort-on-container-exit
	docker-compose -f docker-compose.yml down --volumes
