import re
with open("input.txt", "r") as f:
    program = f.read()

pos = 0
muls_sum = 0
mul_p = re.compile(r"mul\([0-9]{1,3},[0-9]{1,3}\)")
num_p = re.compile(r"[0-9]{1,3}")
for mul in re.findall(mul_p,program):
    x,y = re.findall(num_p,mul)
    muls_sum += int(x)*int(y)
print(muls_sum)

