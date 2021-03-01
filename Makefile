dev: devserver devclient
	env EDITHD_WORKING_DIR=./build/working_dir \
	EDITHD_SERVER_ROOT=./build/server_root \
	./build/edithd

devserver:
	go build --race -o ./build/edithd ./cmd/server/

devclient:
	go build --race -o ./build/edith ./cmd/client/

prodclient:
	go build -o /home/nana/go/bin/edith ./cmd/client/

prod:
	go build --race -o /home/nana/Desktop/edithd/edithd ./cmd/server/
	go build -o /home/nana/go/bin/edith ./cmd/client/
	sudo systemctl restart edithd

release:
	env GOOS=windows GOARCH=amd64 go build -ldflags="-X main.serverAddr=edith.local:54920" -o ./release/clients/windows/edith.exe ./cmd/client
	env GOOS=darwin GOARCH=amd64 go build -ldflags="-X main.serverAddr=edith.local:54920" -o ./release/clients/mac/edith ./cmd/client
	env GOOS=linux GOARCH=amd64 go build -ldflags="-X main.serverAddr=edith.local:54920" -o ./release/clients/linux/edith ./cmd/client
	env GOOS=linux GOARCH=amd64 go build -o ./release/server/edithd ./cmd/server
	go build -ldflags="-X main.serverAddr=edith.local:54920" -o /home/nana/go/bin/edith ./cmd/client/
	zip -rm /home/nana/edith_release.zip ./release/*