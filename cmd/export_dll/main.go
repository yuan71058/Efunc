package main

import "C"

import (
	"encoding/json"
	"unsafe"
)

//export Efunc_Call
func Efunc_Call(name *C.char, params *C.char) *C.char {
	funcName := C.GoString(name)
	jsonParams := C.GoString(params)

	result := globalRegistry.Call(funcName, jsonParams)

	resultJSON, _ := json.Marshal(result)
	return C.CString(string(resultJSON))
}

//export Efunc_Free
func Efunc_Free(ptr unsafe.Pointer) {
	C.free(ptr)
}

//export Efunc_List
func Efunc_List() *C.char {
	list := globalRegistry.List()
	data, _ := json.Marshal(list)
	return C.CString(string(data))
}

//export Efunc_Version
func Efunc_Version() *C.char {
	return C.CString("1.0.0")
}

func init() {
	registerTextFunctions()
	registerFileFunctions()
	registerTimeFunctions()
	registerEncodingFunctions()
	registerCryptoFunctions()
	registerChecksumFunctions()
	registerCoreFunctions()
	registerArrayFunctions()
	registerSystemFunctions()
	registerEnvFunctions()
	registerHTTPFunctions()
}

func main() {}
