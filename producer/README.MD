# Kafka

* Create topic
```sh
./kafka-topics.sh --create --topic sales-ingestion --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1
```

* Listen from topic
```sh
./kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic sales-ingestion
```
