#!/usr/bin/env python
""" randrange.py Rolls dice, demonstrating random.randrage(,
 and a tuple with accessing a particular element with an index"""

import random
doubles = ("Can't happen", "Snake eyes!", "Little Joe!", "Hard six!", "Hard eight!","Fever!","Box cars!")

#Rolling Two Dices
def Rollem():
     #Going to random module and find a randrage function
     dice = random.randrange(1,7), random.randrange(1,7);
     print "%d and %d" % dice
     if dice[0] == dice[1]:
         print doubles[dice[0]]

while True:
    response = raw_input("Ready to roll?!");
    if response[0]  in "Qq":
        break
    Rollem()
