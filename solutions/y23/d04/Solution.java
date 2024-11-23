package y23.d04;

import utilities.java.AocUtils;

import java.util.*;
import java.util.regex.MatchResult;
import java.util.regex.Pattern;
import java.util.stream.Stream;

public class Solution {

  public static void part1(String inputPath) {
    Stream<String> stream = AocUtils.readInputDataToStream(inputPath);
    int result = stream
            .map(Solution::getNumberOfWins)
            .map(i -> (int) Math.pow(2, i - 1))
            .reduce(0, Integer::sum);

    System.out.println("Part 1 result: " + result);
  }

  private static int getNumberOfWins(String s) {
    s = s.substring(s.indexOf(':'));
    String[] input = s.split("\\|");
    Pattern numberPattern = Pattern.compile("\\d+");

    List<Integer> winners = numberPattern.matcher(input[0])
            .results()
            .map(MatchResult::group)
            .map(Integer::parseInt)
            .toList();

    List<Integer> numbers = numberPattern.matcher(input[1])
            .results()
            .map(MatchResult::group)
            .map(Integer::parseInt)
            .toList();

    return (int) winners.stream().filter(numbers::contains).count();
  }

  public static void part2(String inputPath) {
    List<String> list = AocUtils.readInputDataToList(inputPath);
    int[] copies = new int[list.size()];
    Arrays.fill(copies, 1);

    for (int i = 0; i < list.size(); i++) {
      int nWins = getNumberOfWins(list.get(i));
      for (int copy = 0; copy < copies[i]; copy++) {
        addNewCopiesToCopies(copies, i + 1, nWins);
      }
    }

    int result = Arrays.stream(copies).reduce(0, Integer::sum);
    System.out.println("Part 2 result: " + result);
  }

  private static void addNewCopiesToCopies(int[] copies, int startIndex, int nWins) {
    for (int i = startIndex; i < startIndex + nWins; i++) {
      copies[i] = copies[i] + 1;
    }
  }

  public static void main(String[] args) {
    part1("solutions/y23/d04/input.txt");
    part2("solutions/y23/d04/input.txt");
  }
}
