PROJECT_DIR=github.com/hieuphq/califit
PROJECT_ROOT=${GOPATH}/src/${PROJECT_DIR}
PROJECT_NAME=califit
IMAGE_NAME=hieupq/califit
DEFAULT_PORT=8080

.PHONY: build run dev docker-build

build:
	CGO_ENABLED=0 go build -a -installsuffix cgo -o ${PROJECT_ROOT}/bin/server ${PROJECT_ROOT}/cmd/server

dev:
	CGO_ENABLED=0 go build -a -installsuffix cgo -o ${PROJECT_ROOT}/bin/server ${PROJECT_ROOT}/cmd/server
	${PROJECT_ROOT}/bin/server

run:
	${PROJECT_ROOT}/bin/server

migrate-up:
	@goose -dir="./migration" postgres "postgres://postgres:postgres@localhost:5432/califit?sslmode=disable" up

migrate-down:
	@goose -dir="./migration" postgres "postgres://postgres:postgres@localhost:5432/califit?sslmode=disable" down

migrate-create:
	@goose -dir="./migration" create ${file} sql

gen:
	@sqlboiler --wipe psql

local-env:
	docker-compose -f test_postgres/docker-compose.yml down
	docker-compose -f test_postgres/docker-compose.yml up -d

test:
	go test -cover -v ${PROJECT_ROOT}/src/service

test-output:
	go test -cover -v ${PROJECT_ROOT}/src/service -coverprofile=coverage.out && go tool cover -html=coverage.out -o coverage.html

docker-build:
	docker build \
	--build-arg PROJECT_DIR="${PROJECT_DIR}" \
	--build-arg DEFAULT_PORT="8080" \
	-t ${IMAGE_NAME} .

docker-run:
	docker run \
	-d --name some-${PROJECT_NAME} -p 8080:${DEFAULT_PORT}  ${IMAGE_NAME} \

Repositories= User
Usecases= User

mocks:
	$(foreach module,$(Repositories), \
	 	mockery -dir=${PROJECT_ROOT}/src/interfaces/repository -output=${PROJECT_ROOT}/src/interfaces/repository/mocks -name=${module}Repository;)
	$(foreach module,$(Usecases), \
	 	mockery -dir=${PROJECT_ROOT}/src/usecase -output=${PROJECT_ROOT}/src/usecase/mocks -name=${module}Usecase;)