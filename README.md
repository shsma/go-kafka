# Go Kafka App

This Go app demonstrates Kafka messaging usiing k8s(minikube) + skaffold


## Prerequisites

- Go (Golang) installed: https://golang.org/dl/
- Accessible Kafka cluster or instance.
- Docker and Docker Compose (optional)
- Skaffold (optional)

## Run Kafka Locally
### Run with Docker Compose
```shell
docker-compose up 
```

### Run with Skaffold
```shell
skaffold run  --port-forward
```
## Run App
1. Install dependencies: `go mod download`
2. Update Kafka connection info for `producers` and `consumers`.
3. Run: `go run main.go`

## Customization

- Modify message processing in Kafka consumer.
- Adjust Kafka topic and format.
- Handle errors, add features.

## Resources

- [Apache Kafka Docs](https://kafka.apache.org/documentation/)
- [Confluent Kafka Go Client](https://github.com/confluentinc/confluent-kafka-go)

## License

MIT License. See [LICENSE](LICENSE).

Fork, enhance, and adapt. Questions/improvements? Create an issue or pull request. **Happy coding!**

