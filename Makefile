.PHONY: create-kafka

create-topic:
	docker run --net=kafka-go-example_svc_net --rm confluentinc/cp-kafka:5.1.1 kafka-topics --create --topic foo \
				--partitions 2 \
				--replication-factor 2 \
				--if-not-exists \
				--zookeeper 172.22.0.2:22181,172.22.0.3:32181
