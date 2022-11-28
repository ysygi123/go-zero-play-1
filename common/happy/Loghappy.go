package happy

import (
	"errors"
	"fmt"
	"os"
	"syscall"
	"time"
)

type LogHappy struct {
	MessageChan chan *[]byte
	FileName    string
	nowSize     int
	MaxSize     int
	fsyncSize   int
	file        *os.File
	theByte     []byte
	//缓存用
	logByte         []byte
	FlushNow        chan struct{}
	operateDuration time.Duration
	t               time.Time
	null            struct{}
}

func (l *LogHappy) SetFileName(filename string) *LogHappy {
	l.FileName = filename
	return l
}
func (l *LogHappy) SetMaxSize(maxSize int) *LogHappy {
	l.MaxSize = maxSize
	return l
}
func (l *LogHappy) InitSpace() (err error) {
	l.nowSize = 0
	l.MessageChan = make(chan *[]byte, 10000)
	l.FlushNow = make(chan struct{}, 1)
	l.t = time.Now()
	l.logByte = make([]byte, 0, 1024*1024)
	if len(l.FileName) == 0 {
		return errors.New("err no filename")
	}
	l.file, err = os.OpenFile(l.FileName, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return
	}
	err = syscall.Ftruncate(int(l.file.Fd()), int64(l.MaxSize))
	if err != nil {
		return
	}
	l.theByte, err = syscall.Mmap(int(l.file.Fd()), 0, l.MaxSize, syscall.PROT_WRITE|syscall.PROT_READ, syscall.MAP_SHARED)
	if err != nil {
		return
	}
	go l.start()
	return
}

func (l *LogHappy) start() {
	var err error
	for {
		select {
		case <-time.After(1 * time.Second):
			err = l.fsync()
			if err != nil {
				goto DEAD
			}
			fmt.Println("l.operateDuration", l.operateDuration, "now fsyncSIze", l.fsyncSize, "write size", l.nowSize-l.fsyncSize, "len(theByte) : ", len(l.theByte))
		case msg, ok := <-l.MessageChan:
			if !ok {
				goto DEAD
			}
			//fmt.Println("查看来的消息", string(*msg))
			copy(l.theByte[(l.nowSize):(l.nowSize+len(*msg))], *msg)
			l.nowSize += len(*msg)
			//fmt.Println("查看len(*msg)和l.nowSize", len(*msg), l.nowSize)
		case _, ok := <-l.FlushNow:
			if !ok {
				goto DEAD
			}
			err = l.fsync()
			if err != nil {
				goto DEAD
			}
		}
	}
DEAD:
}

func (l *LogHappy) fsync() (err error) {
	if l.nowSize <= l.fsyncSize {
		return
	}
	l.t = time.Now()
	err = syscall.Fsync(int(l.file.Fd()))
	l.operateDuration += time.Since(l.t)
	if err != nil {
		return
	}

	l.fsyncSize = l.nowSize
	fmt.Println("操作时间", l.operateDuration)
	return
}
