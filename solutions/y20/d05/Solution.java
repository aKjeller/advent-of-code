package y20.d05;

import utilities.java.AocUtils;

import java.util.List;

public class Solution {
  private static final String DAY = "01";

  public static void part1(String inputPath) {
    List<Long> input = AocUtils.readInputDataToStream(inputPath).map(Long::parseLong).toList();

    long result = 0;
    for(int i = 0; i < input.size(); i++) {
      for (int j = i; j < input.size(); j++) {
        if (input.get(i) + input.get(j) == 2020) {
          result = input.get(i) * input.get(j);
        }
      }
    }

    System.out.println("Part 1 result: " + result);
  }

  public static void part2(String inputPath) {
    List<Long> input = AocUtils.readInputDataToStream(inputPath).map(Long::parseLong).toList();

    long result = 0;
    for(int i = 0; i < input.size(); i++) {
      for (int j = i; j < input.size(); j++) {
        for (int k = j; k < input.size(); k++) {
          if (input.get(i) + input.get(j) + input.get(k) == 2020) {
            result = input.get(i) * input.get(j) * input.get(k);
          }
        }
      }
    }

    System.out.println("Part 2 result: " + result);
  }



  public static void main(String[] args) {
    part1("solutions/y20/d" + DAY + "/input.txt");
    part2("solutions/y20/d" + DAY + "/input.txt");
  }
}
