package main

import (
	"strconv"
	"sync"
)

var (
	itemBytes = []byte{0x69, 0x74, 0x65, 0x6d, 0x2d}
	commaByte = byte(0x2c)
)

var bufPool = sync.Pool{
	New: func() any {
		b := make([]byte, 0, 2*1024*1024)
		return &b
	},
}

func calculateResult(num int) string {

	bp := bufPool.Get().(*[]byte)
	buf := (*bp)[:0]

	var intBuf [20]byte
	for i := range num {
		buf = append(buf, itemBytes...)
		b := strconv.AppendInt(intBuf[:0], int64(i), 10)
		buf = append(buf, b...)
		buf = append(buf, commaByte)
	}

	result := string(buf)
	*bp = buf
	bufPool.Put(bp)
	return result
}
