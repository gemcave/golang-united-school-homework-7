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
func TestLen(t *testing.T) {
	as := assert.New(t)
	person01 := Person {"Lisa", "Joy", time.Date(1984, time.April,14,0,0,0,0, time.UTC)}
	person02 := Person {"Chris", "Nolan", time.Date(1978,time.July,25,0,0,0,0, time.UTC)}
	person03 := Person {"Anthony", "Hop", time.Date(1967,time.January,9,0,0,0,0, time.UTC)}
	peopleList := People{person01, person02, person03}
	as.Equal(peopleList.Len(), 3)
	
}

func TestLessPeopleFirstName(t *testing.T) {
	personA := Person{firstName: "John", lastName: "Smith",
		birthDay: time.Date(2002, time.March, 1, 0, 0, 0, 0, time.UTC),
	}
	personB := Person{firstName: "Jane", lastName: "Smith",
		birthDay: time.Date(2002, time.March, 1, 0, 0, 0, 0, time.UTC),
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

func TestRows(t *testing.T) {
	as := assert.New(t)
	mx01 := Matrix{1,1,[]int{0}}
	mx02 := Matrix{2,1,[]int{13,55}}
	mx03 := Matrix{2,2,[]int{3, 9, 12, 15}}
	tMatrix := []struct{
		matrix Matrix
		expected [][]int
	}{
		{mx01, [][]int{[]int{0}}},
		{mx02, [][]int{[]int{13}, []int{55}}},
		{mx03, [][]int{[]int{3, 9}, []int{12, 15}}},
	}
	for _, value := range tMatrix{
		actual := value.matrix.Rows()
		as.EqualValues(value.expected, actual)
	}
}

func TestCols(t *testing.T) {
	as := assert.New(t)
	mx01 := Matrix{1,1,[]int{0}}
	mx02 := Matrix{2,1,[]int{13,55}}
	mx03 := Matrix{2,2,[]int{3, 9, 12, 15}}
	tMatrix := []struct{
		matrix Matrix
		expected [][]int
	}{
		{mx01, [][]int{[]int{0}}},
		{mx02, [][]int{[]int{13, 55}}},
		{mx03, [][]int{[]int{3, 12}, []int{9, 15}}},
	}
	for _, value := range tMatrix{
		actual := value.matrix.Cols()
		as.EqualValues(value.expected, actual)
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