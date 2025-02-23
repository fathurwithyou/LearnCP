#include <bits/stdc++.h>

using namespace std;

#define int long long

int comb(int n){
    return n * (n+1) / 2;
}


signed main(){
    ios::sync_with_stdio(0);
    cin.tie(0);
    int n; cin >> n;
    int arr[n];  

    for(int i = 0; i < n; i++) cin >> arr[i];
    int l = 0;
    int cnt = 0;
    for(int i = 1; i < n; i++){
        if(arr[i] <= arr[i-1]){
            cnt += comb(i-l);
            l = i;
        }
    }
    cout << cnt + comb(n-l) << endl;
}