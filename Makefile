VERSION=v1.0.6

## clean: runs go clean
clean:
	@echo "Cleaning..."
	@go clean
	@echo "Cleaned!"

## test: runs all tests
test:
	go test -cover -v ./...

## help: displays help
help: Makefile
	@echo " Choose a command:"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

## update-pkg-cache: forces update on go.pkg.dev
update-pkg-cache:
	GOPROXY=https://proxy.golang.org GO111MODULE=on \
	go get github.com/tsawler/signer@${VERSION}