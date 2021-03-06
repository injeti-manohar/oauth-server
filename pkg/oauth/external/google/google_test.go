package google

import (
	"testing"

	"github.com/theodor2311/oauth-server/pkg/oauth/external"
)

func TestGoogle(t *testing.T) {
	p, err := NewProvider("google", "clientid", "clientsecret", "", nil)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	_ = external.Provider(p)
}
