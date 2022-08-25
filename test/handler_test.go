package test

import (
	"bytes"
	"socks/protocol"
	"testing"
)

func TestHandleHandshake(t *testing.T) {
	t.Run("正常流程", func(t *testing.T) {
		b := []byte{0x05, 0x01, 0x00}
		got, _ := protocol.HandleHandshake(b)
		want := []byte{0x05, 0x00}
		if !bytes.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("错误流程:握手请求短", func(t *testing.T) {
		b := []byte{0x05}
		_, err := protocol.HandleHandshake(b)
		if err == nil {
			t.Error("")
		}
	})
}
