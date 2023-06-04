BINARY_NAME = homepage
INSTALLATION_PATH = /usr/local/bin

build:
	./tailwindcss -i ./web/css/input.css -o ./web/css/output.css --minify
	go build -o bin/$(BINARY_NAME) ./web

run: gorun watch

clean:
	go clean
	rm -f bin/$(BINARY_NAME)
	rm -f web/css/output.css

watch:
	./tailwindcss -i ./web/css/input.css -o ./web/css/output.css --watch

gorun:
	go run ./web

install: clean build
	cp bin/$(BINARY_NAME) $(INSTALLATION_PATH)/$(BINARY_NAME)

service:
	cp homepage.service /etc/systemd/system/homepage.service
	systemctl daemon-reload
	systemctl enable homepage.service
	systemctl start homepage.service