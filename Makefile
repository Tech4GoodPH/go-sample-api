GOFLAGS='-mod=vendor'
PORT=8080
BUILD_IMAGE=tech4goodph/bayanihan-basurahan-api:latest

go-build:
	@echo '======================================================='
	@echo ' Compiling Go packages along with their dependencies'
	@echo '======================================================='
	go mod tidy
	go mod vendor
	go build -o ./build -v .
	@echo 'Done.'

go-run: go-build
	@echo '======================================================='
	@echo ' Running application'
	@echo '======================================================='
	@echo 'Running... (CTRL+C to exit)'
	@./build/go-sample-api.git
	

go-test:
	@echo '======================================================='
	@echo ' Running tests: unit test'
	@echo '======================================================='
	go test -v ./test

go-lint: go-build
	@echo '======================================================='
	@echo ' Running tests: lint test'
	@echo '======================================================='
	./bin/golangci-lint-v1.17.1 run -c config/.golangci.yml

build-app: 
	@echo '======================================================='
	@echo ' Packaging application image'
	@echo '======================================================='
	docker build -t ${BUILD_IMAGE} .
	@echo 'Done.'

publish-app:
	@echo '========================================'
	@echo ' Publishing image'
	@echo '========================================'
	@docker push ${BUILD_IMAGE}
	@echo 'Done.'