package struct_response

import (
	"biodata-lukito/struct_model"
)

type SuccessStruct struct { //Berperilaku seolah model
	Message string
	Data    struct_model.PersonStruct
	Code    uint
}
