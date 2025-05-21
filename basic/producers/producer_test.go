package producer_test

import (
	producer "Spam-Masker/basic/producers"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProduce(t *testing.T) {
	createFile := t.TempDir() + "test.txt"
	content := "line1\nline2\nline3"
	os.WriteFile(createFile, []byte(content), 0644)
	nameFile := &producer.FileProducer{FilePath: createFile}
	result, err := nameFile.Produce()
	assert.NoError(t, err)
	assert.Equal(t, []string{"line1\nline2\nline3"}, result, "Result should match expected lines")
}
