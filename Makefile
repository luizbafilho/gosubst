IMAGE=gosubst-image

build_image:
	docker build -t ${IMAGE} .

build:
	docker run -it -v $(shell pwd):/go/src/github.com/luizbafilho/gosubst -w /go/src/github.com/luizbafilho/gosubst ${IMAGE} go build -o _bin/gosubst

cross_build:
	docker run -it -v $(shell pwd):/go/src/github.com/luizbafilho/gosubst -w /go/src/github.com/luizbafilho/gosubst -e CGO_ENABLED=0 ${IMAGE} gox -output="_dist/gosubst_{{.OS}}_{{.Arch}}"

test:
	docker run -it -v $(shell pwd):/go/src/github.com/luizbafilho/gosubst -w /go/src/github.com/luizbafilho/gosubst/gosubst ${IMAGE} go test

run:
	docker run -it -v $(shell pwd):/go/src/github.com/luizbafilho/gosubst -w /go/src/github.com/luizbafilho/gosubst ${IMAGE} bash
