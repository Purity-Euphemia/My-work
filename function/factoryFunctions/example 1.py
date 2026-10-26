def createDiscount(percentage):
    def calculateDiscount(price):
        return price - (price * percentage / 100)

    return calculateDiscount