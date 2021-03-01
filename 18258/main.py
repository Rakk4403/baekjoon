import sys


def main():
    stack = []
    n = int(sys.stdin.readline())

    idx = 0
    cmd_list = []
    for i in range(n):
        cmd_list.append(str(sys.stdin.readline().replace('\n', '')))
    
    for cmd in cmd_list:
        if cmd == 'front':
            if len(stack) == idx:
                print(-1)
                continue
            print(stack[idx])
        elif cmd == 'back':
            if len(stack) == idx:
                print(-1)
                continue
            print(stack[len(stack)-1])
        elif cmd == 'size':
            print(len(stack) - idx)
        elif cmd == 'empty':
            if len(stack) == idx:
                print(1)
            else:
                print(0)
        elif cmd == 'pop':
            if len(stack) == idx:
                print(-1)
                continue
            print(stack[idx])
            idx += 1
        else:
            # push x
            cmd_split = cmd.split()
            stack.append(int(cmd_split[1]))


if __name__ == '__main__':
    main()
