export tag=prod-20240706-1
export GOPROXY=https://goproxy.cn,direct
build:
	echo "building demo binary"
	go build -o demo .

build_image:
	echo "building demo container"
	docker build -t your-registry:${tag} .