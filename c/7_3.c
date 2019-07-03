// Ahad Bawany
// Daily code 7/3/2019:
// Given a list of numbers, return the smallest positive integer that isn't part of the list

#include <stdio.h>
#include <stdlib.h>

int smallestNotPresent(int* a, int len)
{ 
	// Set b to 1 since that's the smallest positive value possible
	int b = 1;
	for (int i = 0; i < len; i++)
	{
		if (a[i] == b)
		{
			// If the value within a is equal to our current b, that can't be the smallest possible value not present
			b++;
		}
	}
	return b;
}

void main()
{
	int *a= malloc(sizeof(int)*4);
	a[0] = 3;
	a[1] = -1;
	a[2] = 0;
	a[3] = 1;
	printf("%d\n", smallestNotPresent(a, 4));
	free(a);
}
