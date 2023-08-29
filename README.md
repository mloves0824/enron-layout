# Enron Project Template

## Install Enron
```
go install github.com/mloves0824/enron/cmd/enron@latest
```
## Create a service
```
# Create a template project
enron new server

cd server
# Add a proto template
enron proto add api/server/server.proto
# Generate the proto code
enron proto client api/server/server.proto
# Generate the source code of service by proto file
enron proto server api/server/server.proto -t internal/service

go generate ./...
go build -o ./bin/ ./...
./bin/server -conf ./configs
```
## Generate other auxiliary files by Makefile
```
# Download and update dependencies
make init
# Generate API files (include: pb.go, http, grpc, validate, swagger) by proto file
make api
# Generate all files
make all
```
## Automated Initialization (wire)
```
# install wire
go get github.com/google/wire/cmd/wire

# generate wire
cd cmd/server
wire
```

## Docker
```bash
# build
docker build -t <your-docker-image-name> .

# run
docker run --rm -p 8000:8000 -p 9000:9000 -v </path/to/your/configs>:/data/conf <your-docker-image-name>
```

