#include <stdio.h>

// 1 1 [2] 3 5 [8] 13 21 [34] ...
// a b [c] a b [c] a b [c] ...

int main()
{
	int a=1, b=1, c, d=0;
	while (c < 4000000){
		c = a + b; 
		a = c + b; 
		b = a + c; 
	if (c < 4000000){
		d += c;
		}
	}
	printf("d: %d\n", d);	
}
