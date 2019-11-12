package testscenarios

import (
	"fmt"
	"testing"

	emp "github.com/rkrux/org/employee"
	org "github.com/rkrux/org/organization"
)

type testRow struct {
	emp1 int
	emp2 int
	ccm  *emp.Employee
}

func buildTestScenario(o *org.Organization, orgSize int, heirarchy map[int][]int) {
	o.Create(orgSize, heirarchy)
	o.Display()
}

func assertTestScenario(t *testing.T, table []testRow, o org.Organization) {
	for _, row := range table {
		result := o.FindClosestCommonManager(row.emp1, row.emp2)
		if row.ccm == nil && result != nil {
			t.Errorf("Closest Common Manager of %d and %d was incorrect, got: %d, want: nil", row.emp1, row.emp2, result.ID)
		} else if row.ccm != nil && result == nil {
			t.Errorf("Closest Common Manager of %d and %d was incorrect, got: nil, want: %d", row.emp1, row.emp2, row.ccm.ID)
		}
		if row.ccm != nil && result.ID != row.ccm.ID {
			t.Errorf("Closest Common Manager of %d and %d was incorrect, got: %d, want: %d", row.emp1, row.emp2, result.ID, row.ccm.ID)
		}
	}
}

//empty organization
func TestScenario1(t *testing.T) {
	fmt.Println("Test Scenario: 1")

	heirarchy := make(map[int][]int)

	o := org.Organization{}
	buildTestScenario(&o, 0, heirarchy)
	table := []testRow{
		{0, 9, nil},
		{1, 1, nil},
		{1, 2, nil},
	}
	assertTestScenario(t, table, o)
}

//organization with just founder
func TestScenario2(t *testing.T) {
	fmt.Println("Test Scenario: 2")

	heirarchy := make(map[int][]int)
	heirarchy[1] = []int{}

	o := org.Organization{}
	buildTestScenario(&o, 1, heirarchy)

	table := []testRow{
		{0, 9, nil},
		{1, 1, nil},
		{1, 2, nil},
	}
	assertTestScenario(t, table, o)
}

//skewed organization
func TestScenario3(t *testing.T) {
	fmt.Println("Test Scenario: 3")

	heirarchy := make(map[int][]int)
	heirarchy[1] = []int{2}
	heirarchy[2] = []int{3}
	heirarchy[3] = []int{4}
	heirarchy[4] = []int{5}

	o := org.Organization{}
	buildTestScenario(&o, 5, heirarchy)

	table := []testRow{
		{0, 9, nil},
		{1, 1, nil},
		{1, 2, &emp.Employee{ID: 1}},
		{3, 2, &emp.Employee{ID: 2}},
		{4, 3, &emp.Employee{ID: 3}},
		{5, 1, &emp.Employee{ID: 1}},
	}
	assertTestScenario(t, table, o)
}

//equally distributed organization
func TestScenario4(t *testing.T) {
	fmt.Println("Test Scenario: 4")

	heirarchy := make(map[int][]int)
	heirarchy[1] = []int{2, 3, 4}
	heirarchy[2] = []int{5, 6, 7}
	heirarchy[3] = []int{8, 9, 10}
	heirarchy[4] = []int{11, 12, 13}

	o := org.Organization{}
	buildTestScenario(&o, 13, heirarchy)

	table := []testRow{
		{0, 20, nil},
		{1, 1, nil},
		{1, 2, &emp.Employee{ID: 1}},
		{3, 2, &emp.Employee{ID: 1}},
		{5, 6, &emp.Employee{ID: 2}},
		{7, 3, &emp.Employee{ID: 1}},
		{5, 13, &emp.Employee{ID: 1}},
		{11, 4, &emp.Employee{ID: 4}},
		{1, 9, &emp.Employee{ID: 1}},
		{10, 8, &emp.Employee{ID: 3}},
	}
	assertTestScenario(t, table, o)
}

//unequally distributed organization
func TestScenario5(t *testing.T) {
	fmt.Println("Test Scenario: 5")

	heirarchy := make(map[int][]int)
	heirarchy[1] = []int{2, 3}
	heirarchy[2] = []int{4, 5}
	heirarchy[4] = []int{6, 7}
	heirarchy[5] = []int{8}

	o := org.Organization{}
	buildTestScenario(&o, 8, heirarchy)

	table := []testRow{
		{0, 10, nil},
		{1, 1, nil},
		{1, 2, &emp.Employee{ID: 1}},
		{3, 2, &emp.Employee{ID: 1}},
		{5, 6, &emp.Employee{ID: 2}},
		{7, 3, &emp.Employee{ID: 1}},
		{4, 8, &emp.Employee{ID: 2}},
	}
	assertTestScenario(t, table, o)
}

//unequally distributed organization
func TestScenario6(t *testing.T) {
	fmt.Println("Test Scenario: 6")

	heirarchy := make(map[int][]int)
	heirarchy[1] = []int{2, 3, 4, 5}
	heirarchy[2] = []int{6, 7}
	heirarchy[3] = []int{8, 9, 10}
	heirarchy[4] = []int{11}

	o := org.Organization{}
	buildTestScenario(&o, 11, heirarchy)

	table := []testRow{
		{0, 9, nil},
		{1, 16, nil},
		{9, 9, nil},
		{6, 7, &emp.Employee{ID: 2}},
		{8, 9, &emp.Employee{ID: 3}},
		{6, 10, &emp.Employee{ID: 1}},
		{6, 11, &emp.Employee{ID: 1}},
		{6, 2, &emp.Employee{ID: 2}},
		{11, 4, &emp.Employee{ID: 4}},
		{9, 1, &emp.Employee{ID: 1}},
	}
	assertTestScenario(t, table, o)
}
