#include <bits/stdc++.h>

using namespace std;

#define int long long

int dp[1000001];
const int MOD = 1e9 + 7;
signed main(){
    ios_base::sync_with_stdio(0);
    cin.tie(0);
    int x,n; cin >> n >> x;
    int arr[n+1];
    dp[0] = 1;
    for(int i = 1; i <= n; i++)cin >> arr[i];
    for(int i = 1; i <= n; i++){
        for(int j = 1; j <= x; j++){
            if(arr[i] <= j)
            dp[j] = (dp[j] + dp[j-arr[i]]) % MOD;
        }
    }
    cout << dp[x];
}