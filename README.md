# API Search importer

Microservice used for importing products into database. Part of apisearch project.

### Docs

- [apiary](docs.apisearchimporter.apiary.io)

### Dependencies

- Golang >= 1.6
- Elasticsearch >= 5.0.0

### Usage

- `docker-compose up -d`
- Create index: `go run main.go createIndex`
- Run HTTP server: `go run main.go server`
- Import products: `go run main.go import`
- Sample data: `http://localhost:8081/heureka_cz.xml`
- Kibana: `http://localhost:5601`
- Cerebro: `http://localhost:9000`
