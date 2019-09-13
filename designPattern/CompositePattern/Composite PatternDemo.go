package main

import "fmt"

type Employee struct {
	name         string
	dept         string
	salary       string
	subordinates []*Employee
}

func (e *Employee) add(ele *Employee) {
	e.subordinates = append(e.subordinates, ele)

}

func (e *Employee) getSubordinates() []*Employee {
	return e.subordinates
}
func (e *Employee) printEmployee() {
	fmt.Println("Employee :[ Name : " + e.name + ", dept : " + e.dept + ", salary :" + e.salary + " ]")
}

func main() {
	CEO := Employee{name: "John", dept: "CEO", salary: "30000"}

	headSales := Employee{name: "Robert", dept: "Head Sales", salary: "20000"}
	headMarketing := Employee{name: "Michel", dept: "Head Marketing", salary: "20000"}

	clerk1 := Employee{name: "Laura", dept: "Marketing", salary: "10000"}
	clerk2 := Employee{name: "Bob", dept: "Marketing", salary: "10000"}

	salesExecutive1 := Employee{name: "Richard", dept: "Sales", salary: "10000"}
	salesExecutive2 := Employee{name: "Rob", dept: "Sales", salary: "10000"}

	CEO.add(&headSales)
	CEO.add(&headMarketing)

	headMarketing.add(&clerk1)
	headMarketing.add(&clerk2)

	headSales.add(&salesExecutive1)
	headSales.add(&salesExecutive2)

	CEO.printEmployee()
	for _, headEmployee := range CEO.getSubordinates() {

		headEmployee.printEmployee()
	
		for _, employee := range headEmployee.getSubordinates() {

			employee.printEmployee()
		}
	}
}
