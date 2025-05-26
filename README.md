# Quotes API Service

## Run Instructions

### Using Docker
```bash
docker pull nikita213hub/quotes-api:latest
docker run -p 8080:8080 -e ADDRESS=0.0.0.0:8080 nikita213hub/quotes-api:latest
```

### From Source
```bash
git clone https://github.com/Nikita213-hub/testTaskGo.git
cd testTaskGo
go build -o quotes-api ./cmd/main.go
./quotes-api
```

### Running Tests
```bash
go test ./...
```