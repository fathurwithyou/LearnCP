#include <bits/stdc++.h>

using namespace std;

vector<int> dp(1e5);

int main(){
    int n, m; cin >> n >> m;
    int arr[n+1];
    dp[0] = 1;
    dp[1] = 1;
    for(int i = 1; i <= n; i++) cin >> arr[i];
    for(int i = 2; i <= n-1; i++){
        dp[i] = dp[i-1];
        if(!arr[i])
        dp[i] = dp[i] * (3-abs(arr[i-1]-arr[i+1]));
    } 
    for(int i = 1; i <= n; i++){
        cout << dp[i] << " ";
    }
    cout << endl;
    cout << dp[n];
}

/*
4 5
3 4 4 5
b-a = 3
zero = 2
a x y b
zero = banyak nol kontigu
if menaik
x = a+(b-zero)

|| |

inclusion/exclusion
total = 3^zero
kurangi = 

2 1 2 3



*/