#include <stdio.h>

int main(){
    int x=0, y, w=0, z=0, m, t;
    while (x<=100) {
        y = x * x;
        x++;
        w = w + y;
        if (x<101) {
            z = z + x;
            m = z * z;
        }
        t = m - w;
    }
    printf("Difference: %d\n", t);
}