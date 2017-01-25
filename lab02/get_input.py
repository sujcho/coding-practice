#!/usr.bin/env python

answer = raw_input('What is your favorite number?')
print 'You like', answer, '?'
print 'You really like ' + answer + '?'
number = int(answer)
if number < 10:
    print number, 'is small.'
elif number >= 1000:
    print number, 'is big.'
else:
    print number, 'is medium.'

    print 'But you like ' + str(number) + '!'
