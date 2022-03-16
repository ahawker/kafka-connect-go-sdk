.PHONY
generate:
  docker run -v $(pwd):/tmp/go-sdk -v $(pwd)/kafka-connect-spec.yaml:/tmp/kafka-connect-spec.yaml openapitools/openapi-generator-cli generate -i /tmp/kafka-connect-spec.yaml -g go -o /tmp/go-sdk
