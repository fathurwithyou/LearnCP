#include <bits/stdc++.h>

using namespace std;

#define int long long

/*
0
1

00 10
01 11
11 01
10 10

10
*/

vector<string> solve(int n){
    if(n == 1){
        return {"0", "1"};
    }
    vector<string> prev = solve(n-1);
    vector<string> rev = prev;
    reverse(rev.begin(), rev.end());
    int i = 0, m = prev.size();
    while(i < m){
        string tmp = "0" + prev[i];
        prev[i] = "1" + rev[i++];
        prev.push_back(tmp);
    }
    return prev;
}

signed main()
{
    cin >> n;
    vector<string> ans = solve(n);
    for (auto a : ans)
        cout << a << "\n";
}

/*
000
001
011
010
111
110
100
101

*/