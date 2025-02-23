#include <bits/stdc++.h>

using namespace std;

int main() {
    ios::sync_with_stdio(0);
    cin.tie(0);
    cout.tie(0);

    int n, x;
    cin >> n >> x;
    
    vector<pair<int, int>> a(n);
    
    for (int i = 0; i < n; i++) {
        cin >> a[i].first;
        a[i].second = i + 1;
    }

    sort(a.begin(), a.end());

    int l = 0, r = n - 1;

    while (l < r) {
        int sum = a[l].first + a[r].first;
        if (sum == x) {
            cout << a[l].second << " " << a[r].second << "\n";
            return 0;
        }
        if (sum < x) {
            l++;
        } else {
            r--;
        }
    }

    cout << "IMPOSSIBLE\n";
    return 0;
}
