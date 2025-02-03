MAIN_FILE=cmd/main.go

BINARY_OUT=server
OUT_DIR=out

all: build

build:
	@echo "Building ${MAIN_FILE}..."
	@go build -o ${OUT_DIR}/${BINARY_OUT} ${MAIN_FILE}
	@echo "Builded in directory '${OUT_DIR}'"

clean:
	@echo "Cleaning..."
	@rm -rf out
	@echo "Cleaned"