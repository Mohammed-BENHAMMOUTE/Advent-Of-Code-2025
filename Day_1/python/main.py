def get_answer_1():
    count_zeroes = 0
    current_pos = 50
    MOD = 100

    with open("./input.txt") as f:
        lines = [line.strip() for line in f if line.strip()]

    for line in lines:
        direction = line[0]
        distance = int(line[1:])

        if direction == "L":
            current_pos = (current_pos - distance) % MOD
        elif direction == "R":
            current_pos = (current_pos + distance) % MOD
        if current_pos == 0:
            count_zeroes += 1

    return count_zeroes


def get_answer_2():
    count_zeroes = 0
    current_pos = 50
    MOD = 100

    with open("./input.txt") as f:
        lines = [line.strip() for line in f if line.strip()]

    for line in lines:
        direction = line[0]
        distance = int(line[1:])

        if direction == "L":
            count_zeroes += distance // 100
            rest = distance % 100

            if current_pos > 0 and rest >= current_pos:
                count_zeroes += 1

            current_pos = (current_pos - distance) % MOD
        elif direction == "R":
            count_zeroes += (current_pos + distance) // 100
            current_pos = (current_pos + distance) % MOD

    return count_zeroes


print(get_answer_1())
print(get_answer_2())
