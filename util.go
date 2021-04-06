// Copyright (C) 2016-2018 Teamwork.com
// All rights reserved.

package tnef

func byteToInt(data []byte) int {
	var num int
	var n uint
	for _, b := range data {
		num += (int(b) << n)
		n += 8
	}
	return num
}
