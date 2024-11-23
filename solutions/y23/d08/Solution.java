package y23.d08;

import utilities.java.AocUtils;

import java.math.BigInteger;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class Solution {

  private static final String DAY = "08";
  private static String path;
  private static final Map<String, String[]> map = new HashMap<>();

  public static void part1(String inputPath) {
    List<String> input = AocUtils.readInputDataToList(inputPath);
    String path = input.get(0);

    Map<String, String[]> map = new HashMap<>();
    for (String element : input.subList(2, input.size())) {
      String key = element.substring(0, 3);
      String L = element.substring(7, 10);
      String R = element.substring(12, 15);
      map.put(key, new String[]{L, R});
    }

    String currentLocation = "AAA";

    int index = 0;
    int result = 0;
    while (!currentLocation.equals("ZZZ")) {
      if (index == path.length()) {
        index = 0;
      }

      if (path.charAt(index) == 'L') {
        currentLocation = map.get(currentLocation)[0];
      } else {
        currentLocation = map.get(currentLocation)[1];
      }

      index++;
      result++;
    }

    System.out.println("Day " + DAY + " part 1 result: " + result);
  }

  public static void part2(String inputPath) {
    List<String> input = AocUtils.readInputDataToList(inputPath);
    path = input.get(0);

    for (String element : input.subList(2, input.size())) {
      String key = element.substring(0, 3);
      String L = element.substring(7, 10);
      String R = element.substring(12, 15);
      map.put(key, new String[]{L, R});
    }

    var result = input.stream()
            .skip(2)
            .filter(s -> (s.charAt(2) == 'A'))
            .map(s -> s.substring(0, 3))
            .map(Solution::cycleLength)
            .reduce(1L, (a, b) -> lcm(BigInteger.valueOf(a), BigInteger.valueOf(b)));

    System.out.println("Day " + DAY + " part 2 result: " + result);
  }

  private static long cycleLength(String startLocation) {
    String currentLocation = startLocation;
    int index = 0;
    long toZ = 0;
    while (!currentLocation.endsWith("Z")) {
      if (index == path.length()) {
        index = 0;
      }

      if (path.charAt(index) == 'L') {
        currentLocation = map.get(currentLocation)[0];
      } else {
        currentLocation = map.get(currentLocation)[1];
      }

      index++;
      toZ++;
    }

    return toZ ;
  }

  public static long lcm(BigInteger number1, BigInteger number2) {
    BigInteger gcd = number1.gcd(number2);
    BigInteger absProduct = number1.multiply(number2).abs();
    return absProduct.divide(gcd).longValue();
  }

  public static void main(String[] args) {
    //part1("solutions/y23/d" + DAY + "/input.txt");
    part2("solutions/y23/d" + DAY + "/input.txt");
  }
}
