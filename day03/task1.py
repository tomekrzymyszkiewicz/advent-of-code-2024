import re
from functools import reduce
from operator import mul

with open("input.txt", "r") as f:
    program = f.read()

pos = 0
muls_sum = 0
mul_pattern = re.compile(r"mul\([0-9]{1,3},[0-9]{1,3}\)")
num_pattern = re.compile(r"[0-9]{1,3}")


muls_sum = sum(
    [
        reduce(mul, map(int, re.findall(num_pattern, mul_part)))
        for mul_part in re.findall(mul_pattern, program)
    ]
)

print(muls_sum)
assert muls_sum == 174103751
