"""Write a function that, when passed a string of alphanumeric characters, returns a string of digits, Each character that is in the input string is converted to the digit that corresponds to it on a phone keypad:
"""

def convertToDigit(str):
    number = ""
    for ch in str:
        ch = ch.lower()
        if ch in "abc":
            number += '2'
        elif ch in "def":
            number += '3'
        elif ch in "ghi":
            number += '4'
        elif ch in "jkl":
            number += '5'
        elif ch in "mno":
            number += '6'
        elif ch in "pqrs":
            number += '7'
        elif ch in "tuv":
            number += '8'
        elif ch in "wxyz":
            number += '9'
        else:
            number += ch
    return number

DATA = ("peanut", "salt", "lemonade", "good time", ":10", "Zilch");

for item in DATA:
    print convertToDigit(item)
