#hi give it a try

#defining registry , image_name and tag
#values of each  vars can be overritten from cli
# make push registry="heheh" ....

registry ?= heheh13
image_name ?= apiserver
tag ?= latest

## build the docker image 

## since dokcer image contains all information to build the binary
## its seems better to me to buid the binary os independently
build:
	docker build -t ${registry}/${image_name}:${tag} .
	touch build

## make push needs to build the docker image before push?
push: build

	docker push ${registry}/${image_name}:${tag}

## helm install and unistall using make...
install:
	helm install ${image_name} ./helm-charts 
uninstall:
	helm uninstall ${image_name}

clean:
	rm -f build
