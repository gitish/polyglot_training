
# Assignment 1: Temperature Converter
## Description: 
Create a Go program that converts temperatures between Celsius and Fahrenheit.

* Requirements:
    - Write functions ```celsiusToFahrenheit(c float64) float64``` and ```fahrenheitToCelsius(f float64) float64```.
    - Take user input to decide which conversion to perform and to input the temperature value.
    - Print the converted temperature.
* Concepts Covered: Functions, user input, basic data types, conditional statements.

# Assignment 2: Simple Calculator
## Description: 
Create a simple command-line calculator that can add, subtract, multiply, and divide two numbers.

* Requirements:
    - Write functions for each operation: ```add```, ```subtract```, ```multiply```, and ```divide```.
    - Use a switch-case or ```if-else``` statement to determine the operation based on user input.
    - Handle division by zero in the ```divide``` function by returning an error message.
* Concepts Covered: Functions, control structures, error handling, user input.

# Assignment 3: Prime Number Checker
## Description: 
Write a program that checks if a given number is a prime number.

* Requirements:
    - Write a function ```isPrime(n int) bool``` that returns ```true``` if the number is prime and ```false``` otherwise.
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
    - Define a ```Student``` struct with fields for ```Name```, ```Age```, and ```Grades``` (a slice of integers).
    - Add a method ```averageGrade()``` that calculates the average grade for the student.
    - Prompt the user to enter details for multiple students and display each student’s average grade.
* Concepts Covered: Structs, methods, slices, basic arithmetic.

# Assignment 6: Sorting Names Alphabetically
## Description: 
Write a program that sorts a list of names alphabetically.

* Requirements:
    - Take a list of names from the user as input.
    - Sort the names in alphabetical order using Go’s built-in ```sort``` package.
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
Set up a simple HTTP server that responds with "Hello, Go!" on the ```/hello``` endpoint.

* Requirements:
    - Use the ```net/http``` package to create an HTTP server.
    - Set up a handler function that responds with a greeting when accessing /hello.
    - Run the server locally on port 8080 and test it by visiting ```http://localhost:8080/hello```.
* Concepts Covered: Basic web server, HTTP handlers, ```net/http``` package.

# Assignment 9: File I/O – Reading and Writing Files
## Description: 
Write a program that reads a text file, counts the number of lines, and writes the output to a new file.

* Requirements:
    - Prompt the user for the name of an input file.
    - Count the lines in the file and write the line count to a new file called ```output.txt```.
    - If the file doesn’t exist, display an error message.
* Concepts Covered: File I/O, error handling, basic Go file operations.

# Assignment 10: Shopping Cart (Using Structs and Slices)
## Description: 
Implement a basic shopping cart where users can add, view, and remove items.

* Requirements:
    - Define a ```Product``` struct with fields like ```Name```, ```Price```, and ```Quantity```.
    - Allow the user to add products to a shopping cart, view the cart, and remove items.
    - Display the total price of items in the cart when requested.
* Concepts Covered: Structs, slices, functions, basic input handling.

# Assignment 11: Concurrent Sum Calculator
## Description: 
Write a program that calculates the sum of all integers in a slice concurrently by dividing the slice into two halves and summing each half in separate goroutines.

* Requirements:
    - Split the slice of integers into two halves.
    - Use two goroutines to calculate the sum of each half concurrently.
    - Use channels to retrieve the partial sums and calculate the total sum in the main function.
    - Print the total sum.
* Concepts Covered: Goroutines, channels, concurrent calculations.

# Assignment 12: Prime Number Checker with Goroutines
## Description: 
Create a program that checks if multiple numbers are prime concurrently.

* Requirements:
    - Take a list of integers as input.
    - Use a goroutine to check if each number is prime.
    - Use a channel to send the result (whether each number is prime) back to the main function.
    - Print the results after all goroutines finish.
* Concepts Covered: Goroutines, channels, synchronization.

# Assignment 13: Worker Pool Implementation
## Description: 
Implement a worker pool that processes a set of tasks concurrently.

* Requirements:
    - Define a task as a function that takes an integer and performs a time-consuming operation (e.g., calculating factorial).
    - Use multiple worker goroutines to process tasks concurrently.
    - Use a channel to send tasks to the worker pool and another channel to collect results.
    - Limit the number of concurrent workers (e.g., 3 workers).
* Concepts Covered: Goroutines, channels, worker pool pattern, channel buffering.

# Assignment 14: Concurrent Web Scraper
## Description: 
Create a basic web scraper that fetches data from multiple URLs concurrently.

* Requirements:
    - Define a list of URLs.
    - Use goroutines to fetch data from each URL concurrently.
    - Use channels to send the fetched data back to the main function.
    - Print the response length or status for each URL after all goroutines complete.
* Concepts Covered: Goroutines, channels, concurrency control, network operations.

# Assignment 15: Ping-Pong Simulation with Channels
## Description: 
Simulate a ping-pong game where two goroutines "ping" and "pong" each other using a channel.

* Requirements:
    - Create two goroutines, ping and pong.
    - Use a channel to pass a message back and forth between the goroutines.
    - Stop the game after a certain number of "ping-pong" exchanges (e.g., 10 rounds).
    - Print each "ping" and "pong" to the console.
* Concepts Covered: Goroutines, channels, back-and-forth message passing.

# Assignment 16: Race Condition Detector
## Description: 
Write a program that demonstrates a race condition and then uses channels to avoid it.

* Requirements:
    - Create a shared variable (e.g., a counter).
    - Start multiple goroutines that increment this counter without using channels to see the race condition in action.
    - Use channels to control access to the shared variable and avoid the race condition.
    - Print the counter's final value to verify correct results.
* Concepts Covered: Race conditions, goroutines, channels for synchronization.

# Assignment 17: Fibonacci Sequence Generator with Channels
## Description: 
Implement a program that generates Fibonacci numbers concurrently.

* Requirements:
    - Use a goroutine to generate Fibonacci numbers and send them to a channel.
    - In the main function, receive numbers from the channel and print the first N Fibonacci numbers (where N is provided by the user).
    - Ensure the program exits cleanly after printing N numbers.
* Concepts Covered: Goroutines, channels, concurrency with sequential data generation.

# Assignment 18: Producer-Consumer Pattern
## Description: 
Implement a producer-consumer model using channels where one goroutine produces random numbers, and another consumes them.

* Requirements:
    - Create a producer goroutine that generates random numbers and sends them to a channel.
    - Create a consumer goroutine that receives numbers from the channel and prints them.
    - Run the program for a fixed number of items or seconds, and ensure it exits cleanly.
* Concepts Covered: Goroutines, channels, producer-consumer pattern.

# Assignment 19: Parallel Matrix Multiplication
## Description: 
Write a program that multiplies two matrices concurrently by assigning each row’s computation to a separate goroutine.

* Requirements:
    - Define two matrices and multiply them.
    - Use goroutines to calculate the result of each row in parallel.
    - Use channels to collect the results from each goroutine.
    - Print the resulting matrix.
* Concepts Covered: Goroutines, channels, concurrent data processing.

# Assignment 20: Timeout with Channels
## Description: 
Implement a program that performs a time-consuming operation but stops if it takes too long.

* Requirements:
    - Create a function that simulates a time-consuming operation (e.g., sleep).
    - Use a channel to signal the operation’s completion.
    - Use select to implement a timeout so that if the operation takes too long, a timeout message is printed instead.
    - Print whether the operation finished successfully or timed out.
* Concepts Covered: Channels, select statement, timeouts, concurrency control.


