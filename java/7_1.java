// Ahad Bawany
// Daily code 7/1/2019:
// Given a list of numbers, return a list of numbers where each number is the result of multiplying
// all other numbers in the list

public class Main
{
	public int[] multiplArr(int[] a)
	{
    int[] b = new int[a.length];

		// Set the values of the entire array to 1
    for (int i = 0; i < a.length; i++)
    {
      b[i] = 1;
    }

		// Cycle through the answers array
    for (int i = 0; i < a.length; i ++)
    {
			// Cycle through the first array
      for (int j = 0; j < a.length; j++)
      {
        if (i != j)
        {
					// Multiply by all the elements of the first array EXCEPT the index of the answers array
          b[i] *= a[j];
        }
      }
    }
		return b;
	}
	public static void main(String[] args)
	{
    Main ob = new Main();
    int[] first = {3, 2, 1};
    int[] answers = ob.multiplArr(first);

		// Now that the answers array exists, we can print it!
    for (int i = 0; i < first.length; i++)
    {
      System.out.println(answers[i]);
    }
	}
}
