#!/usr/bin/env python
"""Write a funciton called "Coin" that emulates the flip of a coin, returning 'heads' or 'tails/ """


import random

#Flipping a Coin
def Coin():
    #Going to flip the Coin
    coin = random.randrange(0,2)
    if coin == 0:
        flip = "heads"
    else:
         flip ="tails"
    return flip;

print Coin()
