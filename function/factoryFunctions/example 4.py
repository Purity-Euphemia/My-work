def createGreeting(greeting):
    def greet(name):
        return greeting + ", " + name + "!"

    return greet