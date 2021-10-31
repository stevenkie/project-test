coverage.out: $(shell find . -type f -print | grep -v vendor | grep "\.go")
	@go test -cover -coverprofile ./coverage.out.tmp ./...
	@cat ./coverage.out.tmp | grep -v '.pb.go' | grep -v 'mock_' > ./coverage.out
	@rm ./coverage.out.tmp

test: coverage.out

cover: coverage.out
	@echo ""
	@go tool cover -func ./coverage.out

cover-html: coverage.out
	@go tool cover -html=./coverage.out

clean:
	@rm ./coverage.out

localenv-up:
	@docker-compose -f ./localenv/docker-compose.yml up -d

localenv-down:
	@docker-compose -f ./localenv/docker-compose.yml down --rmi all

build:
	@go build

run:
	@go run app.go