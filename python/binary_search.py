
from __future__ import print_function

"""1 2 3, 4 5 6, 7 8 9 : median = 5"""

""" 1 2 3 4 5 6 7 8 9 에서 7보다 작은 엘리먼트 인덱스 찾기 """


import math

def findMedianInRow(row, x):


    start = 0;
    end = len(row)-1

    print("x",x)
    while(start <= end):
        mid = math.floor((start + end)/2)
        print("row[mid]",row[mid], mid)
        if row[mid] == x:
            return mid + 1

        if row[mid] > x:
            end = mid - 1
        else:
            start = mid + 1

    if x > row[end]:
        return end + 1
    else:
        return 0

def findMedian(A):

    n = len(A[0])
    m = len(A)
    med_idx = math.floor(n*m/2)

    mid = 0
    start = 0
    end = 1000000

    while start <= end :
        mid = int(math.floor((start + end)/2))
        count = 0

        for i in range(0,m,1):
            count += findMedianInRow(A[i],mid)
            print("count", count)

        count = count -1

        if count == med_idx :
            return mid

        if count < med_idx :
            start = mid + 1

        else :
            end = mid - 1

    return mid


def main():
    A = [[1,1,1,1,3,3,3]]

    print(findMedian(A))

if __name__ == '__main__':
    main()
