def get_answer():
    count_zeroes = 0
    current_pos = 50 
    MOD = 100
    
    with open("./input.txt") as f:
        # Use strip() to handle trailing newlines safely
        lines = [line.strip() for line in f if line.strip()]

    for line in lines:
        direction = line[0]
        distance = int(line[1:])

        if direction == 'L':
            # Subtract for Left. Python's % handles negatives correctly.
            current_pos = (current_pos - distance) % MOD
        elif direction == 'R':
            # Add for Right.
            current_pos = (current_pos + distance) % MOD

        # If it lands exactly on 0 after the move
        if current_pos == 0:
            count_zeroes += 1

    return count_zeroes

print(get_answer())
