build_go_binary:
	go build -o musizticle

build_docker:
	#docker buildx build --platform linux/arm/v7 -t musizticle-builer --load .
	docker build -t musizticle-builder .
	docker create --name extract musizticle-builder
	docker cp extract:/src/musizticle/musizticle ./
	docker rm extract
	scp musizticle musizticle.local:/home/admiral/musizticle/app
	#rm musizticle

