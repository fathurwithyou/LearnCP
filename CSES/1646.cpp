#include <bits/stdc++.h>

#define int long long
using namespace std;
signed main(){
    ios_base::sync_with_stdio(0);
    cin.tie(0);
    int n,q; cin >> n >> q;
    int pref[n+1] = {};
    for(int i = 1; i <= n; i++){
        cin >> pref[i];
        pref[i] += pref[i-1];
    }
    while(q--){
        int a, b; cin >> a >> b;
        cout << pref[b] - pref[--a] << "\n";
    }
}