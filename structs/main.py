
from math import sqrt

class Point(object):
    def __init__(self, x, y):
        self.x = x
        self.y = y

def distance(point1, point2):
    return sqrt((point1.x - point2.x)**2 + (point1.y - point2.y)**2)


p1 = Point(3, 4) 
p2 = Point(6, 8)

print(distance(p1, p2))
