build:
	@echo "===> Building"
	@go build

coverage: test
	@echo "===> Creating coverage report"
	go tool cover -html=coverage.txt

lint:
	@echo "===> Linting with vet"
	@go vet

test: lint build
	@echo "===> Testing"
	@scripts/test.sh
