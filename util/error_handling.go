package util

import (
	"biodata-lukito/struct_response"
	"fmt"
	"runtime/debug"
)

// Wrapper function untuk menangkap panic
func SafeExecute(fn func()) {
	defer func() {
		if r := recover(); r != nil {
			var error = struct_response.ErrorStruct{}
			error.Message = fmt.Sprintf("Terjadi error: %v", r)
			error.Code = 500
			error.Data = string(debug.Stack())
			fmt.Println(StructToJson(error))
			return
		}
	}()
	fn()
}
