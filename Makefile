run:
	@go run server/main.go

test:
	@go test -coverprofile=coverage.server.out ./server
	@go test -coverprofile=coverage.alert.out ./alert
	@go test -coverprofile=coverage.auth.out ./auth
	@go test -coverprofile=coverage.dashboard.out ./dashboard
	@go test -coverprofile=coverage.monitor.out ./monitor
	@go test -coverprofile=coverage.session.out ./session
	@go test -coverprofile=coverage.template.out ./template
	@echo "mode: set" > coverage.out && cat coverage.*.out | sed '/mode: set/d' >> coverage.out

cov:
	@go tool cover -html coverage.out
