import sys


def main():
    n, t = map(int, sys.stdin.readline().split())
    line = []
    for i in range(n):
        line.append(i+1)

    idx = 0
    counter = 1
    tt = []
    while(len(line) - idx > 0):
        if counter % t == 0:
            tt.append(line[idx])
        else:
            line.append(line[idx])
        idx += 1
        counter += 1

    print_str = '<'
    for idx, i in enumerate(tt):
        print_str += str(i)
        if idx + 1 < len(tt):
            print_str += ', '
    print_str += '>'
    print(print_str)

if __name__ == '__main__':
    main()
