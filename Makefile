compile-go:
	@echo "Compiling asbemulator go files"
	@CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o asbemulator ./cmd/asbemulator.go
	@echo "Finished compiling the static Go app"

compile-python:
	@echo "Compiling amqp_python gateway.py"
	@pyinstaller gateway.spec
	@echo "Finished compiling the static Python app"
