package biodata

import (
	"biodata-lukito/struct_model"
)

func BioLukito() struct_model.PersonStruct {
	var value = struct_model.PersonStruct{}
	value.Name = "Lukito Andriansyah"
	value.Age = 25
	value.Gender = "L"
	value.IsEmployee = true

	value.Education = educationLukito()

	if value.IsEmployee {
		value.Company = employeeLukito(true)

	} else {
		value.Company = employeeLukito(false)
	}

	return value
}

func educationLukito() struct_model.EducationStruct {
	var value = struct_model.EducationStruct{}
	value.HighSchool = "SMA NEGERI 84 JAKARTA"
	value.University = append(value.University, "IPB UNIVERSITY")
	return value
}

func employeeLukito(isEmployee bool) struct_model.EmployeeStruct {
	var value = struct_model.EmployeeStruct{}
	if isEmployee {
		value.CompanyName = "DIFINI TEKNOLOGI"
		value.TimeJoining = "2023-06-05"
		value.Division = "IT"
		value.Position = "Developer and RnD"
		value.IsActive = isEmployee
	} else {
		value.CompanyName = ""
		value.TimeJoining = ""
		value.Division = ""
		value.Position = ""
		value.IsActive = isEmployee
	}

	return value
}
