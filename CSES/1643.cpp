#include <bits/stdc++.h>

using namespace std;

#define int long long

signed main() {
    ios::sync_with_stdio(0);
    cin.tie(0);
    cout.tie(0);

    int n; cin >> n;
    int a[n];
    for(int i = 0; i < n; i++){
        cin >> a[i];
    }
    int ans = 0;
    int maxi = INT_MIN;
    // Kadane's Algorithm
    for(int i = 0; i < n; i++){
        ans += a[i];
        maxi = max(maxi, ans);
        if(ans < 0){
            ans = 0;
        }
    }
    cout << maxi;
}