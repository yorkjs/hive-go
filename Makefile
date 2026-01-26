.PHONY: test test-v cover

# 运行所有测试
test:
	go test ./...

# 运行所有测试并显示详细日志
test-v:
	go test -v ./...

# 运行测试并生成覆盖率报告
cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
