# Set default environment variables
export PORT=8080
export MONGODB=mongodb://localhost:27017/
export EMAIL_CHECK_SERVICE=https://evservice.azurewebsites.net/Email
export AUTHORIZATION_ENABLED=false

build:
	go build -o app

run:
	make build && ./app
