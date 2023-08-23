package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
	"io"
	"k8s.io/klog/v2"
	"time"
)

func main() {
	pwd := "123@Conan"
	remoteAddr := "180.184.75.82:22"
	am := ssh.Password(pwd)
	sshConfig := &ssh.ClientConfig{
		User:            "root",
		Auth:            []ssh.AuthMethod{am},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         120 * time.Second,
	}

	conn, err := ssh.Dial("tcp", remoteAddr, sshConfig)
	if err != nil {
		klog.Warningf("remote host is not connected, host %v, error: %v", remoteAddr, err)
	}
	defer func() {
		_ = conn.Close()
	}()

	err = sshmethod(conn)
	if err != nil {
		klog.ErrorS(err, "errrr")
	}
	klog.Info("main end")
}

func sshmethod(conn *ssh.Client) error {
	cmd := "bash /root/sleep.sh"

	session, err := conn.NewSession()
	if err != nil {
		return err
	}
	defer func() {
		_ = session.Close()
	}()

	done := make(chan error)
	var output []byte
	go func() {
		start := time.Now()
		klog.Info("start cmd")
		pipe, err2 := session.StdoutPipe()
		if err2 != nil {
			done <- err
			klog.ErrorS(err2, "StdoutPipe")
			return
		}
		session.Start(cmd)
		reader := bufio.NewReader(pipe)
		for {
			line, err2 := reader.ReadString('\n')
			if err2 != nil || io.EOF == err2 {
				break
			}
			klog.Infof("readline %s", line)
		}
		session.Wait()
		//output, err = session.CombinedOutput(cmd)
		// TODO 日志：命令执行完毕，结果返回
		klog.Info(fmt.Sprintf("run cmd: %s \n~~~~~~~~~~~~~~~~~\n output: %s \n~~~~~~~~~~~~~~~~~\n error:%v \n~~~~~~~~~~~~~~~~~", cmd, string(output), err))
		since := time.Since(start)
		klog.Infof("耗时：%f", since.Seconds())
		if since.Seconds() > 70 {
			err = errors.Errorf("复现拉!!!")
		}
		done <- err
	}()
	select {
	case err = <-done:
		klog.Info("执行完毕")
		return err
	// to avoid the execution stuck, forced to exit the timeout period
	case <-time.After(20 * 60 * time.Second):
		return errors.Errorf("cmd excute timed out, force to exit: %s", cmd)
	}
	return nil
}
