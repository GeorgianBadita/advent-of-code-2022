from functools import cmp_to_key


def read_from_file(file):
    with open(file, "r") as f:
        lines = f.readlines()

        res = []
        for ln in lines:
            if len(ln) == 0 or ln == " " or ln == "\n":
                continue
            res.append(eval(ln))

        return res


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
    lists += [[[2]], [[6]]]
    lists = sorted(lists, key=cmp_to_key(cmp))
    return (lists.index([[2]]) + 1) * (lists.index([[6]]) + 1)


res = read_from_file("./in-day13.txt")

print(solve(res))
