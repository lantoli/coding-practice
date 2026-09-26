/*
Tuenti Challenge 4, Apr 2014, Challenge 19 - Password Recovery

Challenge not submitted to the contest.


First we simplify the simulated computer, see original_fn.
In the code we see which are the expected values for the registers, so we then reverse the operations to
try to find the original values, see reverse_fn.

The result is: 4834636b7c643165
h = 4834636b
l = 7c643165
r[0] = 'H';
r[1] = '4';
r[2] = 'c';
r[3] = 'k';
r[4] = '|';
r[5] = 'd';
r[6] = '1';
r[7] = 'e';

Code:

0, w: OP, x: ASG, y: 9, z: 0
1, w: OP, x: ADD, y: 9, z: 1
2, w: OP, x: ADD, y: 9, z: 2
3, w: OP, x: ADD, y: 9, z: 3
4, w: OP, x: ADD, y: 9, z: 4
5, w: OP, x: ADD, y: 9, z: 5
6, w: OP, x: ADD, y: 9, z: 6
7, w: OP, x: ADD, y: 9, z: 7
8, w: LOAD, x: 8, y: 3, z: 2			R8 = "50"

9, w: LOAD, x: 15, y: 4, z: 8			R15 = "72"
10, w: OP, x: XOR, y: 5, z: 15          R5 ^= R15
11, w: OP, x: ASG, y: 15, z: 7          R15 = R7
12, w: LOAD, x: 14, y: 0, z: 15         R14 = "15"
13, w: OP, x: AND, y: 15, z: 14         R15 &= R14
14, w: LOAD, x: 13, y: 0, z: 4          R13 = "4"
15, w: OP, x: SHL, y: 15, z: 13         R15 <<= R13
16, w: OP, x: SHR, y: 7, z: 13			R7 >>= R13
17, w: OP, x: AND, y: 7, z: 14          R7 &= R14
18, w: OP, x: OR, y: 7, z: 15           R7 |= R15
19, w: LOAD, x: 15, y: 0, z: 13         R15 = "13"
20, w: OP, x: MUL, y: 4, z: 15          R4 *= R15
21, w: OP, x: MIN, y: 6, z: 6			R6 = -R6
22, w: OP, x: ASG, y: 15, z: 3          R15 = R3
23, w: OP, x: ADD, y: 15, z: 7          R15 += R7
24, w: OP, x: XOR, y: 15, z: 1          R15 ^= R1
25, w: OP, x: ADD, y: 15, z: 4          R15 += R4
26, w: OP, x: XOR, y: 0, z: 15          R0 ^= R15
27, w: OP, x: NEG, y: 3, z: 3           R3 ~= R3
28, w: OP, x: ADD, y: 2, z: 0           R2 += R0

29, w: LOAD, x: 15, y: 8, z: 0			R15 = "128"
30, w: OP, x: CMP, y: 4, z: 15          CMP R4, R15
31, w: GOTO, x: 6, y: 2, z: 2           JMP <= "34"
32, w: LOAD, x: 15, y: 2, z: 3          R15 = "35"
33, w: OP, x: ADD, y: 1, z: 15          R1 += R15
34, w: OP, x: SWP, y: 0, z: 7           SWP R0, R7
35, w: OP, x: SWP, y: 1, z: 6           SWP R1, R6
36, w: OP, x: SWP, y: 2, z: 5           SWP R2, R5
37, w: OP, x: SWP, y: 3, z: 4           SWP R3, R4

38, w: LOAD, x: 15, y: 0, z: 1          R15 = "1"  LOOP
39, w: OP, x: SUB, y: 8, z: 15          R8 -= R15
40, w: LOAD, x: 15, y: 0, z: 0          R15 = "0"
41, w: OP, x: CMP, y: 8, z: 15          CPM R8, R15
42, w: GOTO, x: 2, y: 0, z: 9           JMP != "9"    LOOP 50 TIMES

43, w: LOAD, x: 15, y: 13, z: 11        R15 = "219"   CHECK PARAMS
44, w: OP, x: CMP, y: 0, z: 15          CPM R0, R15
45, w: GOTO, x: 2, y: 5, z: 6    ---->  JMP != 86 (WRONG)
46, w: LOAD, x: 15, y: 1, z: 2
47, w: OP, x: CMP, y: 1, z: 15
48, w: GOTO, x: 2, y: 5, z: 6    ---> All these lines check expected regiter values at the end
49, w: LOAD, x: 15, y: 11, z: 5
50, w: OP, x: CMP, y: 2, z: 15
51, w: GOTO, x: 2, y: 5, z: 6
52, w: LOAD, x: 15, y: 1, z: 4
53, w: OP, x: CMP, y: 3, z: 15
54, w: GOTO, x: 2, y: 5, z: 6
55, w: LOAD, x: 15, y: 2, z: 3
56, w: OP, x: CMP, y: 4, z: 15
57, w: GOTO, x: 2, y: 5, z: 6
58, w: LOAD, x: 15, y: 15, z: 4
59, w: OP, x: CMP, y: 5, z: 15
60, w: GOTO, x: 2, y: 5, z: 6
61, w: LOAD, x: 15, y: 3, z: 8
62, w: OP, x: CMP, y: 6, z: 15
63, w: GOTO, x: 2, y: 5, z: 6
64, w: LOAD, x: 15, y: 6, z: 8
65, w: OP, x: CMP, y: 7, z: 15
66, w: GOTO, x: 2, y: 5, z: 6
67, w: LOAD, x: 15, y: 9, z: 2   -->>> Right!  (all paramas as expected)
68, w: OP, x: XOR, y: 15, z: 9
69, w: WRITE, x: 15, y: 0, z: 0
70, w: LOAD, x: 15, y: 10, z: 9
71, w: OP, x: XOR, y: 15, z: 9
72, w: WRITE, x: 15, y: 0, z: 0
73, w: LOAD, x: 15, y: 10, z: 7
74, w: OP, x: XOR, y: 15, z: 9
75, w: WRITE, x: 15, y: 0, z: 0
76, w: LOAD, x: 15, y: 10, z: 8
77, w: OP, x: XOR, y: 15, z: 9
78, w: WRITE, x: 15, y: 0, z: 0
79, w: LOAD, x: 15, y: 11, z: 4
80, w: OP, x: XOR, y: 15, z: 9
81, w: WRITE, x: 15, y: 0, z: 0
82, w: LOAD, x: 15, y: 14, z: 1
83, w: OP, x: XOR, y: 15, z: 9
84, w: WRITE, x: 15, y: 0, z: 0
85, w: END, x: 15, y: 15, z: 15
86, w: LOAD, x: 9, y: 3, z: 12         --->> Wrong!
87, w: LOAD, x: 15, y: 6, z: 11
88, w: OP, x: XOR, y: 15, z: 9
89, w: WRITE, x: 15, y: 0, z: 0
90, w: LOAD, x: 15, y: 4, z: 14
91, w: OP, x: XOR, y: 15, z: 9
92, w: WRITE, x: 15, y: 0, z: 0
93, w: LOAD, x: 15, y: 5, z: 3
94, w: OP, x: XOR, y: 15, z: 9
95, w: WRITE, x: 15, y: 0, z: 0
96, w: LOAD, x: 15, y: 5, z: 2
97, w: OP, x: XOR, y: 15, z: 9
98, w: WRITE, x: 15, y: 0, z: 0
99, w: LOAD, x: 15, y: 5, z: 11
100, w: OP, x: XOR, y: 15, z: 9
101, w: WRITE, x: 15, y: 0, z: 0
102, w: LOAD, x: 15, y: 1, z: 13
103, w: OP, x: XOR, y: 15, z: 9
104, w: WRITE, x: 15, y: 0, z: 0
105, w: END, x: 15, y: 15, z: 15

 */

#include <stdio.h>
#include <iostream>
#include <boost/format.hpp>

using namespace std;


int main(int argc, char *argv[]) {
    cout << "4834636b7c643165" << endl;
    return 0;
}


inline void swap (uint8_t& a, uint8_t& b) {
    a ^= b;
    b ^= a;
    a ^= b;
}


uint8_t div13[256];

void init_div13()
{
    uint8_t i13 = 13;

    for (int ii=1; ii<256; ii++) {
        uint8_t i = ii;
        uint8_t mul = ii * i13;
        div13[mul] = i;
    }

}

uint8_t calc7[256];

void init_calc7()
{
    uint8_t r13 = 4;
    uint8_t r14 = 15;

    for (int ii=1; ii<256; ii++) {
        uint8_t i = ii;
        uint8_t r7 = i;

        uint8_t r15 = r7;
        r15 &= r14;
        r15 <<= r13;
        r7 >>= r13;
        r7 &= r14;
        r7 |= r15;

        calc7[r7] = i;
    }
}



void reverse_fn()
{
    uint8_t r0 = 219, r1 = 18, r2 = 181, r3 = 20, r4 = 35, r5 = 244, r6 = 56, r7 = 104;
    uint8_t r13 = 0, r14 = 0, r15 = 0;

    for (int i=0; i<50; i++) {


        swap(r0, r7);
        swap(r1, r6);
        swap(r2, r5);
        swap(r3, r4);

        r15 = 128;
        if (r4 > r15) {
            r15 = 35;
            r1 -= r15;
        }

        r3 = ~r3;
        r2 -= r0;

        r15 = r3;
        r15 += r7;
        r15 ^= r1;
        r15 += r4;
        r0 ^= r15;

        r6 = -r6;

        r4 = div13[r4];

        r7 = calc7[r7];

        r15 = 72;
        r5 ^= r15;


    }

    // 'H' '4' 'c' 'k' '|' 'd' '1' 'e'

}

bool original_fn(uint32_t h, uint32_t l)
{
    uint8_t r0 = 0, r1 = 0, r2 = 0, r3 = 0, r4 = 0, r5 = 0, r6 = 0, r7 = 0;
    uint8_t r13 = 0, r14 = 0, r15 = 0;


    r7 = l & 0xff;
    r6 = l >> 8 & 0xff;
    r5 = l >> 16 & 0xff;
    r4 = l >> 24;
    r3 = h & 0xff;
    r2 = h >> 8 & 0xff;
    r1 = h >> 16 & 0xff;
    r0 = h >> 24;


    r0 = ','; r1 = 'f'; r2 = 'm'; r3 = 'm'; r4 = '\xf4'; r5 = '+'; r6 = '\x94'; r7 = '\xa3';

    for (int i=0; i<50; i++) {
        r15 = 72;
        r5 ^= r15;
        r15 = r7;
        r14 = 15;
        r15 &= r14;
        r13 = 4;
        r15 <<= r13;
        r7 >>= r13;
        r7 &= r14;
        r7 |= r15;

        r15 = 13;
        r4 *= r15;
        r6 = -r6;
        r15 = r3;
        r15 += r7;
        r15 ^= r1;
        r15 += r4;
        r0 ^= r15;
        r3 = ~r3;
        r2 += r0;
        r15 = 128;
        if (r4 > r15) {
            r15 = 35;
            r1 += r15;
        }
        swap(r0, r7);
        swap(r1, r6);
        swap(r2, r5);
        swap(r3, r4);
    }


    return
            r0 == 219 &&
            r1 == 18 &&
            r2 == 181 &&
            r3 == 20 &&
            r4 == 35 &&
            r5 == 244 &&
            r6 == 56 &&
            r7 == 104;

}


uint8_t r[16], s[7], p;
uint16_t m[128] = {
        0x0090, 0x0391, 0x0392, 0x0393, 0x0394, 0x0395, 0x0396, 0x0397, 0x1832, 0x1f48,
        0x085f, 0x00f7, 0x1e0f, 0x06fe, 0x1d04, 0x0afd, 0x0b7d, 0x067e, 0x077f, 0x1f0d,
        0x054f, 0x0266, 0x00f3, 0x03f7, 0x08f1, 0x03f4, 0x080f, 0x0133, 0x0320, 0x1f80,
        0x0c4f, 0x2622, 0x1f23, 0x031f, 0x0d07, 0x0d16, 0x0d25, 0x0d34, 0x1f01, 0x048f,
        0x1f00, 0x0c8f, 0x2209, 0x1fdb, 0x0c0f, 0x2256, 0x1f12, 0x0c1f, 0x2256, 0x1fb5,
        0x0c2f, 0x2256, 0x1f14, 0x0c3f, 0x2256, 0x1f23, 0x0c4f, 0x2256, 0x1ff4, 0x0c5f,
        0x2256, 0x1f38, 0x0c6f, 0x2256, 0x1f68, 0x0c7f, 0x2256, 0x1f92, 0x08f9, 0x3f00,
        0x1fa9, 0x08f9, 0x3f00, 0x1fa7, 0x08f9, 0x3f00, 0x1fa8, 0x08f9, 0x3f00, 0x1fb4,
        0x08f9, 0x3f00, 0x1fe1, 0x08f9, 0x3f00, 0xffff, 0x193c, 0x1f6b, 0x08f9, 0x3f00,
        0x1f4e, 0x08f9, 0x3f00, 0x1f53, 0x08f9, 0x3f00, 0x1f52, 0x08f9, 0x3f00, 0x1f5b,
        0x08f9, 0x3f00, 0x1f1d, 0x08f9, 0x3f00, 0xffff
};


void goriginal() {
    uint16_t i;uint8_t w,x,y,z;
         /*
    r[0] = 'H';
    r[1] = '4';
    r[2] = 'c';
    r[3] = 'k';
    r[4] = '|';
    r[5] = 'd';
    r[6] = '1';
    r[7] = 'e';
     */

    int r0 = r[0], r1 = r[1], r2 = r[2], r3 = r[3], r4=r[4], r5=r[5], r6 = r[6], r7 = r[7];

    while(1) {
        i=m[p];
        w=i>>12;x=i>>8&0xf;y=i>>4&0xf;z=i&0xf;
        if(w==0) {
            if(x==0)r[y]=r[z];if(x==1)r[y]=~r[z];if(x==2)r[y]=-r[z];if(x==3)r[y]+=r[z];if(x==4)r[y]-=r[z];if(x==5)r[y]*=r[z];
            if(x==6)r[y]&=r[z];if(x==7)r[y]|=r[z];if(x==8)r[y]^=r[z];if(x==9)r[y]^=r[z];if(x==10)r[y]<<=r[z];if(x==11)r[y]>>=r[z];
            if(x==12){s[0]=1;s[1]=r[y]==r[z];s[2]=r[y]!=r[z];s[3]=r[y]>r[z];s[4]=r[y]>=r[z];s[5]=r[y]<r[z];s[6]=r[y]<=r[z];}
            if(x==13){x=r[y];r[y]=r[z];r[z]=x;}
        }
        if(w==1)r[x]=m[p];
        if(w==2&&s[x])p=m[p]-1;
        if(w==3){putchar(r[x]);}
        if(w==15){putchar(0xa);break;}
        p++;
    }
}


void g() {
    uint16_t i;
    uint8_t w, x, y, z;

    r[0] = 'H';
    r[1] = '4';
    r[2] = 'c';
    r[3] = 'k';
    r[4] = '|';
    r[5] = 'd';
    r[6] = '1';
    r[7] = 'e';

    while (1) {

        i = m[p];
        w = i >> 12;
        x = i >> 8 & 0xf;
        y = i >> 4 & 0xf;
        z = i & 0xf;

        int pp = p, ww = w, xx = x, yy = y, zz = z;

        int r0 = r[0];
        int r1 = r[1];
        int r2 = r[2];
        int r3 = r[3];
        int r4 = r[4];
        int r5 = r[5];
        int r6 = r[6];
        int r7 = r[7];
        int r8 = r[8];
        int r9 = r[9];
        int r10 = r[10];
        int r11 = r[11];
        int r12 = r[12];
        int r13 = r[13];
        int r14 = r[14];
        int r15 = r[15];

        if (p==43) {
    /*        r[0] = 219;
            r[1] = 18;
            r[2] = 181;
            r[3] = 20;
            r[4] = 35;
            r[5] = 244;
            r[6] = 56;
            r[7] = 104;

            r[9] = 192; // SUM OF ENTRIES MUST BE 192   // 0xC0 PRODUCES RIGHT    , original 0
     */
        }


        if (w == 0) {
            if (x == 0)r[y] = r[z];
            if (x == 1)r[y] = ~r[z];
            if (x == 2){
                if (p==21) {
                    int debug = 1;
                }
                r[y] = -r[z];
            }
            if (x == 3)r[y] += r[z];
            if (x == 4)r[y] -= r[z];
            if (x == 5)r[y] *= r[z];
            if (x == 6)r[y] &= r[z];
            if (x == 7)r[y] |= r[z];
            if (x == 8) {
                 r[y] ^= r[z];
            }
            if (x == 9){
                 r[y] ^= r[z];
            }
            if (x == 10)r[y] <<= r[z];
            if (x == 11)r[y] >>= r[z];
            if (x == 12) {
                s[0] = 1;
                s[1] = r[y] == r[z];
                s[2] = r[y] != r[z];
                s[3] = r[y] > r[z];
                s[4] = r[y] >= r[z];
                s[5] = r[y] < r[z];
                s[6] = r[y] <= r[z];
            }
            if (x == 13) {
                x = r[y];
                r[y] = r[z];
                r[z] = x;
            }
        }
        if (w == 1) {

            int aold = (uint8_t)r[x];
            int anew = (uint8_t)m[p];

            if (p==8) {
                int debug = 1;
            }

            r[x] = m[p];
        }
        if (w == 2) {
            if (s[x]   || p==45    ) {

                if (p==45) {
                    int debug = 10;
                }

                p = m[p] - 1;

                pp = p;
                cout << boost::format("JMP DONE, old: %1%, new: %2%\n") %  pp % (int)p;
            } else {
                cout << boost::format("NO JUMP, old: %1%\n") %  pp;
            }
        }
        if (w == 3) {
            putchar(r[x]);
        }
        if (w == 15) {
            putchar(0xa);
            break;
        }
        p++;
    }
}


void list() {

    string wstr[16];
    wstr[0] = "OP";
    wstr[1] = "LOAD";
    wstr[2] = "GOTO";
    wstr[3] = "WRITE";
    wstr[15] = "END";

    string xstr[16] = { "ASG", "NEG", "MIN", "ADD", "SUB", "MUL", "AND", "OR", "XOR", "XOR", "SHL", "SHR", "CMP", "SWP"   };


    for (int p=0; p<106; p++) {
        uint8_t w, x, y, z;
        uint16_t i = m[p];
        w = i >> 12;
        x = i >> 8 & 0xf;
        y = i >> 4 & 0xf;
        z = i & 0xf;
        cout << boost::format("%1%, w: %2%, x: %3%, y: %4%, z: %5%\n") % p % wstr[w] % ((w==0) ? xstr[x] : to_string(x)) % (int)y % (int)z;
    }

}



int mainold(int argc, char *argv[]) {

    init_div13();
    init_calc7();
  /*
    original_fn(0, 0);
    reverse_fn();

    for (uint32_t h=0; h < 0xffffffff; h++)
        for (uint32_t l=0; l < 0xffffffff; l++) {
            if (original_fn(h, l)) {
                cout << "FOUND: h: " << h << ", l: " << l << endl;
                return 0;
            }
            if (l % 1000000 ==0) cout << "h: " << h << ", l: " << l << endl;
        }
   // list();
   // return 0;
    */
    uint32_t h=0, l=0;

    // h = 0x4834636b;
    // l = 0x7c643165;

    scanf("%8x%8x", &h, &l);
    r[7] = l & 0xff;
    r[6] = l >> 8 & 0xff;
    r[5] = l >> 16 & 0xff;
    r[4] = l >> 24;
    r[3] = h & 0xff;
    r[2] = h >> 8 & 0xff;
    r[1] = h >> 16 & 0xff;
    r[0] = h >> 24;

    goriginal();
    return 0;
}

