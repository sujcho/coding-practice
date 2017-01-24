#!/usr/bin/env python
""" Write a script that uses nested while loops to produce this pattern:

      *
    * * *
  * * * * *
* * * * * * *
"""

number = 4
num_of_stars = 1;
while number > 0:
    stars = 0;
    space = 0;
    while space < number:
        print " ",
        space += 1
    while stars < num_of_stars:
        print "*",
        stars += 1
    print
    num_of_stars += 2
    number -= 1

#simple way


while number > 0:
    print number * ' ' + num_of_stars * '*'
    num_of_stars += 2
    number -= 1
