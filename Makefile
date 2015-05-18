VERSION = 0.0.1
REPO = aledbf

check-version:
ifndef VERSION
	@echo "Error: VERSION is undefined."
	@exit 1
endif

build-go:
	echo "Building..."
	gb build all

build: build-go mesos-template mesos-master mesos-slave
	echo "Building docker images..."

mesos-template: check-version
	sed "s/VERSION/$(VERSION)/g" mesos/template > mesos/Dockerfile
	docker build -t $(REPO)/$@:$(VERSION) mesos/.
	rm -f mesos/Dockerfile

mesos-master: mesos check-version
	sed "s/VERSION/$(VERSION)/g" mesos/master > mesos/Dockerfile
	docker build -t $(REPO)/$@:$(VERSION) mesos/.
	rm -f mesos/Dockerfile

mesos-slave: mesos check-version
	sed "s/VERSION/$(VERSION)/g" mesos/slave > mesos/Dockerfile
	docker build -t $(REPO)/$@:$(VERSION) mesos/.
	rm -f mesos/Dockerfile
