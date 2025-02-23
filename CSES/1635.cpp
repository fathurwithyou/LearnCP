#include <bits/stdc++.h>

using namespace std;

vector<int> dp(1e6+1, 0);
int arr[101];
const int MOD = 1e9+7;
int main(){
    ios_base::sync_with_stdio(0);
    cin.tie(0);
    dp[0] = 1;
    int n, x; cin >> n >> x;
    for(int i = 1; i <= n; i++) cin >> arr[i]; 
    for(int i = 1; i <= x; i++){
        for(int j = 1; j <= n; j++){
            // dp[i] = dp[i-1];
            if(arr[j] <= i)
            dp[i] = (dp[i] + dp[i-arr[j]]) % MOD;
        }
        // cout << dp[i] << endl; 
    }
    cout << dp[x];
}

/*
3x9
2 3 5
1 0 1 0 1 0 1 0 1 0
1 0 0 1 0 0 2 3 3 0
1 0 0 0 0 1 2 0 0 0

2 3 5
1 1 1 0
0 0 0 1
1 1 1 2
1 1 1 3 
1 1 1 4
1 2 3 5
1 1 0 6
0 0 0 7
0 0 0 8
0 0 0 9
0 0 0 10

*/