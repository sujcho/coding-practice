#!/usr.bin/env python
"""Write a program that reads in two integers and determines wheter the first is a multiple of the seconds."""

while True: # True/False are keywords.
    try:
        first = int(raw_input("Enter the first number: "))
        break
    except ValueError: #built in exception
        print "Please try again."

while True: # True/False are keywords.
    try:
        second = int(raw_input("Enter the second number: "))
        break
    except ValueError: #built in exception
        print "Please try again."

if first % second == 0:
        print '%d is a multiple of %d' % (first, second)
else:
        print '%d is a not multiple of %d' % (first, second)
