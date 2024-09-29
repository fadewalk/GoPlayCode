
def upone(mutable,index):
    mutable[index] = mutable[index].upper()
    
list_ = ['a','b','c']

upone(list_,1)

print(list_)

dict_ = {'a':"anders",'b':"begent"}
upone(dict_,'b')

print(dict_)

