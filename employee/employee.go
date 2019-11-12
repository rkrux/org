package employee

import "fmt"

type findCCM struct {
	ccm         *Employee
	foundFirst  bool
	foundSecond bool
}

//Employee structure : an employee is a manager if he has more than 1 Direct under him - (len(Employee.Directs) > 0)
type Employee struct {
	ID      int
	Directs []*Employee
}

//Create forms the heirarchy among employees recursively
func Create(heirarchy map[int][]int, id int) *Employee {
	//create current employee
	currentEmployee := &Employee{ID: id}

	//check if current employee has directs
	if directs, ok := heirarchy[id]; ok {
		//add directs under the current employee
		currentEmployee.Directs = make([]*Employee, len(directs))

		//create directs of the current employee
		for i := 0; i < len(directs); i++ {
			currentEmployee.Directs[i] = Create(heirarchy, directs[i])
		}
	}
	return currentEmployee
}

//Display displays organization heirarchy recursively
func (e *Employee) Display() {
	//display current employee & his directs IDs
	fmt.Printf("EmployeeID: %d, with directs	->", e.ID)
	for i := 0; i < len(e.Directs); i++ {
		fmt.Printf("	%d", e.Directs[i].ID)
	}
	fmt.Println()

	//traverse his directs
	for i := 0; i < len(e.Directs); i++ {
		e.Directs[i].Display()
	}
}

func (e *Employee) traverse(first, second int) findCCM {
	result := findCCM{nil, false, false}

	//if current employee is one of the employees to be searched for
	if e.ID == first {
		result.foundFirst = true
	}
	if e.ID == second {
		result.foundSecond = true
	}

	for i := 0; i < len(e.Directs); i++ {
		//traverse each direct of the manager
		directResult := e.Directs[i].traverse(first, second)

		//return as-is if CCM found
		if directResult.ccm != nil {
			return directResult
		}

		//update the flags from the results of traversing the direct
		result.foundFirst = result.foundFirst || directResult.foundFirst
		result.foundSecond = result.foundSecond || directResult.foundSecond

		//if both the employees found, then current employee is CCM
		if (result.foundFirst == true) && (result.foundSecond == true) {
			result.ccm = e
			return result
		}
	}

	return result
}

//Traverse traverses organization heirarchy through Founder
func (e *Employee) Traverse(first, second int) *Employee {
	return e.traverse(first, second).ccm
}
