#include <bits/stdc++.h>

using namespace std;

#define int unsigned long long

int binexp(int a, int b, int MOD){
    int res = 1;
    while(b > 0){
        if(b & 1) res = res * a % MOD;
        a = a * a % MOD;
        b >>= 1;
    }
    return res;
}

signed main(){
    ios::sync_with_stdio(0);
    cin.tie(0);
    int a,b,c,n; cin >> a >> b >> c >> n;
    if (n == 1){
        cout << 0;
        return 0;
    }
    cout << binexp(a, binexp(b, c, n-1), n);
}


/*
a^b^c mod n = pow(a, pow(b, c, n-1), n)
*/