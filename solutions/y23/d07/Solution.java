package y23.d07;

import utilities.java.AocUtils;

import java.util.stream.IntStream;
import java.util.stream.Stream;

public class Solution {

  private static final String DAY = "07";

  public static void part1(String inputPath) {
    Stream<String> input = AocUtils.readInputDataToStream(inputPath);

    var hands = input
            .map(s -> s.split(" "))
            .map(s -> new Hand(s[0], Integer.parseInt(s[1])))
            .sorted()
            .toList();

    long result = IntStream.range(0, hands.size())
            .map(i -> (i + 1) * hands.get(i).getBid())
            .mapToObj(Long::valueOf)
            .reduce(0L, Long::sum);

    System.out.println("Day " + DAY + " part 1 result: " + result);
  }

  public static void part2(String inputPath) {
    Stream<String> input = AocUtils.readInputDataToStream(inputPath);

    var hands = input
            .map(s -> s.split(" "))
            .map(s -> new Hand2(s[0], Integer.parseInt(s[1])))
            .sorted()
            .toList();

    long result = IntStream.range(0, hands.size())
            .map(i -> (i + 1) * hands.get(i).getBid())
            .mapToObj(Long::valueOf)
            .reduce(0L, Long::sum);

    System.out.println("Day " + DAY + " part 1 result: " + result);
  }

  public static void main(String[] args) {
    part1("solutions/y23/d" + DAY + "/input.txt");
    part2("solutions/y23/d" + DAY + "/input.txt");
  }
}
