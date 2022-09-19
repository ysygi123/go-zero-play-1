package main

import (
	"fmt"
	"go-zero-play-1/common/queue/syrocketmq/rocketmq_i"
)

func main() {
	fmt.Println("############################ \n#                          # \n#       start              # \n#                          # \n############################ ")
	rocketmq_i.GrocketmqManager.Start()
}
