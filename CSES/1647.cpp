#include <bits/stdc++.h>

#define int long long
using namespace std;

int main(){
    int n, q; cin >> n >> q;
    int pref[n+1], arr[n+1];
    pref[0] = INT_MAX;
    for(int i = 1; i <= n; i++){ 
        cin >> arr[i];
        pref[i] = min(pref[i], pref[i-1]);
    }
    while(q--){
        int a, b; cin >> a >> b;
        if(a==b)
        cout << min(pref[b], pref[a]);
    }
}