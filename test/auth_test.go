package test

import (
	"testing"

	"github.com/RoyJoel/TennisMomentBackEnd/package/web/auth"
)

func TestAuth(t *testing.T) {
	auth.DeletePolicy("1", "", "*")
}
