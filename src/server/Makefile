build:
	@if [ ! -d bin ]; then echo "creating bin folder"; mkdir bin ; fi;
	@if [ -d bin/server ]; then rm bin/server; fi;
	@echo "building server"; 
	@cd cmd/server; \
	go build -o ../../bin/server; 
	@echo "building client"; 
	@cd cmd/client; \
	go build -o ../../bin/client; 

build_run:
	make build && bin/server

check_install:
	which swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger
	
swagger:
	GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models
