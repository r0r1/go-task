.PHONY: test

export PROJECT_PATH = /go/src/github.com/appleboy/gin-jwt

install:
	glide install

update:
	glide up

test:
	go test -v -cover -coverprofile=.cover/coverage.txt

html:
	go tool cover -html=.cover/coverage.txt

docker_test: clean
	docker run --rm \
		-v $(PWD):$(PROJECT_PATH) \
		-w=$(PROJECT_PATH) \
		appleboy/golang-testing \
		sh -c "make install && coverage all"

clean:
	rm -rf .cover
