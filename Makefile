CASSANDRA_HOST:=127.0.0.1
KEYSPACE:=twitter_example
TARGET_PACKAGE:=twitter
GENERATED_SRC_DIR:=src/$(TARGET_PACKAGE)
GENERATED_GO_FILE:=generated_schema.go

default: $(GENERATED_GO_FILE)

$(GENERATED_GO_FILE): $(GENERATED_SRC_DIR)
	cqlc -i $(CASSANDRA_HOST) -k $(KEYSPACE) -p $(TARGET_PACKAGE) -o $(GENERATED_SRC_DIR)/$@

$(GENERATED_SRC_DIR):
	mkdir -p $@

schema:
	CQLSH_HOST=$(CASSANDRA_HOST) cqlsh -f schema.cql

clean:
	rm -r src
