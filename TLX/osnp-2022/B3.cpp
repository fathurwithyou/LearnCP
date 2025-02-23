#include <bits/stdc++.h>
using namespace std;

const int MAX_N = 1e5;
int E[MAX_N], pos_map[MAX_N], R[MAX_N], L[MAX_N], ans[MAX_N];
pair<int, int> arr[MAX_N];

// Fast input implementation
char inputBuffer[1 << 20];
int pos = 0;
int sz = 0;

inline char getChar()
{
    if (pos == sz)
    {
        pos = 0;
        sz = fread(inputBuffer, 1, 1 << 20, stdin);
        if (!sz)
            return EOF;
    }
    return inputBuffer[pos++];
}

inline int getInt()
{
    int x = 0;
    char c = getChar();
    while (c < '0' || c > '9')
        c = getChar();
    while (c >= '0' && c <= '9')
    {
        x = (x << 3) + (x << 1) + (c - '0');
        c = getChar();
    }
    return x;
}

int main()
{
    ios::sync_with_stdio(false);
    cin.tie(nullptr);
    cout.tie(nullptr);

    int n = getInt();
    int P = getInt();
    int q = getInt();

    for (int i = 0; i < n; i++)
    {
        E[i] = getInt();
        arr[i] = {E[i], i};
    }

    sort(arr, arr + n, [](auto &a, auto &b)
         { return a.first < b.first; });

    for (int i = 0; i < n; i++)
        pos_map[arr[i].second] = i;

    int j = 0;
    for (int i = 0; i < n; i++)
    {
        while (j < n && arr[j].first - arr[i].first <= P)
            j++;
        R[i] = j - 1;
    }

    int l = 0;
    for (int j = 0; j < n; j++)
    {
        while (arr[j].first - arr[l].first > P)
            l++;
        L[j] = l;
    }

    deque<int> dq;
    for (int pos = 0; pos < n; pos++)
    {
        while (!dq.empty() && R[dq.front()] < pos)
            dq.pop_front();
        while (!dq.empty() && (R[dq.back()] - dq.back()) <= (R[pos] - pos))
            dq.pop_back();
        dq.push_back(pos);
        ans[pos] = R[dq.front()] - dq.front() + 1;
    }

    deque<int> dq2;
    for (int pos = n - 1; pos >= 0; pos--)
    {
        while (!dq2.empty() && L[dq2.front()] > pos)
            dq2.pop_front();
        while (!dq2.empty() && (dq2.back() - L[dq2.back()]) <= (pos - L[pos]))
            dq2.pop_back();
        dq2.push_back(pos);
        ans[pos] = max(ans[pos], dq2.front() - L[dq2.front()] + 1);
    }

    for (int i = 0; i < q; i++)
    {
        int X = getInt();
        int pos = pos_map[X - 1];
        cout << ans[pos] << "\n";
    }

    return 0;
}