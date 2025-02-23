#include <bits/stdc++.h>

using namespace std;


vector<vector<int>> dp(2, vector<int>(1001));

const int MOD = 1e9 + 7;
char grid[1001][1001];
int main(){
    int n; cin >> n;
    for(int i = 1; i <= n; i++){
        for(int  j = 1; j <= n; j++){
            cin >> grid[i][j];
        }
    }
    dp[0][1] = grid[1][1] == '.';
    for(int i = 1; i <= n; i++){
        for(int j = 1; j <= n; j++){
            if(grid[i][j] == '.'){
                dp[1][j] = (dp[1][j-1] + dp[0][j]) % MOD;
            }else dp[1][j] = 0;
        }
        dp[0] = dp[1];
    }
    cout << dp[1][n];
}