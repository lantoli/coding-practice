#include <algorithm>
#include <cstdio>
#include <cstdlib>
#include <list>
#include <map>
#include <queue>
#include <set>
#include <stack>
#include <vector>
#include <cmath>
#include <cstring>
#include <string>
#include <iostream>
#include <complex>
#include <sstream>
using namespace std;
 
typedef long long LL;
typedef unsigned long long ULL;
typedef long double LD;
typedef vector<int> VI;
typedef pair<int,int> PII;
 
#define REP(i,n) for(int i=0;i<(n);++i)
#define SIZE(c) ((int)((c).size()))
#define FOR(i,a,b) for (int i=(a); i<(b); ++i)
#define FOREACH(i,x) for (__typeof((x).begin()) i=(x).begin(); i!=(x).end(); ++i)
#define FORD(i,a,b) for (int i=(a)-1; i>=(b); --i)
#define ALL(v) (v).begin(), (v).end()
 
#define pb push_back
#define mp make_pair
#define st first
#define nd second

int N,M,K;
int T[2000001];

void scase() {
  scanf("%d%d%d",&N,&M,&K);
  REP(i,2000001) T[i] = 999999999;
  FOR(i,1,N+1)FOR(j,1,M+1) T[i*j] = min(T[i*j], (i == 1 || j == 1) ? max(i,j) : 2 * (i + j - 2));
  FORD(i,N*M,0) T[i] = min(T[i], T[i+1]);
  
  if (N > M) swap(N,M);
  int best = N * M;
  REP(a1,N+1)REP(b1,N+1)REP(a2,N+1)REP(b2,N+1) {
    if (a1 + b1 >= N) continue;
    if (a2 + b2 >= N) continue;
    if (a1 + a2 >= M) continue;
    if (b1 + b2 >= M) continue;
    
    int needed = K + (a1 * (a1 + 1)) / 2 + (b1 * (b1 + 1)) / 2 + (a2 * (a2 + 1)) / 2 + (b2 * (b2 + 1)) / 2;
    
    int X = T[needed] - a1 - a2 - b1 - b2;
    best = min(best, X);
  }
    
  printf("%d\n", best);
}

int main() {
  int Z;
  scanf("%d", &Z);
  REP(z,Z) {
    printf("Case #%d: ", z+1);
    scase();
  }
}    
