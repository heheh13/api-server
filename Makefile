#hi give it a try

#defining registry , image_name and tag
#values of each  vars can be overritten from cli
# make push registry="heheh" ....

registry ?= heheh13
image_name ?= apiserver
tag ?= latest

## build the docker image and push to registry
## since dokcer image contains all information to build the binary
## its seems better to me to buid the binary os independently
push:
#need to run the docker file
	docker build -t ${registry}/${image_name}:${tag} .;\
	docker push ${registry}/${image_name}:${tag}


## helm install and unistall using make...
install:
	helm install ${image_name} ./helm-charts 
uninstall:
	helm uninstall ${image_name}

