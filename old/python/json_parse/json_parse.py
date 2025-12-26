def sum_json(data):
    sub_total = 0
    if isinstance(data, dict):
        for item in data.values():
            sub_total += sum_json(item)
    elif isinstance(data, list):
        for item in data:
            sub_total += sum_json(item)
    else:
        sub_total += data
    return sub_total


if __name__ == "__main__":
    data = {"a": 10, "b": {"c": 20, "d": {"e": 30}}, "f": [1, 2, {"g": 3}]}
    data2 = {"k1": 10, "k2": [5, 6], "k3": {"k4": 30}}

    print(sum_json(data))
    print(sum_json(data2))
