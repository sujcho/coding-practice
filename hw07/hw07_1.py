#!/usr/bin/env python
# -*- coding: utf-8 -*-
"""
Python identiﬁers that tests a string.
"""
import keyword

def IsValidIdentifier(my_str):
    """
    IsValidIdentifier ( ) receives a string and returns a tuple (Boolean, reason string).
    Returns True, if:
    (a)   The ﬁrst character must be a letter or underscore.
    (b)   Letters beyond the ﬁrst can be alphanumeric or underscore.
    (c)   Identiﬁers can not be keywords.
    Else returns False with a reason string.
    """

    #Test the string for if it is a keyword
    if keyword.iskeyword(my_str):
        return (False, "%s: %s" % ("Invalid", "this is a keyword!"))

    #Test if the first character is a letter or underscore
    if not my_str[0].isalpha() and my_str[0] != '_':
        return (False, "%s: %s" % ("Invalid", "first symbol must be alphabetic or underscore."))

    #For each character in the string,
    for ch in my_str[1:]:
        #if not, the string is invalid
        if not ch.isalnum() and ch != '_':
            return (False, "%s: '%s' %s" % ("Invalid", ch, "is not allowed."))

    return (True, "Valid!")

if __name__ == '__main__':
    DATA = ('x', '_x', '2x', 'x,y ', 'yield', 'is_this_good')
    for case in DATA:
        result = IsValidIdentifier(case)
        print "%s -> %s" % (case, result[1])
