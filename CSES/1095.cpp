#include <bits/stdc++.h>

using namespace std;

#define int long long
const int MOD = 1e9+7;

int binexpo(int a, int b){
    if(b == 0) return 1;
    int res = binexpo(a, b/2);
    res *= res;
    res %= MOD;
    if(b&1)
        res *= a;
    // } else res *= res;
    return res%MOD;
}

signed main(){
    int q;
    int n; cin >> q;
    while(q--){
        int a, b; cin >> a >> b;
        cout << binexpo(a, b) << "\n";
    }
}