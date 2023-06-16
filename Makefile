BINARY_NAME=go-demo
OUTPUT=build

.PHONY: build

build: clean
	mkdir ${OUTPUT} -p
	cp resources/* ${OUTPUT} -r
	# GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin main.go
	# GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux main.go
	CGO_ENABLED=1 GOARCH=amd64 GOOS=windows go build -o ${OUTPUT}/${BINARY_NAME}-win.exe
	#go build -o ${BINARY_NAME}-windows main.go

run: build
	${OUTPUT}/${BINARY_NAME}-win.exe

clean:
	rm ${OUTPUT}/ -rf
	go clean


test:
	echo 123
# go test ./

test_coverage:
	go test ./ -coverprofile=coverage.out

dep:
	go mod download

vet:
	go vet

lint:
	golangci-lint run --enable-all