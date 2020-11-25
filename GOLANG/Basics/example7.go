package main

import "fmt"

func main() {

	StudentRoll := make(map[int]string)
	StudentRoll[10] = "Febin"
	StudentRoll[5] = "Favas"
	StudentRoll[12] = "Murshid"
	StudentRoll[15] = "Rizwan"
	StudentRoll[20] = "Shibil"

	fmt.Println(StudentRoll[12])

	Employee := map[string]map[string]string{
		"Febin": map[string]string{
			"Position":   "Trainee",
			"EmployeeId": "20",
		},
		"Favas": map[string]string{
			"Position":   "Developer",
			"EmployeeId": "15",
		},
	}

	if temp, name := Employee["Febin"]; name {
		fmt.Println(temp["Position"], temp["EmployeeId"])
	}
}
