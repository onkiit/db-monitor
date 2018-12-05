package mongo

import (
	"testing"
)

func TestOpen(t *testing.T) {
	err := Open("localhost:270172")
	if err != nil {
		t.Error(err)
	}
}
