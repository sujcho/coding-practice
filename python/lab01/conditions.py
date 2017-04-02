#!/usr/bin/env python
"""concions.py demeonstrates if/elif/els and while/else."""

number = 4

# if/elit/else

if number < 10:
    print number, 'is small.'
elif number >= 1000:
    print number, 'is big.'
else: print number, 'is medium'

#Alternate syntax for 2.5 -- all one line but less readable.

print number, "is",

print "small." if number < 10 \
               else "big." if number >= 1000\
               else "medium."

# else occurs in a loop too

while number < 6:
    if number % 3 == 0:
        print number, 'is divisible by 3.'
        break
    number += 1
else:
    print 'Nothing in the loop was divisible by 3.'
