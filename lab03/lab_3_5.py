"""Print the decimal equivalent of a binary string that is given by the user:"""



input = raw_input("Enter a binary number: ")
sum = 0
mult = 0
for digit in input:
    if digit != '0' and digit != '1':
        print "Not a binary number"
        break;
    else:
        sum += int(digit) * pow(2,mult)
        mult += 1

if sum != 0:
    print "%d" % sum
