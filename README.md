# kafka-example
The repository have 5 part :

1.  Async producer
2. Sync producer
3. Consumer
4. Mocking testing
5. Http test aka Handler testing

------------
## How to use
### 1. Create kafka cluster by docker:
 To run kafka cluster by docker, you can go to [landoop/kafka-lenses-dockerhub](https://hub.docker.com/r/landoop/kafka-lenses-dev/ "landoop/kafka-lenses")
- Lenses kafka- Web UI for kafka:  `localhost:3030`
- Kafka broker : `localhost:9092`
- Create a topic: `demo_topic`
- Download necessary library:

    	govendor sync

### 2. Run async producer:
    cd producer
    go run main.go
- To stop the producer
    press: `ctrl + C`
### 3. Run sync producer:
    cd producersync
    go run main.go
- To stop the producer
    press: `ctrl + C`
### 4. Run sync consumer:
    cd consumer
    go run main.go
- To stop the consumer
    press: `ctrl + C`