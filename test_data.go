package goextract

// 常量注释
const CONST_ONE = "one" // 常量注释

const CONST_TEST = "test"

// 自定义类型
type TRstType uint32 // 自定义类型

// 新类型
type TResType string // 新类型

// 新类型2
type TResTypeString = string // 新类型2

const (
	//RstOk 注释all 1
	//RstOk 注释all 2
	RstOk TRstType = 1 // RstOk 注释2
	//RstOk 注释all 3
	RstNoChar TRstType = 2 // RstOk 注释3
)
const RstErrData TRstType = 3

const RstErrTransFer = TRstType(5)

//自定义类型测试常量
const ResHaha TResType = "hahahhahah"

const Const_Int_Data1 = 1

const (
	Const_Int_Iota_1 = iota
	Const_Int_Iota_2
)

const Const_String_Data = "str1111"
