initials = {}
for name in ("peter", "anders", "bengt", "bengtsson"):
    initial = name[0]
    # if initial not in initials:
    #     initials[initial] = 0
    initials.setdefault(initial, 0)
    # 使用 setdefault 方法：对于每个首字母，
    # 使用 setdefault 方法检查字典中是否已存在该键。
    # 如果不存在，则设置其值为 0；如果已存在，则不做改动。
    initials[initial] += 1

print(initials)
# outputs
# {'a': 1, 'p': 1, 'b': 2}