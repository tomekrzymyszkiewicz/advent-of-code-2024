with open("input.txt", "r") as f:
    one, two = [], []
    for line in f.read().split("\n"):
        elements = " ".join(line.split()).split()
        one.append(int(elements[0]))
        two.append(int(elements[1]))

sum_of_diffs = 0
while one or two:
    sum_of_diffs += abs(one.pop(one.index(min(one))) - two.pop(two.index(min(two))))
print(sum_of_diffs)