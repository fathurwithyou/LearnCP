#include <bits/stdc++.h>

using namespace std;

int main(){
    int n, x; cin >> n >> x;
    int arr[n];
    for(int i = 0; i < n; i++){
        cin >> arr[i];
    }
    sort(arr, arr+n);
    int l = 0, r = n-1, ans = 0;
    while(l <= r){
        if(arr[l] + arr[r] <= x){
            l++;
        }
        r--;
        ans++;
    }
    cout << ans << endl;
}