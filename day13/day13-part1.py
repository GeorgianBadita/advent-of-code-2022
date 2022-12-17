def read_from_file(file):
    with open(file, 'r') as f:
        lines = f.readlines()

        res = []
        for ln in lines:
            if len(ln) == 0 or ln == ' ' or ln == '\n':
                continue
            res.append(eval(ln))

        return [(res[i], res[i + 1]) for i in range(len(res) - 1)]


def cmp(left, right):
    pass


def solve(lists):
    res = 0
    for idx in range(len(lists)):
        left, right = lists[idx]
        if cmp(left, right):
            res += idx + 1

    return res


res = read_from_file("./in-day13.txt")
print(res)
