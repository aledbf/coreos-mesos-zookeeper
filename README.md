# coreos-mesos-zookeeper

Services and images to start mesos and zookeeper in coreos

- [X] zookeeper
  - [X] use alpine linux
  - [X] service unit
  - [X] add node to the cluster on start
  - [X] remove node from the cluster on restart

- [X] mesos
  - [X] master
  - [X] service unit
  - [X] slave using global
  - [X] service unit
- [X] marathon

### Why zookeeper 3.5.0-alpha?
Starting in 3.5.0 is possible to resize the zookeeper cluster using [dynamic reconfiguration](http://zookeeper.apache.org/doc/trunk/zookeeperReconfig.html#ch_reconfig_dyn)

### Required fleet metadata for zookeeper
Any node where is possible to rung zookeeper must have the `zookeeper=true` metadata.

```
  - name: fleet.service
    command: start
    content: |
      [Unit]
      Description=fleet
      Requires=etcd.service
      After=etcd.service

      [Service]
      Environment="FLEET_METADATA=zookeeper=true"
      ExecStart=/usr/bin/fleetd
```

### Steps to start the zookeeper cluster

The folder units contains the required services to run the service.
The file `zookeeper@1.service` is the template required to run the zookeeper server in one node.
To run the zookeeper cluster is required at least 3 nodes.
Copy the template `zookeeper@1.service` to `zookeeper@2.service` and `zookeeper@3.service` and submit the units to `fleet`:

```
fleetctl submit zookeeper@1.service zookeeper@2.service zookeeper@3.service
```

and start the services

```
fleetctl start zookeeper@1.service zookeeper@2.service zookeeper@3.service

```

This will add the next key/values to etcd:
```
/zookeeper
/zookeeper/nodes
/zookeeper/nodes/<ip node 1>
/zookeeper/nodes/<ip node 1>/id
/zookeeper/nodes/<ip node 2>
/zookeeper/nodes/<ip node 2>/id
/zookeeper/nodes/<ip node 3>
/zookeeper/nodes/<ip node 3>/id
```

Why?
Zookeeper need unique identifiers for every node so when the container starts it will check if this "mapping" exists or it will create a new one. This keys do not use any TTL (are permanent).


### Restarting zookeeper nodes
During the start of the zookeeper service in the node it will register the information in the cluster if it does not exists and during the stop it will remove the node from the cluster.
The information to register follows the format defined [here](http://zookeeper.apache.org/doc/trunk/zookeeperReconfig.html#sc_reconfig_clientport):
```
server.$NODE_ID=$HOST_IP:2181:2888:participant;$HOST_IP:3888"
```

