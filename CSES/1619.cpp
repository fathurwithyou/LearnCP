#include <bits/stdc++.h>

using namespace std;

#define int long long

signed main() {
    ios::sync_with_stdio(0);
    cin.tie(0);
    cout.tie(0);

    int n;
    cin >> n;
    
    vector<pair<int, int>> aa;
    
    for (int i = 0; i < n; i++) {
        int x, y;
        cin >> x >> y;
        aa.push_back({x, 1});  
        aa.push_back({y, -1});
    }
    
    sort(aa.begin(), aa.end());

    int ans = 0, maxi = 0;
    
    for (auto xxx : aa) {
        ans += xxx.second;
        maxi = max(maxi, ans);
    }
    
    cout << maxi << "\n";
}
