package protocol

import (
	"errors"
	"log"
	"net"
)

const (
	VER = 0x05
)

/**
 *	ver	|	nMethods | methods
 *	0x05|	1		 |  1 ~ 255
 */
func HandleHandshake(buf []byte) ([]byte, error) {
	if len(buf) < 3 {
		log.Printf("")
		return nil, errors.New("unmatched length")
	}

	ver := buf[0]
	if ver != VER {
		return nil, errors.New("upsupported protocol version")
	}

	nMethods := buf[1]
	if int(nMethods) != len(buf[2:]) {
		return nil, errors.New("unmatched method length")
	}

	// TODO 选择一种验证方式, 先不需要验证

	return []byte{VER, 0x00}, nil
}

/**
 *	ver	|	cmd	|	rsv	|	aType	|	dstAddr	|	dstPort
 *  0x05|	1	|	0x00|	1		|	n		|	2
 */
func HandleConnect(buf []byte) ([]byte, error) {
	if len(buf) < 6 {
		return nil, errors.New("can't resolve format")
	}

	ver := buf[0]
	if ver != VER {
		return nil, errors.New("unsupported protocol version")
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
		return nil, errors.New("unsupported address type")
	}
	dstPort := buf[4+addrLen:]

	replay, err := HandleDstConnect(cmd, aType, dstAddr, dstPort)
	return replay, err
}

func HandleDstConnect(cmd byte, aType byte, dstAdrss []byte, port []byte) ([]byte, error) {
	targetServ := net.JoinHostPort(ToIpv4(dstAdrss), ToPort(port))
	switch cmd {
	case 0x01:
		// connect
		_, err := net.Dial("tcp", targetServ)
		if err != nil {
			log.Printf("connect to target %v failed, err: %v", targetServ, err)
			return nil, err
		}
		b := []byte{VER, 0x00, 0x00, aType}
		rsp := append(b, dstAdrss...)
		rsp = append(rsp, port...)
		log.Printf("connect to target %v succeed", string(dstAdrss))
		return rsp, nil
	case 0x02:
		// bind TODO
	case 0x03:
		// udp associate TODO
	}
	return nil, errors.New("unsupport cmd")
}
