#include <bits/stdc++.h>

using namespace std;

#define int long long
int MOD = 1e9+6;

int binexpo(int a, int b){
    if(b == 0) return 1;
    int res = binexpo(a, b/2);
    res *= res;
    res %= MOD;
    if(b&1)
        res *= a;
    return res%MOD;
}

/*

a^b^c = 

*/

signed main(){
    int q;
    int n; cin >> q;
    while(q--){
        int a, b, c; cin >> a >> b >> c;
        int bc = binexpo(b,c);
        ++MOD;
        cout << binexpo(a, bc) << "\n";
        --MOD;
    }
}