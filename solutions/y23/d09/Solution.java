package y23.d09;

import utilities.java.AocUtils;

import java.util.*;
import java.util.stream.IntStream;

public class Solution {
  private static final String DAY = "09";

  public static void part1(String inputPath) {
    List<String> input = AocUtils.readInputDataToList(inputPath);

    long result = 0;
    for (String line : input) {
      List<Long> firstRow = AocUtils.createListOfLongsFromString(line);
      long[][] ranges = new long[firstRow.size() + 1][firstRow.size() + 1];
      IntStream.range(0, firstRow.size()).forEach(i -> ranges[0][i] = firstRow.get(i));
      calculateMap(ranges, firstRow.size());
      long n = calculateNumber(ranges, 0, firstRow.size());
      result += n;

    }
    System.out.println("Day " + DAY + " part 1 result: " + result);
  }

  private static void calculateMap(long[][] ranges, int size) {
    for (int row = 1; row < size; row++) {
      for (int col = 0; col < size - row; col++) {
        ranges[row][col] = ranges[row - 1][col + 1] - ranges[row - 1][col];
      }
    }
  }

  private static long calculateNumber(long[][] ranges, int row, int col) {
    if (col < 1) {
      return 0;
    }
    ranges[row][col] =  ranges[row][col - 1] + calculateNumber(ranges, row + 1, col - 1);
    return ranges[row][col];
  }

  public static void part2(String inputPath) {
    List<String> input = AocUtils.readInputDataToList(inputPath);

    long result = 0;
    for (String line : input) {
      List<Long> firstRow = AocUtils.createListOfLongsFromString(line);
      long[][] ranges = new long[firstRow.size() + 1][firstRow.size() + 1];
      IntStream.range(0, firstRow.size()).forEach(i -> ranges[0][i] = firstRow.get(i));
      calculateMap(ranges, firstRow.size());
      long n = calculateNumberBackwards(ranges, 0, firstRow.size());
      result += n;

    }
    System.out.println("Day " + DAY + " part 2 result: " + result);
  }

  private static long calculateNumberBackwards(long[][] ranges, int row, int size) {
    if (row > size - 1) {
      return 0;
    }
    return ranges[row][0] - calculateNumberBackwards(ranges, row + 1, size);
  }

  public static void main(String[] args) {
    part1("solutions/y23/d" + DAY + "/input.txt");
    part2("solutions/y23/d" + DAY + "/input.txt");
  }
}
