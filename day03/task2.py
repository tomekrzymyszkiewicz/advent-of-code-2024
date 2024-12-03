import re

with open("input.txt", "r") as f:
    program = f.read()

pos = 0
muls_sum = 0
command_p = re.compile(r"mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don't\(\)")
num_p = re.compile(r"[0-9]{1,3}")
last_command = "do()"
for command in re.findall(command_p, program):
    match command:
        case command if "mul" in command:
            if last_command == "do()":
                x, y = re.findall(num_p, command)
                muls_sum += int(x) * int(y)
            elif last_command == "don't()":
                continue
        case command if "do" in command:
            last_command = command


print(muls_sum)
