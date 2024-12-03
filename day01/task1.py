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

result = sum([abs(i - j) for i, j in zip(sorted(one), sorted(two))])

print(result)
assert result == 2192892
