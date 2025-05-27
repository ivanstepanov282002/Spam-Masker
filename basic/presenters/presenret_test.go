package presenter_test

import (
	presenter "Spam-Masker/basic/presenters"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPresent(t *testing.T) {
	createFile := t.TempDir() + "mask_test.txt"
	content := "line1\nline2\nline3"
	os.WriteFile(createFile, []byte(content), 0644)
	nameFile := &presenter.FilePresenter{FilePath: createFile}
	sliceStringContent := []string{"line1\nline2\nline3"}
	err := nameFile.Present(sliceStringContent)
	assert.NoError(t, err)
}

