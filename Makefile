build:
	@if [ ! -d src/server/bin ]; then echo "creating bin folder"; mkdir src/server/bin ; fi;
	@if [ -d src/server/bin/server ]; then rm src/server/bin/server; fi;
	@echo "building server"; 
	@cd src/server/cmd/server; \
	go build -o ../../bin/server; 
	@echo "building client"; 
	@cd src/server/cmd/client; \
	go build -o ../../bin/client; 

build_run:
	make build && src/server/bin/server
