with open("input.txt", "r") as f:
    one, two = map(
        sorted,
        map(
            list,
            zip(
                *[
                    map(int, " ".join(line.split()).split())
                    for line in f.read().split("\n")
                ]
            ),
        ),
    )

similarity_score = sum([i * two.count(i) for i in one])

print(similarity_score)
assert similarity_score == 22962826
