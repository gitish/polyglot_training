
# Assignment 1: Temperature Converter
## Description: 
Create a Go program that converts temperatures between Celsius and Fahrenheit.

* Requirements:
    - Write functions celsiusToFahrenheit(c float64) float64 and fahrenheitToCelsius(f float64) float64.
    - Take user input to decide which conversion to perform and to input the temperature value.
    - Print the converted temperature.
* Concepts Covered: Functions, user input, basic data types, conditional statements.

# Assignment 2: Simple Calculator
## Description: 
Create a simple command-line calculator that can add, subtract, multiply, and divide two numbers.

* Requirements:
    - Write functions for each operation: add, subtract, multiply, and divide.
    - Use a switch-case or if-else statement to determine the operation based on user input.
    - Handle division by zero in the divide function by returning an error message.
* Concepts Covered: Functions, control structures, error handling, user input.

# Assignment 3: Prime Number Checker
## Description: 
Write a program that checks if a given number is a prime number.

* Requirements:
    - Write a function isPrime(n int) bool that returns true if the number is prime and false otherwise.
    - Prompt the user to enter a number and display whether it is prime.
    - Implement a loop to allow users to check multiple numbers until they choose to exit.
* Concepts Covered: Functions, loops, conditional logic.

# Assignment 4: Word Frequency Counter
## Description: 
Create a program that counts the frequency of each word in a sentence.

* Requirements:
    - Take a sentence as input from the user.
    - Split the sentence into words and store the words in a map where keys are words, and values are their frequencies.
    - Print each word along with its frequency.
* Concepts Covered: Strings, maps, loops.

# Assignment 5: Structs and Methods – Student Grades
## Description: 
Implement a program that stores information about students and their grades.

* Requirements:
    - Define a Student struct with fields for Name, Age, and Grades (a slice of integers).
    - Add a method averageGrade() that calculates the average grade for the student.
    - Prompt the user to enter details for multiple students and display each student’s average grade.
* Concepts Covered: Structs, methods, slices, basic arithmetic.

# Assignment 6: Sorting Names Alphabetically
## Description: 
Write a program that sorts a list of names alphabetically.

* Requirements:
    - Take a list of names from the user as input.
    - Sort the names in alphabetical order using Go’s built-in sort package.
    - Display the sorted list.
* Concepts Covered: Slices, sorting, strings, standard library usage.

# Assignment 7: Countdown Timer using Goroutines
## Description: 
Create a simple countdown timer that displays a countdown from a given number of seconds.

* Requirements:
    - Prompt the user to enter the number of seconds.
    - Use a goroutine to print the remaining time every second.
    - When the countdown reaches zero, print "Time's up!".
* Concepts Covered: Concurrency, goroutines, time package.

# Assignment 8: Basic HTTP Server
## Description: 
Set up a simple HTTP server that responds with "Hello, Go!" on the /hello endpoint.

* Requirements:
    - Use the net/http package to create an HTTP server.
    - Set up a handler function that responds with a greeting when accessing /hello.
    - Run the server locally on port 8080 and test it by visiting http://localhost:8080/hello.
* Concepts Covered: Basic web server, HTTP handlers, net/http package.

# Assignment 9: File I/O – Reading and Writing Files
## Description: 
Write a program that reads a text file, counts the number of lines, and writes the output to a new file.

* Requirements:
    - Prompt the user for the name of an input file.
    - Count the lines in the file and write the line count to a new file called output.txt.
    - If the file doesn’t exist, display an error message.
* Concepts Covered: File I/O, error handling, basic Go file operations.

# Assignment 10: Shopping Cart (Using Structs and Slices)
## Description: 
Implement a basic shopping cart where users can add, view, and remove items.

* Requirements:
    - Define a Product struct with fields like Name, Price, and Quantity.
    - Allow the user to add products to a shopping cart, view the cart, and remove items.
    - Display the total price of items in the cart when requested.
* Concepts Covered: Structs, slices, functions, basic input handling.
