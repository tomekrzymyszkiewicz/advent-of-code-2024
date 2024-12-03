with open("input.txt", "r") as f:
    one, two = [], []
    for line in f.read().split("\n"):
        elements = " ".join(line.split()).split()
        one.append(int(elements[0]))
        two.append(int(elements[1]))

similarity_score = 0
for i in one:
    similarity_score += i * two.count(i)
print(similarity_score)
