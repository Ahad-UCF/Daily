// Ahad Bawany
// Daily code 6/30:
// Given a list of numbers and a number key, print whether any two numbers from the list add up
// to the key

#include <stdio.h>
#include <stdlib.h>

void oneSum(int key, int arr[], int len)
{

	// Check each value at each index individually
	for (int i = 0; i < len; i++)
	{
		// Start at the next index as we can't add a number to itself
		for (int j = i+1; j < len; j++)
		{
			// If the current value + any of the next values is the key, print true
			if (arr[i] + arr[j] == key)
			{
				printf("True\n");
				return;
			}
		}
	}

}

void main()
{
	int arr[] = {10, 5, 3, 16};
	int len = 4;
	oneSum(17, arr, 4);
}
