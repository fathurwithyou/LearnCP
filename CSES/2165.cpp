#include <bits/stdc++.h>

using namespace std;
/*
a b c
cout << a << c << endl;
a c b

b c a 
*/
void hanoi(int a, int b, int c, int n){
    if(!n) return;
    hanoi(a, c, b, n-1);
    cout << a << " " << c << "\n";
    hanoi(b, a, c, n-1);
} 

int main(){
    int n; cin >> n;
    int total = (1 << n) - 1;
    cout << total << "\n";
    hanoi(1,2,3,n);
}