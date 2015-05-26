REPO = aledbf
MESOS = 0.22.1-1.0.ubuntu1404
MESOS_VERSION = 0.22.1
ZOOKEEPER_VERSION = 3.5.0

MARATHON_VERSION = 0.8.2-RC3

repo_path = github.com/aledbf/coreos-mesos-zookeeper

GO = godep go
GOFMT = gofmt -l
GOLINT = golint
GOTEST = $(GO) test --cover --race -v
GOVET = $(GO) vet

GO_PACKAGES = pkg/boot pkg/confd pkg/etcd pkg/fleet pkg/log pkg/net
GO_PACKAGES_REPO_PATH = $(addprefix $(repo_path)/,$(GO_PACKAGES))

build: zookeeper-go build-tools mesos-template mesos-master mesos-slave mesos-marathon zookeeper

mesos-go:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 godep go build -a -installsuffix cgo -ldflags '-s' -o mesos/bin/master-boot pkg/boot/mesos/master/main.go
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 godep go build -a -installsuffix cgo -ldflags '-s' -o mesos/bin/slave-boot pkg/boot/mesos/slave/main.go
	go-bindata -pkg bindata -o pkg/boot/mesos/marathon/bindata/bindata.go pkg/boot/mesos/marathon/bash/; \
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 godep go build -a -installsuffix cgo -ldflags '-s' -o mesos/bin/marathon-boot pkg/boot/mesos/marathon/main.go

mesos-template:
	sed "s/#VERSION#/$(MESOS)/g" mesos/template > mesos/Dockerfile
	docker build -t $(REPO)/$@:$(MESOS_VERSION) mesos/.
	rm -f mesos/Dockerfile

mesos-master: mesos-go mesos-template
	sed "s/#VERSION#/$(MESOS_VERSION)/g" mesos/master > mesos/Dockerfile
	docker build -t $(REPO)/$@:$(MESOS_VERSION) mesos/.
	rm -f mesos/Dockerfile

mesos-slave: mesos-go mesos-template
	sed "s/#VERSION#/$(MESOS_VERSION)/g" mesos/slave > mesos/Dockerfile
	docker build -t $(REPO)/$@:$(MESOS_VERSION) mesos/.
	rm -f mesos/Dockerfile

build-mesos-marathon: mesos-template
	sed "s/#MARATHON_VERSION#/$(MARATHON_VERSION)/;s/#VERSION#/$(MESOS_VERSION)/" mesos/build-marathon > mesos/Dockerfile
	docker build -t $(REPO)/$@:$(MESOS_VERSION) mesos/.
	rm -f mesos/Dockerfile

mesos-marathon: build-mesos-marathon mesos-go
	cp mesos/marathon mesos/Dockerfile
	docker cp `docker create $(REPO)/build-mesos-marathon:$(MESOS_VERSION) /bin/bash`:/marathon/target/marathon-assembly-$(MARATHON_VERSION).jar mesos/
	mv mesos/marathon-assembly-$(MARATHON_VERSION).jar mesos/marathon-assembly.jar
	sed "s/#MARATHON_VERSION#/$(MARATHON_VERSION)/;s/#VERSION#/$(MESOS_VERSION)/" mesos/marathon > mesos/Dockerfile
	docker build -t $(REPO)/$@:$(MESOS_VERSION) mesos/.
	rm -f mesos/Dockerfile

zookeeper: zookeeper-go
	docker build -t $(REPO)/$@:$(ZOOKEEPER_VERSION) zookeeper/.

zookeeper-go:
	echo "Building..."
	go-bindata -pkg bindata -o pkg/boot/zookeeper/bindata/bindata.go pkg/boot/zookeeper/bash/; \
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 godep go build -a -installsuffix cgo -ldflags '-s' -o zookeeper/bin/boot pkg/boot/zookeeper/main/boot.go

build-tools:
	echo "Building tools..."

test: test-style
	go test -cover -timeout 10s ./...

test-style:
	@$(GOFMT) $(GO_PACKAGES)
	@$(GOFMT) $(GO_PACKAGES) | read; if [ $$? == 0 ]; then echo "gofmt check failed."; exit 1; fi
	$(GOVET) $(GO_PACKAGES_REPO_PATH)
	@for i in $(addsuffix /...,$(GO_PACKAGES)); do \
		$(GOLINT) $$i; \
	done
