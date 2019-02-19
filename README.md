# Documentation

- article: https://dev.to/davidsbond/golang-implementing-kafka-consumers-using-sarama-cluster-4fko
- kt: https://github.com/fgeller/kt

# prepare

- (docker rm $(docker ps -aq)) 
- docker-compose up
- make create-topic (create topic "foo", check topic: kt topic --brokers 172.22.0.4:19092,172.22.0.5:29092)

# kafkacat

create a alias: kc='kafkacat -b 172.22.0.5:29092,172.22.0.4:19092'

## producer

```
echo "my message" | kc -P -t foo

// or simple
kc -P -t foo

// with Key and delimiter ":"
kc -P -t foo -K:
```