def createMultiplier(factor):
    def multiply(number):
        return number * factor

    return multiply