package protocol

import (
	"errors"
	"log"
)

func HandleHandshake(buf []byte) ([]byte, error) {
	if len(buf) < 3 {
		log.Printf("")
		return nil, errors.New("")
	}

	ver := buf[0]
	if ver != 0x05 {
		return nil, errors.New("")
	}

	nMethods := buf[1]
	if int(nMethods) != len(buf[2:]) {
		return nil, errors.New("")
	}

	// TODO 选择一种验证方式, 先不需要验证

	return []byte{0x05, 0x00}, nil
}

func HandleConnect(buf []byte) ([]byte, error) {
	return nil, nil
}
