#!usr/bin/env python

"""Write a funciton that returns a total cost from the sales price and sales tax rate. Both price and tax_rate shoudl be in the argument list. Provide a defualt value for the tax rate = 8.25%. Test your function."""

def Cost(sales, tax_rate="8.25"):
    print sales * (1 + (float(tax_rate) * 0.01))

Cost(100)

"""Write a breakfast function that takes five arguments: meat, eggs, potatoes, toast, beverage"""

def Breakfast(meat="bacon", eggs="over easy", potatoes="hash brown", toast="white", beverage="coffee"):
    print 'Here is your', meat, 'and', eggs ,'with' , potatoes  , 'and' , toast
    print 'Can I bring you more' , beverage ,'?'

Breakfast();
