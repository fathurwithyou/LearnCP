#include <bits/stdc++.h>

using namespace std;

int main(){
    int n, m; cin >> n >> m;
    multiset<int> s;
    for(int i = 0; i < n; i++){
        int x; cin >> x;
        s.insert(x);
    }
    while(m--){
        int x; cin >> x;
        auto it = s.upper_bound(x);
        if(it != s.begin()){
            cout << *--it << "\n";
            s.erase(it);
        }else cout << "-1\n";
    }
}