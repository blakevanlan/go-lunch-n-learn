go-lunch-n-learn
================

Go Lunch 'n Learn example code 

## Hello World

### Packages
* All Go programs are made up of packages
* Programs start running in package _main_
* _import_ is used to import other packages

### Exported Names
* After importing you can refer to the name a package exports
* Go exports anything that begins with a capital letter

### Variables
* Declare with either _var_ or _:=_ (:= can only be used inside functions)
* Type is declared after variable name
* If initializer is present, type can be ommitted
* _:=_ uses implicity typing
* Constants can be a character, string, boolean, or a numeric value

### Functions
* Return value type always comes after parameter list
* When consecutive parameters share a type it can be shortened
* Functions can return any number of results
* Go functions can return _result parameters_ that act as normal variables while inside the function and return their current value on _return_

## Loops and Conditions

### If Statements
* No _( )_ and _{ }_ are required
* Can start with short statement before the execution of condition
* Variables declared in the statement are only in scope until the end of the if block

### Switch Statements
* Evaluates from top to bottom
* Automatically breaks unless a case ends with a _fallthrough_ statement
* Switch statement without condition can be used for long if-then-else chains

## Data Structures

### Pointers
* No pointer arithmetic

### Structs
* A collection of fields
* Uses the dot accessor
* _new(T)_ function allocates a zeroed struct and returns a pointer to it

### Slices
* Points to an array of values
* _len(s)_ is used to get the length of a slice
* Can be re-sliced, creating a new slice value that points to the same array
   * expression: _s[lo:lo]_
   * evaluates from _lo_ through _hi-1_, inclusive
   * _s[lo:lo]_ is empty and _s[lo:lo+1]_ has one element
* Slices are created with the _make_ function
* _cap(s)_ is used to get the capacity of a slice

### Maps
* Key value pair structure