
build:
	go build -o ./pkg/cmd/pubby-mcsubface/bin/pubby-mcsubface github.com/elliotcourant/pubby-mcsubface/pkg/cmd/pubby-mcsubface

docker: build
	docker build -t pubby-mcsubface ./pkg/cmd/pubby-mcsubface
	docker tag pubby-mcsubface docker.pkg.github.com/elliotcourant/pubby-mcsubface/pubby-mcsubface:latest
