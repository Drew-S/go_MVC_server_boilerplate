serve: clean build
	go run *.go

sass:
	./.sass.sh

clean:
	rm -rf wwwroot/build/css

build: sass

watch:
	while true; do find wwwroot/src | entr -d make build; done

deps:
	go get github.com/go-gorp/gorp
	go get github.com/gorilla/mux
	go get github.com/gorilla/sessions