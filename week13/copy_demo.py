import copy

# a = [3, -99, 17]  # immutable values
a = [3, [5,76], 17]  # immutable values, mutable values, immutable values
b = a
c = a[:]  # shallow copy
d = list(a)  # shallow copy
e = a.copy()  # shallow copy
f=copy.deepcopy(a) # deep copy
print(a, b, c, d, e, f)
# a[1] = 100
a[1][0]=100
print(a, b, c, d, e, f)
