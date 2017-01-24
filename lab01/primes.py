#!/usr/bin/env python
"""primes/py -- Produces a list of prime numbers.
Here, we are only cheching the "look" of the Python code.
"""
MAX = 100  #Here is a comment.

print 'primes are :' # A new line is added by default.
number = 3
while number < MAX:
    div = 2
    while div * div <= number:
        if number % div == 0:
            break
        div += 1
    else: #Overloaded 'else', loop didn't break'.
        print number, #Trailing Comma suppresses the new line.
    number += 2
print #This only produces the new line.
