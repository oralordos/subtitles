package subtitles

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateMovieHashFromMovieFile(t *testing.T) {

	fileName := createTempFile(1024 * 1024 * 2)

	hash, err := createMovieHashFromMovieFile(fileName)

	assert.Equal(t, nil, err)
	assert.Equal(t, "38a503307786991a982f8ded498b90e0", hash)

	os.Remove(fileName)
}

func TestDownloadFromTheSubDb(t *testing.T) {
	fileName := createZeroedTempFile(1024 * 1024 * 2)

	text, err := fromTheSubDb(fileName, "en", "sandbox.thesubdb.com")
	assert.Equal(t, nil, err)
	assert.True(t, len(text) > 1000)

	os.Remove(fileName)
}