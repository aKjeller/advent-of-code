package y23.d01;

import utilities.java.AocUtils;

import java.util.List;
import java.util.stream.Stream;

public class Solution {

  public static List<String> wordList = Stream.of("one", "two", "three", "four", "five", "six", "seven", "eight", "nine").toList();
  public static void part1(String inputPath) {
    Stream<String> stream = AocUtils.readInputDataToStream(inputPath);
    int result = stream
            .map(Solution::getDigitsFromString)
            .reduce(0, Integer::sum);
    System.out.println("Part 1 result: " + result);
  }

  private static int getDigitsFromString(String s) {
    String allNumbers = "";
    for (char c : s.toCharArray()) {
      try {
        Integer.parseInt(String.valueOf(c));
        allNumbers += c;
      } catch (NumberFormatException ignored) { }
    }

    String result = allNumbers.substring(0, 1);
    result += allNumbers.substring(allNumbers.length() - 1);

    return Integer.parseInt(result);
  }

  public static void part2(String inputPath) {
    Stream<String> stream = AocUtils.readInputDataToStream(inputPath);
    int result = stream
            .map(Solution::replaceWordsWithNumbers)
            .map(Solution::getDigitsFromString)
            .reduce(0, Integer::sum);
    System.out.println("Part 2 result: " + result);
  }

  private static String replaceWordsWithNumbers(String s) {
    for (int i = 0; i < s.length(); i++) {
        for (int w = 0; w < wordList.size(); w++) {
          if (s.substring(i).startsWith(wordList.get(w))) {
            String result = s.substring(0, i);
            result += String.valueOf(w + 1);
            result += s.substring(i + 1);
            s = result;
          }
      }

    }
    return s;
  }


  public static void main(String[] args) {
    part1("solutions/y23/d01/input.txt");
    part2("solutions/y23/d01/input.txt");
  }
}
