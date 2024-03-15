# BJSS Go Academy - Exercises assessment

## Table of Contents
1. [Introduction](#introduction)
2. [Resources](#resources)
3. [Exercises](#exercises)
4. [How to run and test](#how-to-run-and-test)

### Resources
- [Go Academy](https://bjss.learnamp.com/en/learnlists/golang-academy)
- [Go Assignments (fundamentals)](https://bjss.learnamp.com/en/learnlists/golang-academy/items/go-assignments)

### Exercises
1. Create a program that has multiple string variable and displays the string on one line. [Strings]
2. Create a program that lets the user input a first name, middle name and last name. Display the person's full name on one line. [Keyboard input]
3. Create a program that allows the user to input a number. Check whether the number lies between 1 and 10. [Variables]
4. Create a program that initialises an array with the integer values 1 to 10: [Arrays][Slices][For Loops]
   - Display the array content in ascending sequential order 1 to 10.
   - Display the array content in descending sequential order 10 to 1. 
   - Count even numbers and odd numbers in increasing and decreasing sequential order. 
   - Display the even and odd count sequences to screen.
5. Create a program that accepts and sums nine numbers. [Methods][Arrays][Slices][For loops]
   - Three single digit numbers from one method.
   - Three double digit numbers from a second method.
   - Three triple digit numbers from a third method.
   - Finally sum all methods into a final sum in the main program.
6. Create a program that calculates the age of a person given their date of birth. [Variables][Methods][Arrays][Slices][For Loops][Package Usage]
   - (Use the github.com/bearbin/go-age to aid in the creation of this app. Also review unit testing applied to the application age.go within the imported package.)
7. Create a program that rolls two dice (1 to 6) fifty times. Display the number rolls and the outcomes in sequential order.Resulting rolls are to be processed in the following manner: [Random Numbers][Switches]
   - 7 and 11 are to be called NATURAL
   - 2 is called SNAKE-EYES-CRAPS
   - 3 and 12 is called LOSS-CRAPS
   - Any other combination is called NEUTRAL.
8. Create a program that: [Write File][Read File][I/O Package][I/O]
   - Copies the following list of cities to a new file - "Abu Dhabi", "London", "Washington D.C.", "Montevideo", "Vatican City", "Caracas", "Hanoi".
   - Reads a list of cities from the newly created file.
   - Displays the list of cities in alphabetical order.
9. Extend the program in Exercise 2 by slicing the full name into 3 slices. Display the full-name : <full-name>, middle-name : <middle-name> and surname : <surname> on 3 separate lines. [Slices] [Structures]
10. Create a school register program that lists 10 pupils - full name, date of birth and age. [Structures][Arrays][Interfaces]


## How to run and test
> all commands can be run from root as follows
### Run
- using the `main.go` file in each exercise
```shell
go run ./cmd/exercise_01/main.go
# ...
go run ./cmd/exercise_10/main.go
```

### Test
```shell
go test -v ./pkg/...
```