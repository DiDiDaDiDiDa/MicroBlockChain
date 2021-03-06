package main

import (
	"bytes"
	"encoding/binary"
	"log"
)

func Int2Byte(num int64) (ret []byte) {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	ret = buff.Bytes()

	return
}
