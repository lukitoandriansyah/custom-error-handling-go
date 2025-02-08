package struct_model

type PersonStruct struct { //Berperilaku seolah model
	Name       string
	Age        uint
	Gender     string
	Education  EducationStruct
	IsEmployee bool
	Company    EmployeeStruct
}
