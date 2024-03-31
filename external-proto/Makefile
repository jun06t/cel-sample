.PHONY: proto proto_format

proto: proto_format
	protoc \
	-I=./proto \
	--go_out=./proto \
	--go_opt=module=github.com/jun06t/cel-sample/external-proto/proto \
	./proto/*.proto

proto_format:
	find ./proto/ -name "*.proto" | xargs clang-format -i
