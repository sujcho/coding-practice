#!/usr/bin/env python
# -*- coding: utf-8 -*
"""Palindromize (phrase) that returns the lowercase version of the phrase"""
import string

def Palindromize (phrase):
    """Palindromize (phrase) returns the lowercase version of the phrase with whitespace and punctuation removed if the phrase is a palindrome. If not, it returns  None. """

    #remove punctuation
    phrase = phrase.strip(string.punctuation)
    #remove whitespace
    phrase = "".join(phrase.split(" "))

    length = len(phrase)

    #If only one character, it's not palindrome
    if length == 1 or not phrase :
        return None

    #get first half of the phrase
    first_half = phrase[:length/2].lower()
    #get second half of the phrase

    if length % 2 == 1:
        second_half = phrase[(length/2)+1:].lower()
    else:
        second_half = phrase[length/2:].lower()

    #Is this Plaindrom?
    #reverse the second half and compare with first half
    if first_half == second_half[::-1]:
        return phrase
    else:
        return None


if __name__ == "__main__":
    DATA = ('Murder for a jar of red rum', '12321', 'nope', 'abcbA', '3443', 'what','Never odd or even', 'Rats live on no evil star')
    for data in DATA:
        print "%s: %s" % (data, Palindromize(data))
