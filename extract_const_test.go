package gotool

import (
	"testing"

	"github.com/kr/pretty"
	"github.com/stretchr/testify/assert"
)

func TestExtractConst(t *testing.T) {

	pkgName, constdata, err := ExtractGoFileConst("./test_data.go", "")
	assert.Empty(t, err)
	pretty.Print(pkgName)
	pretty.Print(constdata)

}
