package y23.d14;

import utilities.java.AocUtils;

import java.util.*;

public class Solution {
  private static final String DAY = "14";

  public static void part1(String inputPath) {
    List<String> input = AocUtils.readInputDataToList(inputPath);
    char[][] dish = createDish(input);

    // Gravity
    for (int row = 1; row < input.size(); row++) {
      for (int col = 0; col < input.get(row).length(); col++) {
        if (dish[row][col] == 'O') {
          for (int newRow = row; newRow > 0; newRow--) {
            if (dish[newRow][col] == 'O' && dish[newRow - 1][col] == '.') {
              dish[newRow][col] = '.';
              dish[newRow - 1][col] = 'O';
            }
          }
        }
      }
    }

    // weight
    long result = 0;
    for (int row = 0; row < input.size(); row++) {
      for (int col = 0; col < input.get(row).length(); col++) {
        if (dish[row][col] == 'O') {
          result += input.size() - row;
        }
      }
    }

    System.out.println("Day " + DAY + " part 1 result: " + result);
  }

  private static char[][] createDish(List<String> input) {
    char[][] dish = new char[input.size()][input.get(0).length()];
    for (int row = 0; row < input.size(); row++) {
      for (int col = 0; col < input.get(row).length(); col++) {
        dish[row][col] = input.get(row).charAt(col);
      }
    }
    return dish;
  }


  public static void part2(String inputPath) {
    List<String> input = AocUtils.readInputDataToList(inputPath);
    char[][] dish = createDish(input);

    long result = 0;
    Map<Long, Long> memo = new LinkedHashMap<>();
    for (int i = 0; i < 1000000000; i++) {
      gravityNorth(dish);
      gravityWest(dish);
      gravitySouth(dish);
      gravityEast(dish);
      if (memo.put(getHash(dish), northWeight(dish)) != null) {
        List<Long> keys = new ArrayList<>(memo.keySet());
        int index = keys.indexOf(getHash(dish));
        keys = keys.subList(index, keys.size());

        int resIndex = (1000000000 - index) % (keys.size());
        resIndex = resIndex == 0 ? keys.size() - 1 : resIndex - 1;
        result = memo.get(keys.get(resIndex));
        break;
      }
    }

    System.out.println("Day " + DAY + " part 2 result: " + result);
  }

  private static long getHash(char[][] dish) {
    return Arrays.deepHashCode(dish) * northWeight(dish);
  }

  public static void gravityNorth(char[][] dish) {
    for (int row = 1; row < dish.length; row++) {
      for (int col = 0; col < dish[row].length; col++) {
        if (dish[row][col] == 'O') {
          for (int newRow = row; newRow > 0; newRow--) {
            if (dish[newRow][col] == 'O' && dish[newRow - 1][col] == '.') {
              dish[newRow][col] = '.';
              dish[newRow - 1][col] = 'O';
            }
          }
        }
      }
    }
  }

  public static void gravitySouth(char[][] dish) {
    for (int row = dish.length - 2; row >= 0; row--) {
      for (int col = 0; col < dish[row].length; col++) {
        if (dish[row][col] == 'O') {
          for (int newRow = row; newRow < dish.length - 1; newRow++) {
            if (dish[newRow][col] == 'O' && dish[newRow + 1][col] == '.') {
              dish[newRow][col] = '.';
              dish[newRow + 1][col] = 'O';
            }
          }
        }
      }
    }
  }

  public static void gravityWest(char[][] dish) {
    for (int col = 1; col < dish[0].length; col++) {
      for (int row = 0; row < dish.length; row++) {
        if (dish[row][col] == 'O') {
          for (int newCol = col; newCol > 0; newCol--) {
            if (dish[row][newCol] == 'O' && dish[row][newCol - 1] == '.') {
              dish[row][newCol] = '.';
              dish[row][newCol - 1] = 'O';
            }
          }
        }
      }
    }
  }

  public static void gravityEast(char[][] dish) {
    for (int col = dish[0].length - 2; col >= 0; col--) {
      for (int row = 0; row < dish.length; row++) {
        if (dish[row][col] == 'O') {
          for (int newCol = col; newCol <  dish[0].length - 1; newCol++) {
            if (dish[row][newCol] == 'O' && dish[row][newCol + 1] == '.') {
              dish[row][newCol] = '.';
              dish[row][newCol + 1] = 'O';
            }
          }
        }
      }
    }
  }

  public static long northWeight(char[][] dish) {
    long result = 0;
    for (int row = 0; row < dish.length; row++) {
      for (int col = 0; col < dish[row].length; col++) {
        if (dish[row][col] == 'O') {
          result += dish.length - row;
        }
      }
    }
    return result;
  }

  public static void main(String[] args) {
    part1("solutions/y23/d" + DAY + "/input.txt");
    part2("solutions/y23/d" + DAY + "/input.txt");
  }
}
