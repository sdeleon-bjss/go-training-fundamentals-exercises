package pkg

// Note - if we want to use Godoc examples, this file needs to use the _test suffix to be recognized as a test file
// this allows the `godoc` tool to recognize the examples in this file
// and properly associate + render them proper in godoc for the localhost:port/pkg/package
// keeping it as _test, requires all references to be updated back to using: pkg.
// - caution this will output errors in the terminal when running the tests (since it will be of _test again)

//import (
//	"fmt"
//	"time"
//)
//
//func ExampleConcatStrings() {
//	s1 := "Hello"
//	s2 := " "
//	s3 := "World"
//
//	result := ConcatStrings(s1, s2, s3)
//	fmt.Println(result)
//	// Output: Hello World
//}
//
//func ExampleGetFullName() {
//	fullName := GetFullName("Steven", "William", "DeLeon")
//	println(fullName)
//	// Output: Steven William DeLeon
//}
//
//func ExampleIsNumberBetweenRange_success() {
//	result := IsNumberBetweenRange(5, 1, 10)
//	println(result)
//	// Output: true
//}
//
//func ExampleIsNumberBetweenRange_failure() {
//	result := IsNumberBetweenRange(100, 1, 10)
//	println(result)
//	// Output: false
//}
//
//func ExampleAscending() {
//	n := []int{5, 3, 1, 4, 2}
//	result := Ascending(n)
//	fmt.Println(result)
//	// Output: [1 2 3 4 5]
//}
//
//func ExampleDescending() {
//	n := []int{5, 3, 1, 4, 2}
//	result := Descending(n)
//	fmt.Println(result)
//	// Output: [5 4 3 2 1]
//}
//
//func ExampleOddEvenCount() {
//	n := []int{5, 3, 1, 4, 2}
//	odd, even := OddEvenCount(n, "ascending")
//	fmt.Println(odd)
//	fmt.Println(even)
//	// Output: -9
//	// -6
//}
//
//func ExampleNumArray_SingleDigits() {
//	n := [9]int{1, 2, 3, 40, 50, 60, 700, 800, 900}
//	numArray := NumArray{
//		Numbers: n,
//	}
//
//	result := numArray.SingleDigits()
//	fmt.Println(result)
//	// Output: [1 2 3]
//}
//
//func ExampleNumArray_DoubleDigits() {
//	n := [9]int{1, 2, 3, 40, 50, 60, 700, 800, 900}
//	numArray := NumArray{
//		Numbers: n,
//	}
//
//	result := numArray.DoubleDigits()
//	fmt.Println(result)
//	// Output: [40 50 60]
//}
//
//func ExampleNumArray_TripleDigits() {
//	n := [9]int{1, 2, 3, 40, 50, 60, 700, 800, 900}
//	numArray := NumArray{
//		Numbers: n,
//	}
//
//	result := numArray.TripleDigits()
//	fmt.Println(result)
//	// Output: [700 800 900]
//}
//
//func ExampleNumArray_Sum() {
//	n := [9]int{1, 2, 3, 40, 50, 60, 700, 800, 900}
//	numArray := NumArray{
//		Numbers: n,
//	}
//
//	singleDigits := numArray.SingleDigits()
//
//	result := numArray.Sum(singleDigits)
//	fmt.Println(result)
//	// Output: 6
//}
//
//func ExampleNumArray_FinalSum() {
//	n := [9]int{1, 2, 3, 40, 50, 60, 700, 800, 900}
//	numArray := NumArray{
//		Numbers: n,
//	}
//
//	result := numArray.FinalSum()
//	fmt.Println(result)
//	// Output: 2556
//}
//
//func ExamplePerson_Age() {
//	p := Person{
//		DOB: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
//	}
//
//	result := p.Age()
//	fmt.Println(result)
//	// Output: 4
//}
//
//func ExampleRollDice() {
//	result := RollDice()
//	fmt.Println(result)
//	// Output: 2  // or any number between 1 and 6
//}
//
//func ExampleComboRolls() {
//	result := ComboRolls()
//	fmt.Println(result)
//	// Output: [2 4]
//}
//
//func ExampleProcessRoll_naturalSeven() {
//	result := ProcessRoll(7)
//	fmt.Println(result)
//	// Output: NATURAL
//}
//
//func ExampleProcessRoll_naturalEleven() {
//	result := ProcessRoll(11)
//	fmt.Println(result)
//	// Output: NATURAL
//}
//
//func ExampleProcessRoll_snakeEyesCraps() {
//	result := ProcessRoll(2)
//	fmt.Println(result)
//	// Output: SNAKE-EYES-CRAPS
//}
//
//func ExampleProcessRoll_lossCrapsThree() {
//	result := ProcessRoll(3)
//	fmt.Println(result)
//	// Output: LOSS-CRAPS
//}
//
//func ExampleProcessRoll_lossCrapsTwelve() {
//	result := ProcessRoll(12)
//	fmt.Println(result)
//	// Output: LOSS-CRAPS
//}
//
//func ExampleProcessRoll_neutralDefaultCase() {
//	result := ProcessRoll(4)
//	fmt.Println(result)
//	// Output: NEUTRAL
//}
//
//func ExampleWriteStringsToFile() {
//	s := []string{"Hello", "World"}
//	err := WriteStringsToFile(s, "example.txt")
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	// Output: should create a file named example.txt with "Hello\nWorld\n" at the root of the project
//}
//
//func ExampleReadStringsFromFile() {
//	// given that the file example.txt contains "Hello\nWorld\n"
//	s, err := ReadStringsFromFile("example.txt")
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	fmt.Println(s)
//	// Output: [Hello World]
//}
//
//func ExampleName_SetFullName() {
//	fullName := "Steven William DeLeon"
//
//	name := Name{}
//	_ = name.SetFullName(fullName)
//
//	fmt.Println(name.GetFullName())
//	// Output: Steven William DeLeon
//}
//
//func ExampleName_GetFullName() {
//	name := Name{
//		First:  "Steven",
//		Middle: "William",
//		Last:   "DeLeon",
//	}
//
//	got := name.GetFullName()
//	fmt.Println(got)
//	// Output: Steven William DeLeon
//}
//
//func ExampleName_Print() {
//	name := Name{
//		First:  "Steven",
//		Middle: "William",
//		Last:   "DeLeon",
//	}
//
//	name.Print()
//	// Output: full-name : Steven William DeLeon, middle-name : William, and surname : DeLeon
//}
//
//func ExampleNewPupil() {
//	p, err := NewPupil("Steven", "William", "DeLeon", "01/28/1992")
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	fmt.Println(p)
//	// Output: {{Steven William DeLeon} 1992-01-01 00:00:00 +0000 UTC}
//}
//
//func ExamplePupil_GetFullName() {
//	p, _ := NewPupil("Steven", "William", "DeLeon", "01/28/1992")
//
//	output := p.GetFullName()
//	fmt.Println(output)
//	// Output: Steven William DeLeon
//}
//
//func ExamplePupil_GetDateOfBirth() {
//	p, _ := NewPupil("Steven", "William", "DeLeon", "01/28/1992")
//
//	output := p.GetDateOfBirth()
//	fmt.Println(output)
//	// Output: 01/28/1992
//}
//
//func ExamplePupil_GetAge() {
//	p, _ := NewPupil("Steven", "William", "DeLeon", "01/28/1992")
//
//	output := p.GetAge()
//	fmt.Println(output)
//	// Output: 32
//}
