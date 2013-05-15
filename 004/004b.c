#include <stdio.h>
#include <time.h>

int main()
{
    clock_t start = clock();
    int a, b, c, d;
    int n, rev = 0, temp;

    for (a=999; a<=100; a--){
        for (b=a; b<=100; b--){
            n = a*b;    
            temp = n;
            while(temp != 0){
                rev = rev * 10;
                rev = rev + temp%10;
                temp = temp/10; 
            }
                if (n==rev){
                    c = n;
                    if (c < d){
                       d = c;
                    }
                }
                rev = 0;  
            }
    }
    printf("Largest palindrome >> %d\n", d);
    printf("Elapsed Time >> %f seconds.\n", ((double)clock() - start) / CLOCKS_PER_SEC); 
}