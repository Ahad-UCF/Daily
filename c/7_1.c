// Ahad Bawany
// Daily code 7/1/2019:
// Given a list of numbers, return a list of numbers where each number is the result of multiplying
// all other numbers in the list

#include <stdio.h>
#include <stdlib.h>

int * multiplArr(int arr[], int len)
{
	// Initialize our counters
	int i = 0;
	int j = 0;

	// Allocate memory for our array which we will return
	int *arr2= malloc(sizeof(int)*len);

	// Set the values of the entire array to 1
	for (i = 0; i < len; i++)
	{
		arr2[i] = 1;
	}

	// Cycle through the second (results) array
	for (i = 0; i < len; i++)
	{
		// Cycle through each element of the first (initial) array
		for (j = 0; j < len; j++)
		{
			if (i != j)
			{
				// Multiply each element EXCEPT for the index that's the same in both arrays
				arr2[i] *= arr[j];
			}
		}
	}

	return arr2;
}

void main()
{
	int arr[] = {3, 2, 1};
	int *arr2= multiplArr(arr, 3);
	int i = 0;

	// After generating the answers array, print it out
	for (i = 0; i < 3; i++)
	{
		printf("%d ", arr2[i]);
	}

	// Makes a new line
	puts("");
}
