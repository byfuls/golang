package child

//package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"syscall"
)

type Process struct {
	onoff   bool
	child   *exec.Cmd
	pipeIn  io.WriteCloser
	pipeOut io.ReadCloser
	writer  *bufio.Writer
	reader  *bufio.Reader
}

func (p *Process) Read(buf *[]byte) (int, error) {
	if !p.onoff {
		return 0, errors.New("process does not exist")
	}

	len, err := p.reader.Read(*buf)
	if err != nil {
		if err == io.EOF {
			fmt.Println("[Read] pipe has been closed, error: ", err)
			return 0, err
		} else {
			fmt.Println("[Read] pipe read content failed, error: ", err)
			return 0, err
		}
	}
	return len, err
}

func (p *Process) Write(buf []byte) (int, error) {
	if !p.onoff {
		return 0, errors.New("process does not exist")
	}

	len, err := p.writer.Write(buf)
	if err != nil {
		fmt.Println("[Write] pipe write error: ", err)
		return 0, err
	}
	p.writer.Flush()
	return len, err
}

func (p *Process) Wait() error {
	return p.child.Wait()
}

func (p *Process) Run(fullPath string, ip string, port int) error {
	if p.onoff {
		return errors.New("Already process is running")
	}
	p.child = exec.Command(fullPath, "-ip", ip, "-port", strconv.Itoa(port))

	var err error
	p.pipeIn, err = p.child.StdinPipe()
	if err != nil {
		return nil
	}
	p.pipeOut, err = p.child.StdoutPipe()
	if err != nil {
		return nil
	}
	p.writer = bufio.NewWriter(p.pipeIn)
	p.reader = bufio.NewReader(p.pipeOut)

	if err := p.child.Start(); err != nil {
		return err
	}
	p.onoff = true
	return nil
}

func (p *Process) Status() bool {
	return p.onoff
}

func (p *Process) SetStatus(onoff bool) {
	p.onoff = onoff
}

func (p *Process) Stop() error {
	if !p.onoff {
		return errors.New("process does not exist")
	}
	pcs, err := os.FindProcess(p.child.Process.Pid)
	if err != nil {
		return err
	}
	err = pcs.Signal(syscall.SIGINT)
	if err != nil {
		return err
	}
	err = p.child.Wait()
	if err != nil {
		return err
	}

	//err := syscall.Kill(p.child.Process.Pid, syscall.SIGKILL)
	//if err != nil {
	//	fmt.Printf("[Stop:%d] end child(%d) process %s error: %s\n", os.Getpid(), p.child.Process.Pid, p.child.Path, err)
	//}
	fmt.Printf("[Stop:%d] end child(%d) process %s success\n", os.Getpid(), p.child.Process.Pid, p.child.Path)
	p.pipeIn.Close()
	p.pipeOut.Close()
	p.pipeIn = nil
	p.pipeOut = nil
	p.writer = nil
	p.reader = nil
	p.child = nil
	p.onoff = false
	return nil
}

func Done(p *Process) {
	p = nil
}

func Init() *Process {
	process := new(Process)

	process.onoff = false

	return process
}
