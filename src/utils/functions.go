package utils

import (
	"github.com/gin-gonic/gin"
	"bytes"
	"encoding/binary"
	"strconv"
)

func Response(code int, msg string, data interface{}) interface{} {
	return gin.H{
		"state": code,
		"msg":   msg,
		"data":  data,
	}
}

//字节转换成整形
func BytesToInt(b []byte) int {
	var tmp int
	bytesBuffer := bytes.NewBuffer(b)
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return int(tmp)
}

func BytesToString(b []byte) string {
	var tmp string
	for _, v := range b {
		tmp += strconv.Itoa(int(v))
	}
	return tmp
}
