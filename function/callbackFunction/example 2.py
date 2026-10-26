# a Callback Function

def orderComplete(item):
    return "Order for " + item + " is ready!"

def processOrder(item, callback):
    return callback(item)

def solution(item):
    return processOrder(item, orderComplete)