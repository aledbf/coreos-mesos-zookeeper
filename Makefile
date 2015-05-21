REPO = aledbf
MESOS = 0.22.1-1.0.ubuntu1404
MESOS_VERSION = 0.22.1
ZOOKEEPER_VERSION = 3.5.0

build: zookeeper-go build-tools mesos-template mesos-master mesos-slave zookeeper

zookeeper-go:
	echo "Building..."
	go-bindata -pkg bindata -o pkg/boot/zookeeper/bindata/bindata.go pkg/boot/zookeeper/bash/; \
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 godep go build -a -installsuffix cgo -ldflags '-s' -o zookeeper/bin/boot pkg/boot/zookeeper/main/boot.go

mesos-template:
	sed "s/VERSION/$(MESOS)/g" mesos/template > mesos/Dockerfile
	docker build -t $(REPO)/$@:$(MESOS_VERSION) mesos/.
	rm -f mesos/Dockerfile

mesos-master: mesos-template
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 godep go build -a -installsuffix cgo -ldflags '-s' -o mesos/bin/boot pkg/boot/mesos/master/main.go
	sed "s/VERSION/$(MESOS)/g" mesos/master > mesos/Dockerfile
	docker build -t $(REPO)/$@:$(MESOS_VERSION) mesos/.
	rm -f mesos/Dockerfile

mesos-slave: mesos-template
	sed "s/VERSION/$(MESOS)/g" mesos/slave > mesos/Dockerfile
	docker build -t $(REPO)/$@:$(MESOS_VERSION) mesos/.
	rm -f mesos/Dockerfile

zookeeper: zookeeper-go
	docker build -t $(REPO)/$@:$(ZOOKEEPER_VERSION) zookeeper/.

build-tools:
	echo "Building tools..."
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 godep go build -a -installsuffix cgo -ldflags '-s' -o tools/zkNodeUrls pkg/boot/zookeeper/main/zkNodesUrl.go