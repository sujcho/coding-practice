def DoThis():
        def Report():
            print("test")
            print (amount)
        amount = 3
        Report()

def DoThis_2():
        def Report():
            print (amount)
        Report()
        amount = 3

def main():
    DoThis()
    DoThis_2()


if __name__ == "__main__":
    main()
