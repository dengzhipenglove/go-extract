package goextract

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtractConst(t *testing.T) {

	_, constdata, err := ExtractGoFileConst("./test_data.go")
	require.NoError(t, err)
	//pretty.Println(pkgName)
	//pretty.Println(constdata)

	for _, v := range constdata {
		if v.Name == "Const_String_Data" {
			require.Equal(t, v.ValueString, "str1111\"")
		}
	}
}
