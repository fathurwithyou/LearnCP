#include <bits/stdc++.h>

using namespace std;

int main(){
    ios_base::sync_with_stdio(0);
    cin.tie(0);
    int n; cin >> n;
    for(int k = 1; k <= n; k++){
        long long ans = (long long)k*k*(k*k-1)/2;
        ans -= (long long) 4 * (k-1) * (k-2); 
        cout << ans << "\n";
    }
}

/*
0 1 
0 0 

1 1 
0 0

1 0 0
0 0 0
0 0 0 -> 6

0 1 0
0 0 0
0 0 0 -> 6

0 0 1
0 0 0
0 0 0 -> 6

1 0 0 
0 0 1

0 0 1
1 0 0

inclusion/exclusion
perm = k * k * (k*k-1)
comb = perm / 2! 
attack = 4 * (k - 1) * (k - 2)



*/