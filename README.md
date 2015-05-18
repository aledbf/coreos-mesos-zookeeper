# coreos-mesos-zookeeper

Services and images to start mesos and zookeeper in coreos

- [X] zookeeper
  - [X] use alpine linux
  - [X] service unit
  - [X] ~~command to generate configuration file~~ (included in boot)
  - [ ] remove node from the cluster on restart

- [ ] mesos
  - [ ] master
  - [ ] service unit
  - [ ] slave using global
  - [ ] service unit


### Why zookeeper 3.5.0-alpha?
Starting in 3.5.0 is possible to resize the zookeeper cluster using (dynamic reconfiguration)[http://zookeeper.apache.org/doc/trunk/zookeeperReconfig.html#ch_reconfig_dyn]

### Required fleet metadata for zookeeper
Any node where is possible to rung zookeeper must have the `role=zookeeper` metadata.

```
  - name: fleet.service
    command: start
    content: |
      [Unit]
      Description=fleet
      Requires=etcd.service
      After=etcd.service

      [Service]
      Environment="FLEET_METADATA=role=zookeeper"
      ExecStart=/usr/bin/fleetd
```