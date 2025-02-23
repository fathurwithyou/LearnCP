#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
const int MAX_N = 1e5;
int A[MAX_N];
pair<int, int> arr[MAX_N];
int main(){
    ios::sync_with_stdio(false);
    cin.tie(nullptr);
    int n;
    ll K, X;
    cin >> n >> K >> X;
    for (int i = 0; i < n; i++){
        cin >> A[i];
        arr[i] = {A[i], i};
    }
    sort(arr, arr + n, [](auto &a, auto &b){ return a.first < b.first; });
    int mini = INT_MAX;
    for (int i = 0; i < n; i++){
        mini = min(mini, A[i]);
    }
    ll lo = mini, hi = (ll)mini + K * X, ans = mini;
    while (lo <= hi){
        ll mid = (lo + hi) / 2;
        ll sum = 0;
        for (int i = 0; i < n; i++){
            if (arr[i].first < mid){
                ll diff = mid - arr[i].first;
                sum += (diff + X - 1) / X;
                if (sum > K) break;
            }
        }
        if (sum <= K){
            ans = mid;
            lo = mid + 1;
        } else {
            hi = mid - 1;
        }
    }
    cout << ans;
    return 0;
}
