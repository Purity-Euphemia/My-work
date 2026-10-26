# Parent Variable

def ratioBinder(ratio):
    def getRatio():
        return ratio

    return getRatio


def solution(ratio):
    return ratioBinder(ratio)()