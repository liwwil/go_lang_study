package main

import "fmt"

type employee struct {
	employeeID   string
	employeeName string
	age          int
	phone        string
}

func main() {
	employee1 := employee{ //call Struct
		employeeID:   "E01",
		employeeName: "John",
		age:          25,
		phone:        "081-1234567",
	}
	fmt.Println("employee1 : ", employee1)

	/*employeeList := [3]employee{} //Struct with Array
	employeeList[0] = employee{
		employeeID:   "E01",
		employeeName: "John",
		age:          25,
		phone:        "081-1234567",
	}
	employeeList[1] = employee{
		employeeID:   "E02",
		employeeName: "Mary",
		age:          26,
		phone:        "082-1234567",
	}
	employeeList[2] = employee{
		employeeID:   "E03",
		employeeName: "Tom",
		age:          27,
		phone:        "083-1234567",
	}
	fmt.Println("employeeList : ", employeeList)*/

	employeeList := []employee{} //Struct with Slice (Slice do not have fixed size)
	employee2 := employee{
		employeeID:   "E01",
		employeeName: "John",
		age:          25,
		phone:        "081-1234567",
	}
	employee3 := employee{
		employeeID:   "E02",
		employeeName: "Jenny",
		age:          25,
		phone:        "081-1234567",
	}
	employee4 := employee{
		employeeID:   "E03",
		employeeName: "Jill",
		age:          25,
		phone:        "081-1234567",
	}
	employeeList = append(employeeList, employee2) //add employee2 to employeeList
	employeeList = append(employeeList, employee3)
	employeeList = append(employeeList, employee4)

	fmt.Println("employeeList : ", employeeList)

}
