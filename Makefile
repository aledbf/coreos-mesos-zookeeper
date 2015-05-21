REPO = aledbf
MESOS = 0.22.1-1.0.ubuntu1404
MESOS_VERSION = 0.22.1
ZOOKEEPER_VERSION = 3.5.0

build-go:
	echo "Building..."
	go-bindata -pkg bindata -o pkg/boot/zookeeper/bindata/bindata.go pkg/boot/zookeeper/bash/; \
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 godep go build -a -installsuffix cgo -ldflags '-s' -o zookeeper/bin/boot pkg/boot/zookeeper/main/boot.go

build: build-go mesos-template mesos-master mesos-slave zookeeper
	echo "Building docker images..."

mesos-template:
	sed "s/VERSION/$(MESOS)/g" mesos/template > mesos/Dockerfile
	docker build -t $(REPO)/$@:$(MESOS_VERSION) mesos/.
	rm -f mesos/Dockerfile

mesos-master: mesos
	sed "s/VERSION/$(MESOS)/g" mesos/master > mesos/Dockerfile
	docker build -t $(REPO)/$@:$(MESOS_VERSION) mesos/.
	rm -f mesos/Dockerfile

mesos-slave: mesos
	sed "s/VERSION/$(MESOS)/g" mesos/slave > mesos/Dockerfile
	docker build -t $(REPO)/$@:$(MESOS_VERSION) mesos/.
	rm -f mesos/Dockerfile

zookeeper: build-go
	docker build -t $(REPO)/$@:$(ZOOKEEPER_VERSION) zookeeper/.

build-tools:
	echo "Building tools..."
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 godep go build -a -installsuffix cgo -ldflags '-s' -o tools/zkNodeUrls pkg/boot/zookeeper/zkNodesUrl.go