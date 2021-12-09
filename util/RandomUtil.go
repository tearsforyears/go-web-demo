package util

import (
	"bytes"
	"math/rand"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix() % 1000)
}
func RandomStr(len int) string {
	var buffer bytes.Buffer
	for i := 0; i < len; i++ {
		buffer.WriteString(strconv.Itoa(rand.Intn(10)))
	}
	return buffer.String()
}
