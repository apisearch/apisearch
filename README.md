# apisearch

Full text search service which communicates through JSON API.

### Docs

- [Apiary docs](http://docs.apisearch.apiary.io/)

### Dependencies

- Golang >= 1.6
- Docker

### Usage

- `docker-compose up -d`
- Show usage: `go run main.go`
- Create index: `go run main.go c`
- Run HTTP server: `go run main.go s`
- Import products: `go run main.go i`
- Sample data: `http://localhost:8081/heureka_cz.xml`
- Kibana: `http://localhost:5601`
- Cerebro: `http://localhost:9000`
