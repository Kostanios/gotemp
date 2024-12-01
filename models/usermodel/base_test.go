package usermodel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkAllocationForNullPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = (*AppUser)(nil)
	}
}

func BenchmarkAllocationForStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = AppUser{}
	}
}

func TestGetPronouns(t *testing.T) {
	pronounce := "something"

	res, err := getPronouns(pronounce)
	assert.ErrorIs(t, ErrInvalidPronounce, err)
	assert.EqualValues(t, Other, res)
}
