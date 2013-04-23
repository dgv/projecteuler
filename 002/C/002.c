#include <stdio.h>

// 1 1 [2] 3 5 [8] 13 21 [34] 55 89 [144] 233 377 [610] ...

int main()
{
<<<<<<< HEAD
	int a, b, c;
=======
 int a, b, c;
>>>>>>> 0f5ed305be53fd392ca629d2a1c31dead75796b7
	a = 1;
	b = 1;
	c = 0;
	c = a + b;

	while (c <= 4000000) {
		a = c + b;
		b = a + c;
		c = a + b;

		if (c < 4000000) {
			printf("%d \n", c);	
		}
	}
<<<<<<< HEAD
}
=======
}
>>>>>>> 0f5ed305be53fd392ca629d2a1c31dead75796b7
