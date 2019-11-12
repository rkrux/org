Notes
- No external dependencies required
- 40 Test Cases covered among 6 Test Scenarios
- An Employee is represented by a struct containing ID and a slice of Employee pointers
- An Organization is represented as a N-Ary Tree by a struct containing TotalEmployees and a Founder  
- Organization heirarchy and Closest Common Manager (CCM) for every test case is displayed in console
- To add test cases, update scenarios_test.go file
- Organization's heirarchy is defined by a map that is passed onto Organization.Create() function

Assumptions
- ID of Organization's founder is always 1
- An employee is a manager if he has more than 1 Direct under him - (len(Employee.Directs) > 0)
- All relations among the employees roll up to the Founder - the heirarchy is connected for all the employees
- CCM of a manager and his employee is the manager himself

CCM negative scenarios
- Same employee passed in both the arguments
- Either of the employee absent from the organization

4 Files:
- main.go
- employee.go
- organization.go
- scenarios_test.go

Build: go install github.com/rkrux/org
Run from bin directory: ./org
Run from /org/ directory: go run main.go
Test from /org/testscenarios/ directory: go test