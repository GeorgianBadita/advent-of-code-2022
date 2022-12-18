def read_from_file(file):
    with open(file, "r") as f:
        lines = f.readlines()

        res = []
        for ln in lines:
            if len(ln) == 0 or ln == " " or ln == "\n":
                continue
            res.append(eval(ln))

        return [(res[i], res[i + 1]) for i in range(0, len(res) - 1, 2)]


def cmp(left, right):
    if type(left) is list and type(right) is list:
        for l, r in zip(left, right):
            rr = cmp(l, r)
            if rr:
                return rr

        if len(left) < len(right):
            return -1
        elif len(left) > len(right):
            return 1
        return 0
    elif type(left) is int and type(right) is int:
        if left < right:
            return -1
        elif left > right:
            return 1
        return 0

    elif type(right) is int:
        return cmp(left, [right])
    else:
        return cmp([left], right)


def solve(lists):
    res = 0
    for idx in range(len(lists)):
        left, right = lists[idx]
        if cmp(left, right) == -1:
            res += idx + 1

    return res


res = read_from_file("./in-day13.txt")
print(solve(res))
