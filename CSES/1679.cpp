#include <bits/stdc++.h>

using namespace std;

vector<int> adj[100001];
bool vis[100001];
bool chk[100001];
stack<int> s;

bool impos;
void dfs(int u){
    vis[u] = 1;
    chk[u] = 1;
    for(auto a : adj[u]){
        if(!vis[a]){
            dfs(a);
        }
        if(chk[a]){
            cout<<"IMPOSSIBLE"; exit(0);
        }
    }
    chk[u] = 0;
    s.push(u);
}

int main(){
    ios_base::sync_with_stdio(0);
    cin.tie(0);
    int n, m; cin >> n >> m;
    for(int i = 0; i < m; i++){
        int x, y; cin >> x >> y;
        adj[x].push_back(y);
    }
    for(int i = 1; i <= n; i++) if(!vis[i]) dfs(i);
    while(!s.empty()){
        cout << s.top() << " "; 
        s.pop();
    }
}

