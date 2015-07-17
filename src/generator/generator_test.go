package generator

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestRandStr(t *testing.T) {
	var length int64 = 10
	passInfo := PassInfo{"alpha", length}
	pass := RandStr(passInfo)

	assert.Equal(t, length, int64(len(pass)))
	assert.NotEqual(t, pass, RandStr(passInfo), "Password isn't random")
}

func TestGetDictionary(t *testing.T) {
	var set string = "alpha"

	assert.Equal(t, char_alpha, getDictionary(set))
}