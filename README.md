Go Go?
================

Go Lunch 'n Learn example code 

## Hello World

### Packages
* All Go programs are made up of packages
* Programs start running in package ```main```
* ```import``` is used to import other packages

### Exported Names
* After importing you can refer to the name a package exports
* Go exports anything that begins with a capital letter

### Variables
* Declare with either ```var``` or ```:=``` (```:=``` can only be used inside functions)
* Type is declared after variable name
* If initializer is present, type can be ommitted
* ```:=``` uses implicity typing
* Constants can be a character, string, boolean, or a numeric value

### Functions
* Return value type always comes after parameter list
* When consecutive parameters share a type it can be shortened
* Functions can return any number of results
* Go functions can return **result parameters** that act as normal variables while inside the function and return their current value on ```return```

## Loops and Conditions

### If Statements
* No **( )** and **{ }** are required
* Can start with short statement before the execution of condition
* Variables declared in the statement are only in scope until the end of the if block

### For loop
* Go only has a `for` loop
* As with C, the pre and post statements can be empty and then no `;` are required
* For an infinite loop the keyword `for` can just be used

### Switch Statements
* Evaluates from top to bottom
* Automatically breaks unless a case ends with a ```fallthrough``` statement
* Switch statement without condition can be used for long if-then-else chains

## Data Structures

### Pointers
* No pointer arithmetic

### Structs
* A collection of fields
* Uses the dot accessor
* ```new(T)``` function allocates a zeroed struct and returns a pointer to it

### Slices
* Points to an array of values
* ```len(s)``` is used to get the length of a slice
* Can be re-sliced, creating a new slice value that points to the same array
   * expression: ```s[lo:lo]```
   * evaluates from ```lo``` through ```hi-1```, inclusive
   * ```s[lo:lo]``` is empty and ```s[lo:lo+1]``` has one element
* Slices are created with the ```make()``` function
* ```cap(s)``` is used to get the capacity of a slice

### Maps
* Key value pair structure
* Can be created with ```make()```
* ```delete()``` is used to remove an object from a map
* A loopup returns two parameters, the value (if any) and a boolean indicating if the element exists in the map

## Closures
* Work just like closures in JavaScript

## Methods and Interfaces

### Methods
* Go does not have classes
* Methods can be attached to structs
* The *method receiver* appears between ```func``` and the method name
* The method receiver should be a pointer to the type, this allows the method to modify properties of the type

### Interfaces
* Interface type is defined as a set of methods
* A value of interface type can hold any value that implements those methods
* Allows for increased decoupling

## Concurrent Programming

### Goroutines
* Lightweight thread, managed by Go runtime
* ```go``` is the keyword used to start a new goroutine

### Channels
* Typed conduit through which you can send and receive values with the channel operator, ```<-```
* Must be created before use
* By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables
* Can be buffered, provide buffer amount as the second argument to ```make```
* Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.

### Range and Close
* A sender can close a channel to indicate that no more values with be sent
* A receiver can check if a channel is open by utilizing the second return value
* A range loop can be used to receive values until the channel is closed

### Select blocks
* A select block blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
* ```default``` can be used to create a non-blocking pattern

