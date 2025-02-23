#include <bits/stdc++.h>

using namespace std;

#define int long long

int solve(int l, int r, int k, int arr[]){
    if(l > r){
        cout << 0;
        return 0;
    }
    int m = (l+r)/2;
    

}

signed main(){
    int n, m, k; cin >> n >> m >> k;
    int arr[m];
    for(int i = 0; i < m; i++) cin >> arr[i];
    sort(arr, arr+m);
    if(k == n){
        cout << 0;
        return 0;
    }
    cout << solve(0, m-1, k, arr);
}