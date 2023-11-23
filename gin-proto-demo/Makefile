protoc:
	 protoc -I proto/ --go_out . --go-gin_out . proto/hello.proto

protoc-go-inject-tag:
	protoc-go-inject-tag -input proto/*.pb.go

.PHONY: build_autotest
build_autotest:
	@echo "build_autotest"

.PHONY: autotest_run
autotest_run:
	@echo "autotest_run"

.PHONY: autotest
autotest: build_autotest autotest_run