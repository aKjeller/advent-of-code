package y23.d13;

import utilities.java.AocUtils;

import java.util.ArrayList;
import java.util.List;

public class Solution {
  private static final String DAY = "13";

  public static void part1(String inputPath) {
    List<String> input = AocUtils.readInputDataToList(inputPath);
    List<char[][]> maps = createMaps(input);

    long result = 0;
    for (char[][] map : maps) {
      result += getVerticalMirros(map);
      result += getHorizontalMirros(map);
    }

    System.out.println("Day " + DAY + " part 1 result: " + result);
  }

  private static long getVerticalMirros(char[][] map) {
    long best = 0;
    for (int col = 0; col < map[0].length - 1; col++) {
     boolean matches = true;
     for (int mirrorSize = 1; mirrorSize <= Math.min(col + 1, map[0].length - col - 1); mirrorSize++) {
       int leftIndex = col - mirrorSize + 1;
       int rightIndex = col + mirrorSize;
       for (int row = 0; row < map.length; row++) {
         if (map[row][leftIndex] != map[row][rightIndex]) {
           matches = false;
         }
       }
     }
      if (matches) {
        best = Math.max(col + 1, best);
      }
    }

    return best;
  }

  private static long getHorizontalMirros(char[][] map) {
    long best = 0;
    for (int row = 0; row < map.length - 1; row++) {
      boolean matches = true;
      for (int mirrorSize = 1; mirrorSize <= Math.min(row + 1, map.length - row - 1); mirrorSize++) {
        int leftIndex = row - mirrorSize + 1;
        int rightIndex = row + mirrorSize;
        for (int col = 0; col < map[0].length; col++) {
          if (map[leftIndex][col] != map[rightIndex][col]) {
            matches = false;
          }
        }
      }
      if (matches) {
        best = Math.max((row + 1) * 100, best);
      }
    }

    return best;
  }

  public static void part2(String inputPath) {
    List<String> input = AocUtils.readInputDataToList(inputPath);
    List<char[][]> maps = createMaps(input);

    long result = 0;
    for (char[][] map : maps) {
      result += calculateNewMap(map);
    }

    System.out.println("Day " + DAY + " part 2 result: " + result);
  }

  private static long calculateNewMap(char[][] map) {
    List<int[]> rowCols = new ArrayList<>();
    long result = 0L;
    for (int row = 0; row < map.length; row++) {
      for (int col = 0; col < map[0].length; col++) {
        long first = getVerticalMirros(map);
        map[row][col] = getOpposite(map[row][col]);
        long second = getVerticalMirros(map, first);
        if (first != second && second > 0) {
          result = second;
          rowCols.add(new int[]{row, col});
        }
        map[row][col] = getOpposite(map[row][col]);

        first = getHorizontalMirros(map);
        map[row][col] = getOpposite(map[row][col]);
        second = getHorizontalMirros(map, first);
        if (first != second && second > 0) {
          result = second;
          rowCols.add(new int[]{row, col});
        }
        map[row][col] = getOpposite(map[row][col]);
      }
    }
    return result;
  }

  private static long getVerticalMirros(char[][] map, long first) {
    for (int col = 0; col < map[0].length - 1; col++) {
      boolean matches = true;
      for (int mirrorSize = 1; mirrorSize <= Math.min(col + 1, map[0].length - col - 1); mirrorSize++) {
        int leftIndex = col - mirrorSize + 1;
        int rightIndex = col + mirrorSize;
        for (int row = 0; row < map.length; row++) {
          if (map[row][leftIndex] != map[row][rightIndex]) {
            matches = false;
          }
        }
      }
      if (matches) {
        long newResult = col + 1;
        if (newResult != first) {
          return newResult;
        }
      }
    }

    return 0;
  }

  private static long getHorizontalMirros(char[][] map, long first) {
    for (int row = 0; row < map.length - 1; row++) {
      boolean matches = true;
      for (int mirrorSize = 1; mirrorSize <= Math.min(row + 1, map.length - row - 1); mirrorSize++) {
        int leftIndex = row - mirrorSize + 1;
        int rightIndex = row + mirrorSize;
        for (int col = 0; col < map[0].length; col++) {
          if (map[leftIndex][col] != map[rightIndex][col]) {
            matches = false;
          }
        }
      }
      if (matches) {
        long newResult = (row + 1) * 100;
        if (newResult != first) {
          return newResult;
        }
      }
    }

    return 0;
  }

  private static char getOpposite(char c) {
    return switch (c) {
      case '.' -> '#';
      case '#' -> '.';
      default -> throw new IllegalStateException("Unexpected value: " + c);
    };
  }

  private static List<char[][]> createMaps(List<String> input) {
    input.add("");

    List<char[][]> maps = new ArrayList<>();
    List<String> map = new ArrayList<>();
    for (String line : input) {
      if (line.isBlank()) {
        char[][] mapArray = new char[map.size()][map.get(0).length()];
        for (int i = 0; i < map.size(); i++) {
          mapArray[i] = map.get(i).toCharArray();
        }
        maps.add(mapArray);
        map = new ArrayList<>();
      } else {
        map.add(line);
      }
    }

      return maps;
  }

  public static void main(String[] args) {
    part1("solutions/y23/d" + DAY + "/input.txt");
    part2("solutions/y23/d" + DAY + "/input.txt");
  }
}
