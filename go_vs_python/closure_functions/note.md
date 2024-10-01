请注意，在 Python 示例中，您可以访问内部函数中的数字，但不能更改它。假设您想这样做

```
def increment(amount):
    number += amount
increment(1)
increment(2)
```
然后你会得到一个UnboundLocalError错误，因为变量会被绑在内的范围增加功能。 注：可以使用的全局声明，获得的周围，例

def increment(amount):
    global number
    number += amount
increment(1)
increment(2)