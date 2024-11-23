package y23.d02;

import utilities.java.AocUtils;

import java.util.List;
import java.util.stream.Stream;

public class Solution {

  public static void part1(String inputPath) {
    Stream<String> stream = AocUtils.readInputDataToStream(inputPath);
    int result = stream
            .filter(Solution::gameIsPossible)
            .map(Solution::getGameID)
            .reduce(0, Integer::sum);
    System.out.println("Part 1 result: " + result);
  }

  private static int getGameID(String s) {
    return Integer.parseInt(s.substring(s.indexOf(' ') + 1, s.indexOf(':')));
  }

  private static boolean gameIsPossible(String s) {
    s = s.substring(s.indexOf(':') + 1);
    List<String> rounds = List.of(s.split(";"));
    for (String round : rounds) {
      List<String> balls = List.of(round.split(","));
      for (String ball : balls) {
        if (!ballNumberAllowed(ball)) {
          return false;
        }
      }
    }
    return true;
  }

  private static boolean ballNumberAllowed(String ball) {
    // only 12 red cubes, 13 green cubes, and 14 blue cubes?
    if (ball.endsWith("red")) {
      return Integer.parseInt(ball.substring(1, ball.length() - 4)) <= 12;
    } else if (ball.endsWith("green")) {
      return Integer.parseInt(ball.substring(1, ball.length() - 6)) <= 13;
    } else if (ball.endsWith("blue")) {
      return Integer.parseInt(ball.substring(1, ball.length() - 5)) <= 14;
    }
    return true;
  }


  public static void part2(String inputPath) {
    Stream<String> stream = AocUtils.readInputDataToStream(inputPath);
    int result = stream
            .map(Solution::getGamePower)
            .reduce(0, Integer::sum);
    System.out.println("Part 2 result: " + result);
  }

  private static int getGamePower(String s) {
    s = s.substring(s.indexOf(':') + 1);

    int maxRed = 0;
    int maxBlue = 0;
    int maxGreen = 0;

    List<String> rounds = List.of(s.split(";"));
    for (String round : rounds) {
      List<String> balls = List.of(round.split(","));
      for (String ball : balls) {
        if (ball.endsWith("red")) {
          maxRed = Math.max(Integer.parseInt(ball.substring(1, ball.length() - 4)), maxRed);
        } else if (ball.endsWith("green")) {
          maxGreen = Math.max(Integer.parseInt(ball.substring(1, ball.length() - 6)), maxGreen);
        } else if (ball.endsWith("blue")) {
          maxBlue = Math.max(Integer.parseInt(ball.substring(1, ball.length() - 5)), maxBlue);
        }
      }
    }
    return maxRed * maxGreen * maxBlue;
  }


  public static void main(String[] args) {
    part1("solutions/y23/d02/input.txt");
    part2("solutions/y23/d02/input.txt");
  }
}
