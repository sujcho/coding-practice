class PoliteList(list):
    def __str__(self):
        return """Here are your data:
%s
Thank you for asking.""" % (list.__str__(self))

print PoliteList([1,2,3])
