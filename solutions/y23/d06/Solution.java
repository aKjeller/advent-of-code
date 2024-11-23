package y23.d06;

import utilities.java.AocUtils;

import java.util.List;
import java.util.regex.MatchResult;
import java.util.stream.IntStream;
import java.util.stream.LongStream;

public class Solution {

  public static void part1(String inputPath) {
    List<String> input = AocUtils.readInputDataToList(inputPath);

    List<Integer> times = AocUtils.numberPattern.matcher(input.get(0)).results().map(MatchResult::group).map(Integer::parseInt).toList();
    List<Integer> distances = AocUtils.numberPattern.matcher(input.get(1)).results().map(MatchResult::group).map(Integer::parseInt).toList();

    List<int[]> races = IntStream.range(0, times.size())
            .mapToObj(i -> new int[]{times.get(i), distances.get(i)})
            .toList();

    int result = races.parallelStream()
            .map(r -> (int) winsInRace(r[0], r[1]))
            .reduce(1, Math::multiplyExact);
    System.out.println("Part 1 result: " + result);
  }

  private static long winsInRace(long time, long distance) {
    return LongStream.range(0, time + 1).parallel()
            .map(i -> i * (time - i))
            .filter(t -> t > distance)
            .count();
  }

  public static void part2(String inputPath) {
    List<String> input = AocUtils.readInputDataToList(inputPath);

    long time = Long.parseLong(AocUtils.numberPattern.matcher(input.get(0)).results().map(MatchResult::group).reduce("", String::concat));
    long distance = Long.parseLong(AocUtils.numberPattern.matcher(input.get(1)).results().map(MatchResult::group).reduce("", String::concat));

    long result = winsInRace(time, distance);
    System.out.println("Part 2 result: " + result);
  }

  public static void main(String[] args) {
    part1("solutions/y23/d06/input.txt");
    part2("solutions/y23/d06/input.txt");
  }
}
