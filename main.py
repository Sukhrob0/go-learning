# # a = 10
# # b = 4 
# # for i in range(10):
# #     c = a + b
# #     print(i, c)

# def binarySearch(list, item):
#     low = 0
#     high = len(list) - 1
#     while low <= high:
#         mid = (low + high)
#         guess = list[mid]
#     if guess == item:
#         return mid
#     if guess > item:
#         high - 1
#     else:
#         low = mid + 1
#         return None

# MyList = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 17 , 100]
# print(binarySearch(MyList, 6))  # Output: 16 (the first index where 17 is found)

 

# # def binarySearch(list, item):
# #     low = 0
# #     high = len(list) - 1
# #     while low <= high:
# #         mid = (low + high) // 2  # Add // 2 for correct middle index
# #         guess = list[mid]
# #         if guess == item:
# #             return mid
# #         if guess > item:
# #             high = mid - 1  # Correct indentation and missing return
# #         else:
# #             low = mid + 1
# #     return None  # Return None outside the loop if not found


