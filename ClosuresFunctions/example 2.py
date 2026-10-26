# Counter Function
def createCupCounter():
    count = 0

    def counter():
        nonlocal count
        count += 1
        return count

    return counter


def solution():
    counter = createCupCounter()
    return counter()