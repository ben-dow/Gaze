
current_dir := $(shell pwd)
version := $(shell cat version.txt)-$(shell date +%s%N | cut -b1-13)
branch := $(shell git rev-parse --abbrev-ref HEAD)

build:
	echo $(version)
	docker build -t gaze:$(branch)-$(version) --progress=plain -f build/package/Dockerfile .

run: build
	docker run -p 3000:3000 gaze:$(branch)-$(version)