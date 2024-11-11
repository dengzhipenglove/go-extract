package goextract

import (
	"fmt"
	"testing"

	"github.com/kr/pretty"
	"github.com/stretchr/testify/assert"
)

func TestExtractConst(t *testing.T) {

	pkgName, constdata, err := ExtractGoFileConst("./test_data.go")
	assert.Empty(t, err)
	pretty.Println(pkgName)
	pretty.Println(constdata)

	for _, v := range constdata {
		if v.Name == "Const_String_Data" {
			assert.Equal(t, v.ValueString, "str1111\"")
		}
	}
}
