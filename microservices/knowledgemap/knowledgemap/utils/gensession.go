package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"
)

func Random24() string {
	src := fmt.Sprintf("%v%v%v%v%v", time.Now().String(), time.Now().Unix(), rand.Int(), time.Now().Nanosecond(), time.Now().String())
	srcbyte := []byte(src)
	for idx := range srcbyte {
		srcbyte[idx]++
	}
	s := base64.StdEncoding.EncodeToString(srcbyte)
	sbyte := []byte(s)
	buffer := new(bytes.Buffer)
	for idx := range sbyte {
		if idx == 24 {
			break
		}
		buffer.WriteByte(sbyte[(idx*idx+int(time.Now().Nanosecond()))%len(sbyte)])
	}
	return buffer.String()
}

// +/= 这三个符号要转变
func GenSession(uid string) string {
	randomstr := Random24()
	ranbytes := []byte(randomstr)
	uidbytes := []byte(uid)
	var idx int
	buffer := new(bytes.Buffer)
	for {
		if idx == 48 {
			break
		}
		if idx%2 == 0 { // random
			if ranbytes[idx/2] == '+' || ranbytes[idx/2] == '/' || ranbytes[idx/2] == '=' {
				buffer.WriteByte('x')
			} else {
				buffer.WriteByte(ranbytes[idx/2])
			}
		} else { // uid
			buffer.WriteByte(uidbytes[idx/2])
		}
		idx++
	}
	return buffer.String()
}
func GetUIDFrSession(sessionstr string) string {
	sessionbytes := []byte(sessionstr)
	var idx int
	buffer := new(bytes.Buffer)
	for {
		if buffer.Len() == 24 {
			break
		}
		buffer.WriteByte(sessionbytes[idx+1])
		idx += 2
	}
	return buffer.String()
}
