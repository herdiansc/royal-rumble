# Royal Rumble
This is an app to sort list of royal names consist of ordinal number in roman numeral.
## Problem
An ordinal number is a word representing rank or sequential order. The naming convention for royal names is to follow a given name with an ordinal number using a Roman numeral to indicate the birth order of two people of the same name.
The Roman numerals from 1 to 50 are defined as follows: The numbers 1 through 10 are written I, II, III, IV, V, VI, VII, VIII, IX, and X. The Roman numerals corresponding to the numbers 20, 30, 40, and 50 are XX, XXX, XL, and L. For any other two-digit number < 50, its Roman numeral representation is constructed by concatenating the numeral(s) for its multiples of ten with the numeral(s) for its values < 10. For example, 47 is 40 + 7 = "XL" + "VII" = "XLVII".
In this challenge, you will be given a list of royal name strings consisting of a given name followed by an ordinal number. You must sort the list first alphabetically by name, then by ordinal increasing within any given name.
For example, if you are given the royal names [George VI, William II, Elizabeth I, William I] the result of the sort is [Elizabeth I, George VI, William I, William II].

## Requirement
- Golang: Please you can use this url https://golang.org/doc/install to install go on your machine.

## How to install
- Clone this repo to your $GOPATH/src directory

## How to Run
I have built this app and included the executable version in this repo. Here is step to run the file:

- Enter app root directory
- Execute command `./royal-rumble <input_filename>` for example `./royal-rumble RoyalRumble/input1.txt`

## How to Build
- Execute command `go build .`

## How to Run Unit Test
- Enter app root directory
- Execute command `go test ./...`

