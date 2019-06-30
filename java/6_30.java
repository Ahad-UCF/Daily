// Ahad Bawany
// Daily code 6/30:
// Given a list of numbers and a number key, print whether any two numbers from the list add up
// to the key

import java.util.HashMap;
import java.util.Map;

public class Main
{
	// Create a field of a HashMap called map
	HashMap<Integer, Integer> map;

  public void oneSum(Map<Integer,Integer> map, int key)
	{
		// Cycle through all the keys in the Map
		for (Object k: map.keySet())
		{
			// key - value at k is equal to the number we need to add to a value to get the key
			// each index is equal to its value, so if an index exists, that value exists
			if (map.containsKey(key - map.get(k)))
			{
				// If the value exists, we know that the map contains a value which would work!
				System.out.println("True");
        return;
			}
		}
	}

	public static void main(String[] args)
	{
		// We need to create an instance of main to use any of its methods
		Main m = new Main();
    m.map = new HashMap<>();
		m.map.put(10,10);
		m.map.put(7,7);
		m.map.put(15,15);
		m.map.put(3,3);
		m.oneSum(m.map, 17);
	}
}
