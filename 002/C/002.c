#include <stdio.h>

// 1 1 [2] 3 5 [8] 13 21 [34] ...
// a b [c] a b [c] a b [c] ...

int main()
{
	int a=1, b=1, c, d=0;
	while (c < 4000000){
<<<<<<< HEAD
		c = a + b; // 2
		a = c + b; // 3
		b = a + c; // 5
=======
		c = a + b; 
		a = c + b; 
		b = a + c; 
>>>>>>> 3d5c1408fb7613b64955e1bda9a22c51d9ae0624
	if (c < 4000000){
		d += c;
		}
	}
	printf("d: %d\n", d);	
}
<<<<<<< HEAD

=======
>>>>>>> 3d5c1408fb7613b64955e1bda9a22c51d9ae0624
