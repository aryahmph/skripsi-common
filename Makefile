GO_MODULE := github.com/aryahmph/skripsi-common

.PHONY: protoc
protoc:
	protoc --go_opt=module=${GO_MODULE} --go_out=. \
	--go-grpc_opt=module=${GO_MODULE} --go-grpc_out=. \
	./proto/ticket/*.proto \
	./proto/order/*.proto \
	./proto/payment/*.proto \