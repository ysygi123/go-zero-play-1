package task_test

import "sync"

type TaskNode struct {
	HandleFunc    func() error
	HandleErrFunc func(err error)
	Next          *TaskNode
}

type TaskHead struct {
	Tn    []*TaskNode
	Use   []uint
	Index uint
	Lock  sync.Mutex
}

func NewTaskHead(gNum int) *TaskHead {
	t := &TaskHead{
		Tn:   make([]*TaskNode, gNum),
		Use:  make([]uint, gNum),
		Lock: sync.Mutex{},
	}
	return t
}

func (t *TaskHead) Add(handleFunc func() error, handleErrFunc func(err error)) {

}
