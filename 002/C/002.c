#include <stdio.h>
 
int main()
{
        int x, y;
        x = 1;
        y = 0;
        x = x + y;
        y = x + x;
        printf ("%d %d\n", x, y );
        while (( x <= 4000000 ) && (y <= 4000000)) {
                x = x + y;
                y = x + y;
                printf( "%d %d\n", x, y);
        }
}
