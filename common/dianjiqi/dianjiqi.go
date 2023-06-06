package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

const (
	CommandStart = "s"
	CommandGo    = "go"
	CommandExit  = "exit"
	CommandX     = "x"
)

var text = fmt.Sprintf("请输入需要的操作\n %s： 输入四点矩阵范围；\n %s：开始执行任务；\n %s： 结束", CommandStart, CommandGo, CommandExit)

type jvzhen struct {
	DoubleX  []int
	DoubleY  []int
	NextTime time.Duration
}

func main() {
	jihe := make([]jvzhen, 0)
	var command string
	rand.Seed(time.Now().UnixNano())
	for {
		fmt.Println(text)

		n, err := fmt.Scan(&command)
		if err != nil {
			panic(err)
		}
		if n == 0 {
			continue
		}

		if command == CommandStart {
			pt := jvzhen{
				DoubleX: []int{},
				DoubleY: []int{},
			}
			for i := 0; i < 4; i++ {
				var cd string
				fmt.Printf("请把鼠标放到第 %d 个点 输入x进行采集\n", i+1)
				_, _ = fmt.Scan(&cd)
				if cd != CommandX {
					fmt.Println("输入正确的值啊")
					i--
					continue
				}

				ptt := [2]int{}
				ptt[0], ptt[1] = robotgo.GetMousePos()
				pt.DoubleY = append(pt.DoubleY, ptt[1])
				pt.DoubleX = append(pt.DoubleX, ptt[0])
			}
			for {
				fmt.Println("输入点击完此框后需要等待的ms数量")
				var times string
				_, _ = fmt.Scan(&times)
				if times == "" {
					fmt.Println("重新输入")
					continue
				}
				t, e := strconv.ParseInt(times, 10, 64)
				if e != nil {
					fmt.Println(e)
					continue
				}
				pt.NextTime = time.Duration(t)
				break
			}

			sort.Slice(pt.DoubleX, func(i, j int) bool {
				return pt.DoubleX[i] > pt.DoubleX[j]
			})
			sort.Slice(pt.DoubleY, func(i, j int) bool {
				return pt.DoubleY[i] > pt.DoubleY[j]
			})
			jihe = append(jihe, pt)
		} else if command == CommandExit {
			break
		} else if command == CommandGo {
			fmt.Println("输入你需要执行的次数 : ")
			var times string
			_, _ = fmt.Scan(&times)
			timesInt, _ := strconv.Atoi(times)
			fmt.Println("执行次数为 : ", timesInt)
			for i := 0; i < timesInt; i++ {
				for _, pts := range jihe {
					time.Sleep((pts.NextTime + time.Duration(getrand(1, 2000))) * time.Millisecond)
					x := getrand(pts.DoubleX[3], pts.DoubleX[0])
					y := getrand(pts.DoubleY[3], pts.DoubleY[0])
					robotgo.Move(x, y)
					robotgo.Click()
				}
			}
		}
	}
	fmt.Println("over!!!")
}

func getrand(start, end int) int {
	return rand.Intn(end-start) + start
}
