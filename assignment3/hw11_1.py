#!/usr/bin/env python
# -*- coding: utf-8 -*-
"""
GiveAsciiChart returns ascii chart as a string
"""

def GiveAsciiChart1():
    """
    GiveAsciiChart1() is a version that take the readibility as the first prioriority.
    GiveAsciiChart1() returns ascii chart (only from value 32 to 126) as a string
    """
    chart = ""
    for num in range(32, 126):
        ascii_value = chr(num)
        chart += "%4s" % str(num)+":"+ ascii_value
        if (num+1)%4 == 0:
            chart += "\n"
    return chart

def GiveAsciiChart2():
    """
    GiveAsciiChart2() is a version with .
    GiveAsciiChart1() returns ascii chart (only from value 32 to 126) as a string
    """
    return "\n".join([" ".join(["%4s" % (str(num+i)+":"+ chr(num+i)) for i in range(0,4)]) for num in range (32,126,4)])

if __name__ == '__main__':
    GiveAsciiChart1()
    GiveAsciiChart2()
