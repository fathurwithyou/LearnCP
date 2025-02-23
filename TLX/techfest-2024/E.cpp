#include <bits/stdc++.h>

using namespace std;

#define int long long

signed main(){
    ios::sync_with_stdio(0);
    cin.tie(0);
    int n; cin >> n;
    int arr[100001] = {};
    for(int i = 0; i < n; i++){
        int x; cin >> x;
        arr[x]++;
    }
    int ans = 1, i;
    for(i = 1; i <= 100000 && arr[i] >= ans; i++){
        ans = max(ans, arr[i]);
    }
    cout << (i > 1 ? ans : 0);
}