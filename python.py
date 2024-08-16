def binarySearch(list, item):
    low = 0
    high = len(list) - 1
    while low <=high:
        mid = (low + high)
        guess =list[mid]
    if guess == item:
        return mid
    if guess > item:
        high - 1
    else:
        low = mid + 1
        return None
    MyList = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 17 , 100]
    print(binarySearch (list, 17)) 
    print("helo world")

