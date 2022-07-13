package coverage

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

func TestLenPeople(t *testing.T) {
	personA := Person{firstName: "John", lastName: "Smith",
		birthDay: time.Date(1997, time.March, 1, 0, 0, 0, 0, time.UTC),
	}
	personB := Person{firstName: "Jane", lastName: "Smith",
		birthDay: time.Date(1995, time.July, 10, 10, 10, 0, 0, time.UTC),
	}
	expectedResult := 2
	listPeople := People{personA, personB}

	actualResult := listPeople.Len()
	if actualResult != expectedResult {
		t.Errorf("Expeсted Len: %d, Received Len: %d", expectedResult, actualResult)
	}
}

func TestLessPeopleFirstName(t *testing.T) {
	personA := Person{firstName: "John", lastName: "Smith",
		birthDay: time.Date(1998, time.March, 1, 0, 0, 0, 0, time.UTC),
	}
	personB := Person{firstName: "Jane", lastName: "Smith",
		birthDay: time.Date(1996, time.March, 1, 0, 0, 0, 0, time.UTC),
	}
	peopleList := People{personA, personB}
	expectedResult := true
	actualResult := peopleList.Less(0, 1)
	if actualResult != expectedResult {
		t.Errorf("People Less Function: Expeсted Result: %t, Received Result: %t", expectedResult, actualResult)
	}
}


func TestSwap(t *testing.T) {
	birthdayA := time.Date(1997, time.May, 10, 0, 0, 0, 0, time.UTC)
	birthdayB := time.Date(1997, time.May, 5, 0, 0, 0, 0, time.UTC)
	personA := Person{firstName: "John", lastName: "Smith", birthDay: birthdayA}
	personB := Person{firstName: "Jane", lastName: "Smith", birthDay: birthdayB}
	peopleList := People{personA, personB}

	peopleList.Swap(0, 1)
	if peopleList[0].firstName != "Jane" || peopleList[0].lastName != "Smith" || peopleList[0].birthDay != birthdayB ||
		peopleList[1].firstName != "John" || peopleList[1].lastName != "Smith" || peopleList[1].birthDay != birthdayA {
		t.Errorf("People Swap Function: Expeсted Result is different from Received Result")
	}
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////
func TestNew(t *testing.T){
	as := assert.New(t)
	tString := []struct{
		str string
	}{
		{"sequoia"},
		{"1\n2 3"},
		{"3"},
	}
	for _, value := range tString{
		actual, err := New(value.str)
		if err != nil{
			as.Error(err)
			as.Nil(actual)
		} else{
			as.NoError(err)
			as.NotNil(actual)
		}
	}
}

func TestSet(t *testing.T) {
	as := assert.New(t)
	mx01 := Matrix{2,2,[]int{33, 55, 8, 6}}
	tSetMatrix := []struct{
		matrix Matrix
		rov, col, value int
		ecpected bool
	}{
		{mx01, 1, 2, 1, false},
		{mx01, 0, 0, 0, true},
		{mx01, 1, 1, 3, true},
	}
	for _, value := range tSetMatrix{
		actual := mx01.Set(value.rov, value.col, value.value)
		as.Equal(value.ecpected, actual)
	}

	}

	func TestRowsMatrix(t *testing.T) {
		m, err := New("10 20\n30 40")
		if err != nil {
			t.Errorf("Failed to create a new Matrix")
		}
		expected := [][]int{{10, 20}, {30, 40}}
		result := m.Rows()
		for i, value := range result {
			for j, valCol := range value {
				if valCol != expected[i][j] {
					t.Errorf("Expected Result %d, Received Result %d", expected[i][j], valCol)
				}
			}
		}
	}
	
	func TestColsMatrix(t *testing.T) {
		m, err := New("10 20\n30 40")
		if err != nil {
			t.Errorf("Failed to create a new Matrix")
		}
		expected := [][]int{{10, 30}, {20, 40}}
		result := m.Cols()
		for i, value := range result {
			for j, valRow := range value {
				if valRow != expected[i][j] {
					t.Errorf("Expected Result %d, Received Result %d", expected[i][j], valRow)
				}
			}
		}
	}