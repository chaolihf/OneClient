package sshclient

import (
	"fmt"
	"testing"
)

func TestSession(t *testing.T) {
	session := NewSshSession("134.95.237.121:2222", "nmread", "Siemens#202405", 10)
	//session := NewSshSession("192.168.100.6:22", "root", "P&ssw0rd@@@", 10)
	if session == nil {
		t.Fail()
		return
	}
	content, err := session.ExecuteMoreCommand("display arp", "---- More ----", "<TDL-JF-9310-1>")

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(content)
	session.Close()
}
