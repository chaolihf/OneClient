package http

import "testing"

func TestService(t *testing.T) {
	StartServer()
}

func TestSecurityService(t *testing.T) {
	StartSecurityServer()
}
