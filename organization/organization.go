package organization

import (
	"fmt"

	emp "github.com/rkrux/org/employee"
)

//Organization structure
type Organization struct {
	TotalEmployees int
	Founder        *emp.Employee
}

//Create forms the organization heirarchy
func (o *Organization) Create(totalEmployees int, heirarchy map[int][]int) {
	o.TotalEmployees = (totalEmployees)

	if o.TotalEmployees > 0 {
		//form heirarchy among employees
		o.Founder = emp.Create(heirarchy, 1)
	}
}

//Display displays organization structure
func (o Organization) Display() {
	fmt.Println("Organization Heirarchy")
	if o.Founder == nil {
		fmt.Println("Empty Organization")
		return
	}
	o.Founder.Display()
}

func (o Organization) checkInputs(first, second int) bool {
	if (first < 1) || (first > o.TotalEmployees) {
		return false
	}
	if (second < 1) || (second > o.TotalEmployees) {
		return false
	}
	if first == second {
		return false
	}
	return true
}

//FindClosestCommonManager finds ClosestCommonManager of 2 employees
func (o Organization) FindClosestCommonManager(first, second int) *emp.Employee {
	if o.checkInputs(first, second) == false {
		return nil
	}
	return o.Founder.Traverse(first, second)
}
