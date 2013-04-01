#include <stdio.h>

int main()
{	
	int total = 0;
	int n = 0;

	while (n < 1000)
	{
		if (n%3 == 0 || n%5 == 0)
			total += n;
		n++;
	}
	printf("Sum = %d \n", total);
}
