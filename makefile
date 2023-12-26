gen:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    rpc/search/search.proto

build:
	go build -o bin/server cmd/server/main.go

run: build
	./bin/server

grpcurlSearch:
	grpcurl -plaintext -d '{"targetNumber": 42}' localhost:8080 SearchService/Search

grpcurlInsert:
	grpcurl -plaintext -d '{"number": 42}' localhost:8080 SearchService/Insert

grpcurlDelete:
	grpcurl -plaintext -d '{"number": 42}' localhost:8080 SearchService/Delete

loadEnv:
	export $(xargs < .env)

checkIFEnvExists:
    ifeq (,$(wildcard .env))
        $(error .env file does not exist)
    endif

dockerBuildRun: checkIFEnvExists
	docker build -t searcher . && \
	docker run --rm -it -p 8082:8082 --env-file .env searcher

dockerPush: checkIFEnvExists loadEnv
	docker build -t searcher . && \
	docker tag searcher $(DOCKER_REGISTRY) && \
	docker push $(DOCKER_REGISTRY):latest