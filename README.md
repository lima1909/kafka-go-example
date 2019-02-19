# documentation

- article: https://dev.to/davidsbond/golang-implementing-kafka-consumers-using-sarama-cluster-4fko
- kt: https://github.com/fgeller/kt
- kafkacat: https://github.com/edenhill/kafkacat
- docker ref: https://docs.confluent.io/current/installation/docker/docs/config-reference.html

# prepare

- (docker rm $(docker ps -aq)) 
- docker-compose up
- make create-topic (create topic "foo", check topic: kt topic --brokers 172.22.0.4:19092,172.22.0.5:29092)
- create a alias: kc='kafkacat -b 172.22.0.5:29092,172.22.0.4:19092'

# kafkacat

## producer

```
echo "my message" | kc -P -t foo

// or simple
kc -P -t foo

// with Key and delimiter ":"
kc -P -t foo -K:
```
