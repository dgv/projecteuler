#include <stdio.h>

void main(){
    long a=600851475143, b, c;
    for (b=2; b<=a; b++){
        if (a%b==0) {
        a = a/b;
        c = b;
        b--;
        if(a==1)
            break;
        }
    }
    printf("Largest one: %d\n", c);    
}