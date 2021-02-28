import os
import sys

def main():
    n = int(sys.stdin.readline())
    num_list = []
    for i in range(n):
        item = int(sys.stdin.readline())
        num_list.append(item)

    num_list.sort()
    for i in num_list:
        print(i)

if __name__ == '__main__':
    main()
