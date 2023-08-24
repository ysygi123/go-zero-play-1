package leetcode

import "strconv"

func RestoreIpAddresses(s string) []string {
	return restoreIpAddresses(s)
}

func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	rbv("", s, 4, &res)
	return res
}

func rbv(pres, s string, level int, res *[]string) {
	if level == 1 {
		if (s[0] == '0' && len(s) > 1) || len(s) > 3 {
			return
		}
		x, _ := strconv.Atoi(s)
		if x > 255 {
			return
		}
		*res = append(*res, pres+s)
		return
	}
	ls := len(s)
	if ls > 3 {
		ls = 3
	}
	for i := 0; i < ls; i++ {
		start := s[0 : i+1]
		if len(start) > 1 && start[0] == '0' {
			continue
		}
		end := s[i+1:]
		big := 3 * (level - 1)
		if len(end) == 0 || len(end) > big {
			continue
		}
		x, _ := strconv.Atoi(start)
		if x > 255 {
			continue
		}
		rbv(pres+start+".", end, level-1, res)
	}
	return
}

//  1  2  3  4  5
//  1   | 2  3  4  5
// 1 | 2 |  3  4  5
// 1 | 23 | 4  5
//  12  |  3  4  5
