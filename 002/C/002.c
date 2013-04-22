#include <stdio.h>

// 1 1 [2] 3 5 [8] 13 21 [34] 55 89 [144] 233 377 [610] ...

int main()
{
 int a, b, c;
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
}
