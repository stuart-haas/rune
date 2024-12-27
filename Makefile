api-dev:
	go run main.go http start

web-dev:
	cd web && pnpm dev

dev: api-dev web-dev