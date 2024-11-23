package y23.d05;

import utilities.java.AocUtils;

import java.time.LocalTime;
import java.util.*;
import java.util.regex.MatchResult;
import java.util.stream.IntStream;
import java.util.stream.LongStream;

public class Solution {

  public static List<List<List<Long>>> mappers;

  public static void part1(String inputPath) {
    List<String> input = AocUtils.readInputDataToList(inputPath);
    input.add("");

    List<Long> seeds = getSeeds(input.get(0));
    mappers = getMappers(input.subList(2, input.size()));

    long result = seeds.stream()
            .map(Solution::resolveMapping)
            .min(Long::compareTo)
            .orElseThrow();

    System.out.println("Part 1 result: " + result);
  }

  private static List<Long> getSeeds(String s) {
    return AocUtils.numberPattern.matcher(s)
            .results()
            .map(MatchResult::group)
            .map(Long::parseLong)
            .toList();
  }

  private static List<List<List<Long>>> getMappers(List<String> input) {
    List<List<List<Long>>> mappers = new ArrayList<>();

    List<List<Long>> map = new ArrayList<>();
    for (String line : input) {
      List<Long> range = AocUtils.numberPattern.matcher(line).results().map(MatchResult::group).map(Long::parseLong).toList();
      if (range.size() > 0) {
        map.add(range);
      } else if (line.isEmpty()) {
        mappers.add(map);
        map = new ArrayList<>();
      }
    }

    return mappers;
  }

  private static long resolveMapping(long i) {
    for (List<List<Long>> map : mappers) {
      long res = i;
      for (List<Long> range : map) {
        if (i >= range.get(1) && i < range.get(1) + range.get(2)) {
          res = range.get(0) + i - range.get(1);
        }
      }
      i = res;
    }
    return i;
  }

  public static void part2_1(String inputPath) {
    List<String> input = AocUtils.readInputDataToList(inputPath);
    input.add("");

    List<Long> seeds = getSeeds(input.get(0));
    List<long[]> seedRanges = IntStream.range(0, seeds.size()).filter(i -> i % 2 == 0).mapToObj(i -> new long[]{seeds.get(i), seeds.get(i + 1)}).toList();

    mappers = getMappers(input.subList(2, input.size()));

    long result = seedRanges.parallelStream()
            .flatMapToLong(r -> LongStream.range(r[0], r[0] + r[1]))
            .mapToObj(Solution::resolveMapping)
            .min(Long::compareTo)
            .orElseThrow();

    System.out.println("Part 2 result: " + result);
  }

  public static void part2_2(String inputPath) {
    List<String> input = AocUtils.readInputDataToList(inputPath);
    input.add("");

    mappers = getMappers(input.subList(2, input.size()));
    Collections.reverse(mappers);

    List<Long> seeds = getSeeds(input.get(0));
    List<long[]> seedRanges = IntStream.range(0, seeds.size()).filter(i -> i % 2 == 0).mapToObj(i -> new long[]{seeds.get(i), seeds.get(i + 1)}).toList();

    long max = seedRanges.stream().map(l -> l[0] + l[1]).max(Long::compareTo).orElseThrow();
    for (long i = 0; i < max; i++) {
      long seed = resolveMappingInRevers(i);
      boolean win = seedRanges.stream().anyMatch(r -> seed >= r[0] && seed < r[0] + r[1]);
      if (win) {
        System.out.println("Part 2 result: " + i);
        break;
      }
    }
  }

  private static long resolveMappingInRevers(long i) {
    for (List<List<Long>> map : mappers) {
      long res = i;
      for (List<Long> range : map) {
        if (i >= range.get(0) && i < range.get(0) + range.get(2)) {
          res = range.get(1) + i - range.get(0);
        }
      }
      i = res;
    }
    return i;

  }


  public static void main(String[] args) {
    System.out.println(LocalTime.now());
    part1("solutions/y23/d05/input.txt");
    System.out.println(LocalTime.now());
    part2_1("solutions/y23/d05/input.txt");
    System.out.println(LocalTime.now());
    part2_2("solutions/y23/d05/input.txt");
    System.out.println(LocalTime.now());

//    22:30:52.039407300
//    Part 1 result: 424490994
//    22:30:52.054407400
//    Part 2 result: 15290096
//    22:34:10.188402300
//    Part 2 result: 15290096
//    22:34:13.577056700
  }
}
