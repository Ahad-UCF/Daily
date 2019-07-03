// Ahad Bawany
// Daily code 7/3/2019:
// Given a list of numbers, return the smallest positive integer that isn't part of the list

public class Main
{
	int[] a;
	public int smallestNotPresent(int[] a)
	{
		// Set b to 1 since that's the smallest positive value possible
		int b = 1;
		for (int i = 0; i < a.length; i++)
		{
			// If the value within a is equal to our current b, that can't be the smallest possible value not present
			if (a[i] == b)
			{
				b++;
			}
		}
		return b;
	}

	public static void main (String args[])
	{
		Main s = new Main;
		s.a = new int[]{3, 4, -1, 1};
		System.out.println(s.smallestNotPresent(s.a));
	}
}
