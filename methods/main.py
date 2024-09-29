from math import sqrt

class Point(object):
    def __init__(self, x, y):
        self.x = x
        self.y = y
    def distance(self,other):
        return sqrt((self.x - other.x)**2 + (self.y - other.y)**2)

p1 = Point(1,2)
p2 = Point(2,4)

print(p1.distance(p2)) 
print(p2.distance(p1)) 