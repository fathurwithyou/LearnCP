#include <bits/stdc++.h>
using namespace std;
#include <ext/pb_ds/assoc_container.hpp>
#include <ext/pb_ds/tree_policy.hpp>
using namespace __gnu_pbds;
vector<vector<bool>> vis(7, vector<bool>(7, false));
int dx[4] = {1, -1, 0, 0};
int dy[4] = {0, 0, 1, -1};
char dir[4] = {'R', 'L', 'D', 'U'};
string s;
int cnt = 0;

void dfs(int i, int j, int move)
{
    if (i == 6 && j == 0)
    {
        if (move == 48)
            cnt++;
        return;
    }

    for (int k = 0; k < 4; k++)
    {
        int ni = i + dy[k], nj = j + dx[k];
        if (s[move] != '?' && s[move] != dir[k])
            continue;

        if (ni >= 0 && ni < 7 && nj >= 0 && nj < 7 && !vis[ni][nj])
        {
            vis[ni][nj] = true;
            dfs(ni, nj, move + 1);
            vis[ni][nj] = false;
        }
    }
}

int main(){
    ios_base::sync_with_stdio(false);
    cin.tie(NULL);

    cin >> s;

    vis[0][0] = true;
    dfs(0, 0, 0);
    cout << cnt;

    return 0;
}