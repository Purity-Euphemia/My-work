# a Factory Function
def createDiscount(percentage):
    def discountCalculator(price):
        return price - (price * percentage / 100)
    return discountCalculator

def solution(percentage, price):
    discount = createDiscount(percentage)
    return discount(price)