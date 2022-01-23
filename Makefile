GO_FILES := $(shell find . -name '*.go')

all: dist/rackmon.linux.arm

dist:
	mkdir -p dist

release: dist/rackmon.linux.arm.release

dist/rackmon.linux.arm: .pretty $(GO_FILES) dist
	GOARCH=arm GOOS=linux go build -o $@

dist/rackmon.linux.arm.release: .pretty $(GO_FILES) dist
	GOARCH=arm GOOS=linux go build -ldflags="-s -w" -o $@
	upx --brute $@


.pretty: $(GO_FILES)
	find . -name "*.go" -print0 | xargs -0 goimports -w
	touch .pretty

.PHONY: clean
clean:
	rm -fr updatemgr dist .pretty
