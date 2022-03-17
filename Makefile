.PHONY: generate
generate:
	wget https://raw.githubusercontent.com/dm03514/openapi-kafka-connect/main/kafka-connect-spec.yaml
	docker run -v $(shell pwd):/tmp/go-sdk -v $(shell pwd)/kafka-connect-spec.yaml:/tmp/kafka-connect-spec.yaml openapitools/openapi-generator-cli generate -i /tmp/kafka-connect-spec.yaml -g go -o /tmp/go-sdk