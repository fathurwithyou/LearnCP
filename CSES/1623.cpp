#include <bits/stdc++.h>

using namespace std;
#define int long long
int arr[20];
int n; 
int dfs(int i, int l, int r){
    if(i == n) return abs(l-r);
    return min(dfs(i+1, l+arr[i], r), dfs(i+1, l, r+arr[i]));
}

signed main(){
    cin >> n;
    for(int i = 0; i < n; i++)cin>>arr[i];
    cout << dfs(0, 0, 0);
}