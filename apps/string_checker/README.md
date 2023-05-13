# Task: "Custom error"

Write a `MyCheck(input string) error` function that should check the string against the following parameters:
* The string must not contain digits `(found digits)`.
* The line must be less than 20 characters long `(line is too long)`.
* The line must have two spaces `(no two spaces)`.

The texts of the returned errors are given in brackets.

The function must find all possible errors in a single call.
The result should contain a slice of errors for which the line failed the test, or be `nil`.