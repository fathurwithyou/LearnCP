#include <bits/stdc++.h>
using namespace std;

set<string> ans;
int ch[26];
set<char> av;
int n;
void dfs(int i, string s){
    if(i == n){
        ans.insert(s);
        return;   
    }
    for(auto a: av){
        if(ch[a-'a'] > 0){
            ch[a-'a']--;
            dfs(i+1, s+a);
            ch[a-'a']++;
        }
    }
}
int main(){
    string s; cin >> s;
    for(auto a: s){
        ch[a-'a']++;
        av.insert(a);
    }
    n = s.size();
    dfs(0, "");
    cout << ans.size() << "\n";
    for(auto a: ans){
        cout << a << "\n";
    }
}