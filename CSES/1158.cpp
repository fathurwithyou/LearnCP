#include <bits/stdc++.h>

using namespace std;

vector<vector<int>> dp(2, vector<int>(1e5+1));
int h[1001], s[1001];
int main(){
    int n, x; cin >> n >> x;
    for(int i = 1; i <= n; i++) cin >> h[i];
    for(int i = 1; i <= n; i++) cin >> s[i];
    for(int i = 1; i <= n; i++){
        for(int j = 1; j <= x; j++){
            dp[1][j] = dp[0][j];
            if(j >= h[i])
            dp[1][j] = max(dp[1][j], s[i]+dp[0][j-h[i]]);
        }
        dp[0] = dp[1];
    }
    cout << dp[1][x];
}

/*
T(n) = n^2 + n
cari C dan f(n) sehingga C*f(n) >= T(N)
C = 2 dan f(n) = n^2
sehingga 2*n^2 >= n^2 + n
didapat kompleksitas big Oh adalah O(f(n)) = O(n^2)
*/