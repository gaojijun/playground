package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

// unmarshal to nil slice
func TestSomething(t *testing.T) {
	testSlice := []string{"a"}
	byts, _ := json.Marshal(testSlice)
	assert.Equal(t, `["a"]`, string(byts))

	var cmd []string
	err := json.Unmarshal(byts, &cmd)
	assert.Nil(t, err)
	assert.Equal(t, len(cmd), 1)
	assert.Equal(t, cmd[0], "a")
}

