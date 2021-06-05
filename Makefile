proto:
	protoc --go_out=. --go_opt=paths=source_relative \
			--go-amqp-rpc_out=. \
            *.proto
