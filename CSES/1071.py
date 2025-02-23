def main():
    t = int(input())
    for _ in range(t):
        y, x = map(int, input().split())
        start = max(y, x) * max(y, x)
        ans = 0
        
        if y > x:
            if y % 2 == 0:
                ans = start - x + 1
            else:
                ans = start + x
        else:
            if x % 2 == 0:
                ans = start - y + 1
            else:
                ans = start + y
        
        print(ans)

if __name__ == "__main__":
    main()