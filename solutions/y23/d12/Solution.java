package y23.d12;

import utilities.java.AocUtils;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.regex.MatchResult;
import java.util.regex.Pattern;
import java.util.stream.Stream;

public class Solution {
  private static final String DAY = "12";
  private static final Map<Input, Long> memo = new HashMap<>();


  public static void part1(String inputPath) {
    List<String[]> input = AocUtils.readInputDataToStream(inputPath).map(s -> s.split(" ")).toList();

    long result = 0;
    for (String[] row : input) {
      String springConditions = row[0] + ".";
      List<Integer> springGroups = AocUtils.numberPattern.matcher(row[1]).results().map(MatchResult::group).map(Integer::parseInt).toList();
      Pattern pattern = createPattern(springGroups);
      result += magic(pattern, springConditions);
    }

    System.out.println("Day " + DAY + " part 1 result: " + result);
  }

  private static int magic(Pattern pattern, String springConditions) {
    if (!springConditions.contains("?")) {
      return 1L == pattern.matcher(springConditions).results().count() ? 1 : 0;
    }
    char[] row = springConditions.toCharArray();
    for (int i = 0; i < row.length; i++) {
      if (row[i] == '?') {
        row[i] = '#';
        int matches = magic(pattern, String.valueOf(row));
        row[i] = '.';
        return matches + magic(pattern, String.valueOf(row));
      }
    }
    throw new RuntimeException();
  }

  private static Pattern createPattern(List<Integer> values) {
    String pattern = "^\\.*";

    for (Integer i : values) {
      pattern += "#{" + i + "}\\.{1,}";
    }
    pattern += "$";

    return Pattern.compile(pattern);
  }

  public static void part2(String inputPath) {
    List<String[]> input = AocUtils.readInputDataToStream(inputPath).map(s -> s.split(" ")).toList();
    long result = 0;
    for (String[] row : input) {
      String springs = Stream.generate(() -> row[0]).limit(4).reduce(row[0], (a, b) -> a + "?" + b);
      List<Integer> groups = Stream.generate(() -> AocUtils.createListOfIntegersFromString(row[1])).limit(5).collect(ArrayList::new, List::addAll, List::addAll);
      result += getVariants(springs, groups);
    }
    System.out.println("Day " + DAY + " part 2 result: " + result);
  }

  private static long getVariants(String springs, List<Integer> groups) {
    Long memoValue = memo.get(new Input(springs, groups));
    if (memoValue != null) {
      return memoValue;
    }

    if (springs.isBlank()) {
      return groups.isEmpty() ? 1 : 0;
    }

    long variants = switch (springs.charAt(0)) {
      case '.' -> getVariants(springs.substring(1), groups);
      case '?' -> getVariants("." + springs.substring(1), groups) + getVariants("#" + springs.substring(1), groups);
      case '#' -> validVariant(springs, groups);
      default -> throw new IllegalStateException("Unexpected value: " + springs.charAt(0));
    };

    memo.put(new Input(springs, groups), variants);
    return variants;
  }

  private static long validVariant(String springs, List<Integer> groups) {
    if (groups.isEmpty()) {
      return 0;
    }

    int groupSize = groups.get(0);
    if (springStartWithGroup(springs, groupSize)) {
      List<Integer> newGroups = groups.subList(1, groups.size());
      if (groupSize == springs.length()) {
        return newGroups.isEmpty() ? 1 : 0;
      } else if (springs.charAt(groupSize) != '#') {
        return getVariants("." + springs.substring(groupSize + 1), newGroups);
      }
    }
    return 0;
  }

  private static boolean springStartWithGroup(String springs, int groupSize) {
    return groupSize <= springs.length() && springs.chars().limit(groupSize).allMatch(c -> c == '#' || c == '?');
  }

  public record Input(String springs, List<Integer> groups) { }

  public static void main(String[] args) {
    part1("solutions/y23/d" + DAY + "/input.txt");
    part2("solutions/y23/d" + DAY + "/input.txt");
  }
}
