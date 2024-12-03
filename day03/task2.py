import re
from functools import reduce
from operator import mul

with open("input.txt", "r") as f:
    program = f.read()

pos = 0
muls_sum = 0
command_p = re.compile(r"mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don't\(\)")
num_p = re.compile(r"[0-9]{1,3}")
last_command = "do()"

for command in re.findall(command_p, program):
    if "mul" in command and last_command == "do()":
        muls_sum += reduce(mul, map(int, re.findall(num_p, command)))
    else:
        last_command = command

print(muls_sum)
assert muls_sum == 100411201
