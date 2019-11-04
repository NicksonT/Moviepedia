package moviegetter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMovieYear(t *testing.T) {
	var test MovieGetter
	assert.Equal(t, "2001", test.GetMovie("Spirited Away").Year)
}
