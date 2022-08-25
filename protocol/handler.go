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
	if len(buf) < 6 {
		log.Printf("")
		return nil, errors.New("")
	}

	ver := buf[0]
	if ver != 0x05 {
		return nil, errors.New("Unsupport protocol version.")
	}

	cmd := buf[1]
	aType := buf[3]
	var addrLen int
	var dstAddr []byte
	switch aType {
	case 0x01:
		// ipv4
		addrLen = 4
		dstAddr = buf[4 : 4+addrLen]
	case 0x03:
		// domain
		nameLen := int(buf[5])
		addrLen = 1 + nameLen
		dstAddr = buf[5 : 4+addrLen]
	case 0x04:
		// ipv6
		addrLen = 16
		dstAddr = buf[4 : 4+addrLen]
	default:
		return nil, errors.New("Unsupport address type.")
	}
	dstPort := buf[4+addrLen:]

	return nil, nil
}

func HandleDstConnect(cmd byte, dstAdrss []byte, port [2]byte) ([]byte, error) {
	switch cmd {
	case 0x01:
		// connect
	case 0x02:
		// bind
	case 0x03:
		// udp associate
	default:
		return nil, errors.New("Unsupport cmd.")
	}
}
