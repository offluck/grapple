CURRDIR=$(shell pwd)
BINDIR=${CURRDIR}/bin
PACKAGE=github.com/offluck/grapple/cmd/app

build:
	mkdir -p ${BINDIR}/app
	go build -o ${BINDIR}/app ${PACKAGE}

run:
	go run ${PACKAGE}
