#include <bits/stdc++.h>

using namespace std;
int dp[100001];
void precompute(){
    for(int i = 1; i*i<= 100000; i++){
        for(int j = i; j <= 100000; j++)
    }
}

int main(){
    int n; cin >> n;
    precompute();
    while(n--){
        int x; cin >> x;
        cout << dp[x] << "\n";
    }
}