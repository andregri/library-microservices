go-mod:
	GO111MODULE=on go mod tidy -v
	GO111MODULE=on go mod vendor -v

test:
	docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit
	docker-compose -f docker-compose.test.yml down --volumes