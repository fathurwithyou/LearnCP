#include <bits/stdc++.h>

using namespace std;


vector<vector<int>> dp(2, vector<int>(1001));

int main(){
    int n; cin >> n;
    char grid[n][n];
    for(int i = 0; i < n; i++){
        for(int  j = 0; j < n; j++){
            cin >> grid[i][j];
        }
    }
    for(int i = 0; i < n; i++){
        for(int j = 0; j < n; j++){
            if(grid[i][j] == '.'){
                dp[1][j] = dp[1][j-1] + dp[0][j];
            }else dp[1][j] = 0;
        }
    }
    cout << dp[1][n];
}