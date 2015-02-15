run:
	@go run server/main.go

createuser:
	@go run cli/user.go

test:
	@for d in */ ; do \
		go test -coverprofile=coverage.$${d:0:(-1)}.out ./$${d} ; \
	done
	@echo "mode: set" > coverage.out && cat coverage.*.out | sed '/mode: set/d' >> coverage.out

cov:
	@go tool cover -html coverage.out
