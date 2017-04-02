#!/usr/bin/env python
# -*- coding: utf-8 -*-
"""
This module will ask for a starting directory name and will look for and report any words that are palindromes throughout the directory structure.
"""
import os
import sys


def GetText(file_name):
    """Opens a file and returns content"""
    try:
        #if there is a file, read the content
        open_file = open(file_name)
        try:
            text = open_file.read()
        finally:
            open_file.close()

    except IOError, msg:
        print file_name, msg

    return text

def FindPalindromes(anything, dirname, fnames):
    """
    Find plaindromes in a path tree
    """
    #import hw11_2 as a moudle
    sys.path.insert(0,"..")
    import utils.hw11_2 as pal

    for file_name in fnames:
        full_path = os.path.join(dirname,file_name)

        #if it is not a file, skip
        if not os.path.isfile(full_path):
            continue
        #if it is a file, get content of it
        content = GetText(full_path)
        content = content.split()

        for word in content:
            #check palindromes
            result = pal.Palindromize(word)
            if not result == None:
                print result

def main():
    while True:
        print "Enter starting directory"
        starting_dir = raw_input()

        if os.path.exists(starting_dir):
            break
        else:
            print "Path does not exists. Please check again"

    os.path.walk(starting_dir, FindPalindromes, "Walking:")

if __name__ == "__main__":
    main()
