package zookeeper

import (
	"io"
	"os/exec"
	"time"
)

type ZkServer struct {
	Stdout, Stderr io.Writer

	cmd *exec.Cmd
}

func (srv *ZkServer) Start() error {
	srv.cmd = exec.Command("/opt/zookeeper/bin/zkServer.sh", "start-foreground")
	srv.cmd.Stdout = srv.Stdout
	srv.cmd.Stderr = srv.Stderr
	return srv.cmd.Start()
}

func (srv *ZkServer) Pid() int {
	return srv.cmd.Process.Pid
}

func (srv *ZkServer) Stop() {
	go func() {
		time.Sleep(1 * time.Second)
		srv.cmd.Process.Kill()
	}()
	srv.cmd.Process.Wait()
}
