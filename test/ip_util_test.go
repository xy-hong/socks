package test

import (
	"socks/protocol"
	"testing"
)

func TestToIpv4(t *testing.T) {
	t.Run("测试 ipv4", func(t *testing.T) {
		bs := []byte{192, 168, 0, 1}
		got := protocol.ToIpv4(bs)
		want := "192.168.0.1"
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestToPort(t *testing.T) {
	t.Run("test ToPort", func(t *testing.T) {
		bs := []byte{0, 222}
		got := protocol.ToPort(bs)
		want := "222"
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("test ToPort 2", func(t *testing.T) {
		bs := []byte{1, 0}
		got := protocol.ToPort(bs)
		want := "256"
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("test ToPort 3", func(t *testing.T) {
		bs := []byte{0b00011110, 0b11010010}
		got := protocol.ToPort(bs)
		want := "7890"
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
