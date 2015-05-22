REPO = aledbf
MESOS = 0.22.1-1.0.ubuntu1404
MESOS_VERSION = 0.22.1
ZOOKEEPER_VERSION = 3.5.0

repo_path = github.com/aledbf/coreos-mesos-zookeeper

GO = godep go
GOFMT = gofmt -l
GOLINT = golint
GOTEST = $(GO) test --cover --race -v
GOVET = $(GO) vet

GO_PACKAGES = pkg/boot pkg/confd pkg/etcd pkg/fleet pkg/log pkg/net
GO_PACKAGES_REPO_PATH = $(addprefix $(repo_path)/,$(GO_PACKAGES))

build: zookeeper-go build-tools mesos-template mesos-master mesos-slave zookeeper

mesos-go:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 godep go build -a -installsuffix cgo -ldflags '-s' -o mesos/bin/master-boot pkg/boot/mesos/master/main.go
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 godep go build -a -installsuffix cgo -ldflags '-s' -o mesos/bin/slave-boot pkg/boot/mesos/slave/main.go

zookeeper-go:
	echo "Building..."
	go-bindata -pkg bindata -o pkg/boot/zookeeper/bindata/bindata.go pkg/boot/zookeeper/bash/; \
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 godep go build -a -installsuffix cgo -ldflags '-s' -o zookeeper/bin/boot pkg/boot/zookeeper/main/boot.go

mesos-template:
	sed "s/VERSION/$(MESOS)/g" mesos/template > mesos/Dockerfile
	docker build -t $(REPO)/$@:$(MESOS_VERSION) mesos/.
	rm -f mesos/Dockerfile

mesos-master: mesos-go mesos-template
	sed "s/VERSION/$(MESOS)/g" mesos/master > mesos/Dockerfile
	docker build -t $(REPO)/$@:$(MESOS_VERSION) mesos/.
	rm -f mesos/Dockerfile

mesos-slave: mesos-go mesos-template
	sed "s/VERSION/$(MESOS)/g" mesos/slave > mesos/Dockerfile
	docker build -t $(REPO)/$@:$(MESOS_VERSION) mesos/.
	rm -f mesos/Dockerfile

zookeeper: zookeeper-go
	docker build -t $(REPO)/$@:$(ZOOKEEPER_VERSION) zookeeper/.

build-tools:
	echo "Building tools..."
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 godep go build -a -installsuffix cgo -ldflags '-s' -o tools/zkNodeUrls pkg/boot/zookeeper/main/zkNodesUrl.go

test-style:
	@$(GOFMT) $(GO_PACKAGES)
	@$(GOFMT) $(GO_PACKAGES) | read; if [ $$? == 0 ]; then echo "gofmt check failed."; exit 1; fi
	$(GOVET) $(GO_PACKAGES_REPO_PATH)
	@for i in $(addsuffix /...,$(GO_PACKAGES)); do \
		$(GOLINT) $$i; \
	done
