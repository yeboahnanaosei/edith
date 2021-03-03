dev: devserver devclient
	env EDITHD_WORKING_DIR=./build/working_dir \
	EDITHD_SERVER_ROOT=./build/server_root \
	./build/edithd

devserver:
	go build --race -o ./build/edithd ./cmd/edithd

devclient:
	go build --race -o ./build/edith ./cmd/edith/

prodclient:
	go build -o ${HOME}/go/bin/edith ./cmd/edith/

prod:
	go build --race -o ${HOME}/Desktop/edithd/edithd ./cmd/edithd
	go build -o ${HOME}/go/bin/edith ./cmd/edith/
	sudo systemctl restart edithd

release:
	env GOOS=windows GOARCH=amd64 go build -ldflags="-X main.serverAddr=edith.local:54920" -o ./release/clients/windows/edith.exe ./cmd/edith
	env GOOS=darwin GOARCH=amd64 go build -ldflags="-X main.serverAddr=edith.local:54920" -o ./release/clients/mac/edith ./cmd/edith
	env GOOS=linux GOARCH=amd64 go build -ldflags="-X main.serverAddr=edith.local:54920" -o ./release/clients/linux/edith ./cmd/edith
	env GOOS=linux GOARCH=amd64 go build -o ./release/server/edithd ./cmd/edithd
	go build -ldflags="-X main.serverAddr=edith.local:54920" -o ${HOME}/go/bin/edith ./cmd/edith/
	zip -rm ${HOME}/edith_release.zip ./release/*
	rm -r release
