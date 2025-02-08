package main

import (
	"biodata-lukito/biodata"
	"biodata-lukito/struct_response"
	"biodata-lukito/util"
	"fmt"
)

func main() {
	util.SafeExecute(func() { // metodmain dijalanakna dalam method safeExecution untuk menghindari error yang tidak di inginlkan (Selain error bisnins)
		var success = struct_response.SuccessStruct{}
		success.Message = "Successfully to Get Data"
		success.Code = 200
		success.Data = biodata.BioLukito()
		fmt.Println(util.StructToJson(success))
	})
}
