package x

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCore_AllowOriginFunc(t *testing.T) {
	tests := []struct {
		origin    string
		isAllowed bool
	}{
		{"https://stein.systems", true},
		{"https://management.stein.systems", true},
		{"https://www.stein.systems", true},
		{"https://stein.systems.com", false},
	}
	for _, test := range tests {
		t.Run(test.origin, func(t *testing.T) {
			allowOrigin := allowOriginFunc("stein.systems")
			isAllowed := allowOrigin(test.origin)
			assert.Equal(t, test.isAllowed, isAllowed)
		})
	}
}
