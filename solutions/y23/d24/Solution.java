package y23.d24;

import utilities.java.AocUtils;

import java.util.*;

public class Solution {
  private static final String DAY = "24";

  public static void part1(String inputPath) {
    List<String> input = AocUtils.readInputDataToList(inputPath);

    List<Hail> hails = new ArrayList<>();
    for (String line : input) {
      hails.add(new Hail(line));
    }

    long result = 0;
    for (int i = 0; i < hails.size(); i++) {
      for (int j = i + 1; j < hails.size(); j++){
        if (hails.get(i).intersects(hails.get(j), 200000000000000L, 400000000000000L)) {
          result += 1;
        }
      }
    }

    System.out.println("Day " + DAY + " part 1 result: " + result);
  }

  public static void part2(String inputPath) {
    List<String> input = AocUtils.readInputDataToList(inputPath);

    long result = 0;
    System.out.println("Day " + DAY + " part 2 result: " + result);
  }

  public static void main(String[] args) {
    part1("solutions/y23/d" + DAY + "/input.txt");
    part2("solutions/y23/d" + DAY + "/input.txt");
  }
}
