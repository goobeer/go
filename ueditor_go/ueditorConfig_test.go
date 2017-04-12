package ueditor_go

import (
	"net"
	"testing"
)

func TestUeditorGO(t *testing.T) {
	config, err := BuildItems()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(config)
	}
}

func TestNet(t *testing.T) {
	ip := net.ParseIP("192.168.1.199")

}
