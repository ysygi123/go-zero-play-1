package happy

import (
	"errors"
	"fmt"
	"os"
	"syscall"
)

const maxMapSize = 0x8000000000
const maxMmapStep = 1 << 30 // 1GB
const theFileName = "my.log"

func MyMmapRun() (err error) {
	file, err := os.OpenFile(theFileName, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return
	}
	defer file.Close()
	stat, err := os.Stat(theFileName)
	if err != nil {
		return
	}
	size, err := mmapSize(int(stat.Size()))
	fmt.Println(size)
	if err != nil {
		return
	}
	err = syscall.Ftruncate(int(file.Fd()), 1024)
	if err != nil {
		return
	}
	size = 100
	b, err := syscall.Mmap(int(file.Fd()), 0, size, syscall.PROT_WRITE|syscall.PROT_READ, syscall.MAP_SHARED)
	if err != nil {
		return
	}
	for i := 0; i < 20; i++ {
		b[i] = '0'
	}
	err = syscall.Munmap(b)
	if err != nil {
		return
	}
	return
}

func mmapSize(size int) (int, error) {
	for i := uint(15); i <= 30; i++ {
		if size <= 1<<i {
			return 1 << i, nil
		}
	}

	if size > maxMapSize {
		return 0, errors.New("你妈炸")
	}
	sz := int64(size)

	if remainder := sz % int64(maxMmapStep); remainder > 0 {
		sz += int64(maxMmapStep) - remainder
	}

	pageSize := int64(os.Getpagesize())
	if sz%pageSize != 0 {
		sz = ((sz / pageSize) + 1) * pageSize
	}
	if sz > maxMapSize {
		sz = maxMapSize
	}
	return int(sz), nil
}
