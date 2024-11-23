package y23.d11;

import utilities.java.AocUtils;

import java.util.ArrayList;
import java.util.List;

public class Solution {
  private static final String DAY = "11";

  public static void part1(String inputPath) {
    List<String> input = AocUtils.readInputDataToList(inputPath);
    List<Point> galaxies = createGalaxies(input, 2);

    long result = 0;
    for (int i = 0; i < galaxies.size(); i++) {
      for (int g = i; g < galaxies.size(); g++) {
        result += galaxies.get(i).distanceToPoint(galaxies.get(g));
      }
    }

    System.out.println("Day " + DAY + " part 1 result: " + result);
  }

  private static List<Point> createGalaxies(List<String> input, int multiplier) {
    List<Point> galaxies = new ArrayList<>();

    long row = 0;
    for (String l1 : input) {
      long col = 0;
      boolean rowHasGalaxy = false;
      for (int c = 0; c < l1.length(); c++) {
        boolean colHasGalaxy = false;
        for (String l2 : input) {
          if (l2.charAt(c) == '#') {
            colHasGalaxy = true;
          }
        }
        if (l1.charAt(c) == '#') {
          rowHasGalaxy = true;
          galaxies.add(new Point(row, col));
        }
        col = colHasGalaxy ? col + 1 : col + multiplier;
      }
      row = rowHasGalaxy ? row + 1 : row + multiplier;
    }

    return galaxies;
  }

  public static void part2(String inputPath) {
    List<String> input = AocUtils.readInputDataToList(inputPath);
    List<Point> galaxies = createGalaxies(input, 1000000);

    long result = 0;
    for (int i = 0; i < galaxies.size(); i++) {
      for (int g = i; g < galaxies.size(); g++) {
        result += galaxies.get(i).distanceToPoint(galaxies.get(g));
      }
    }

    System.out.println("Day " + DAY + " part 2 result: " + result);
  }

  public static void main(String[] args) {
    part1("solutions/y23/d" + DAY + "/input.txt");
    part2("solutions/y23/d" + DAY + "/input.txt");
  }
}
