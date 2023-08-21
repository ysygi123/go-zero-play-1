package leetcode

import "strings"

func RestoreIpAddresses(s string) []string {
	return restoreIpAddresses(s)
}

func restoreIpAddresses(s string) []string {
	return rbv("", s, 4)
}

func rbv(preS, s string, level int) []string {
	if level > 4 || level < 1 {
		return []string{}
	}
	ls := len(s)
	if ls > level*3 || ls < level {
		return []string{}
	}
	if ls == level {
		x := ""
		for i := 0; i < ls; i++ {
			x += string(s[i]) + "."
		}
		x = strings.Trim(x, ".")
		return []string{x}
	}
	var rts []string
	if ls >= 3 {
		r1 := rbv(s[3:], level-1)
		r2 := rbv(s[2:], level-1)
		r3 := rbv(s[1:], level-1)
		rts = append(append(append(rts, r1...), r2...), r3...)
	}
	if ls >= 2 {
		r1 := rbv(s[2:], level-1)
		r2 := rbv(s[1:], level-1)
		rts = append(append(rts, r1...), r2...)
	}
	if ls >= 1 {
		r1 := rbv(s[1:], level-1)
		rts = append(rts, r1...)
	}
	return []string{}
}

// 1 2 3 4 5   l 4
// 2 3 4 5   l 3
//  3  4  5  l 2
//   4  5  l 1
