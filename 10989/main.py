import os
import sys

def main():
    n = int(sys.stdin.readline())
    num_list = []
    max_num = 0
    for i in range(n):
        num = int(sys.stdin.readline())
        num_list.append(num)
        if max_num < num:
            max_num = num
        

    num_map = [0] * (max_num+1)
    for i in num_list:
        num_map[i] += 1

    counter = 0 
    for idx, num in enumerate(num_map):
        counter += num_map[idx]
        num_map[idx] = counter

    # create array
    newlist = [0] * (counter+1)
    for idx, num in enumerate(num_list):
        newlist[num_map[num]] = num
        num_map[num] -= 1

    for num in newlist:
        if num == 0:
            continue
        print(num)
    print(num_list)
    print(num_map)
    print(newlist)

if __name__ == '__main__':
    main()
