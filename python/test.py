from collections import defaultdict


def generator_function():
    for i in range(10):
        yield i
        yield 2


gen = generator_function()

# print(next(gen))
# print(next(gen))

#generator = (num ** 2 for num in range(10))
# for num in generator:
# print(num)

number_list = range(-5, 5)
less_than_zero = list(filter(lambda x: x < 0, number_list))
print(less_than_zero)

print(__name__)
