package y23.d03;

import utilities.java.AocUtils;

import java.util.Arrays;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

public class Solution {

  public static void part1(String inputPath) {
    List<String> input = AocUtils.readInputDataToList(inputPath);

    boolean[][] symbol = new boolean[input.size()][input.get(0).length()];
    Arrays.stream(symbol).forEach(a -> Arrays.fill(a, false));

    for (int row = 0; row < input.size(); row++) {
      for (int col = 0; col < input.get(row).length(); col++) {
        char current = input.get(row).charAt(col);
        switch (current) {
          case '.', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9': break;
          default: {
            for (int symbolRow = row - 1; symbolRow <= row + 1; symbolRow++) {
              for (int symbolCol = col - 1; symbolCol <= col + 1; symbolCol++) {
                try {
                  symbol[symbolRow][symbolCol] = true;
                } catch (Exception ignore) {}
              }
            }
          }
        }
      }
    }

    Set<Part> parts = new HashSet<>();
    for (int row = 0; row < input.size(); row++) {
      for (int col = 0; col < input.get(row).length(); col++) {
        char current = input.get(row).charAt(col);
        if (Character.isDigit(current)) {
          String partNumber = "";
          boolean isNearSymbol = false;
          while (Character.isDigit(current)) {
            if (symbol[row][col]) {
              isNearSymbol = true;
            }
            partNumber += String.valueOf(current);
            if (col < input.get(row).length() - 1) {
              current = input.get(row).charAt(++col);
            } else {
              current = 'E';
            }
          }
          if (isNearSymbol) {
            parts.add(new Part(Integer.parseInt(partNumber)));
          }
        }
      }
    }

    int result = parts.stream()
            .map(Part::getPartNumber)
            .reduce(0, Integer::sum);

    System.out.println("Part 1 result: " + result);
  }
  public static void part2(String inputPath) {
    List<String> input = AocUtils.readInputDataToList(inputPath);

    boolean[][] gear = new boolean[input.size()][input.get(0).length()];
    Arrays.stream(gear).forEach(a -> Arrays.fill(a, false));

    for (int row = 0; row < input.size(); row++) {
      for (int col = 0; col < input.get(row).length(); col++) {
        if(input.get(row).charAt(col) == '*') {
          gear[row][col] = true;
        }
      }
    }

    int result = 0;
    for (int row = 0; row < input.size(); row++) {
      for (int col = 0; col < input.get(row).length(); col++) {
        if(gear[row][col]) {
          result += calculateGearRatio(input, row, col);
        }
      }
    }
    System.out.println("Part 2 result: " + result);
  }

  private static int calculateGearRatio(List<String> input, int row, int col) {
    int partA = 0;
    int partB = 0;
    for (int rowIndex = row - 1; rowIndex < row + 2; rowIndex++) {
      for (int colIndex = col - 1; colIndex < col + 2; colIndex++) {
        if (Character.isDigit(input.get(rowIndex).charAt(colIndex))) {
          int part = getPartNumberFromRowAndIndex(input.get(rowIndex), colIndex);
          if (partA != 0) {
            partB = part;
          } else {
            partA = part;
            if (Character.isDigit(input.get(rowIndex).charAt(col))) {
              colIndex += 10;
            }
          }

          colIndex++;
        }
      }
    }
    return partA * partB;
  }

  private static int getPartNumberFromRowAndIndex(String s, int colIndex) {
    int startIndex = colIndex;
    int endIndex = colIndex;

    char current = s.charAt(colIndex);
    while (Character.isDigit(current)) {
      startIndex--;
      try {
        current = s.charAt(startIndex);
      } catch (Exception e) {
        current = 'E';
      }
    }

    current = s.charAt(colIndex);
    while (Character.isDigit(current)) {
      endIndex++;
      try {
        current = s.charAt(endIndex);
      } catch (Exception e) {
        current = 'E';
      }
    }

    return Integer.parseInt(s.substring(startIndex + 1, endIndex));
  }


  public static void main(String[] args) {
    part1("solutions/y23/d03/input.txt");
    part2("solutions/y23/d03/input.txt");
  }
}
