def binexp(a, b, MOD):
    res = 1
    while b > 0:
        if b & 1:
            res = res * a % MOD
        a = a * a % MOD
        b >>= 1
    return res

def main():
    a, b, c, n = map(int, input().split())
    if n == 1:
        print(0)
        return
    bc_mod = binexp(b, c, n-1)
    result = binexp(a, bc_mod, n)
    print(result)

if __name__ == "__main__":
    main()
