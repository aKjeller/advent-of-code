package y23.d18;

import utilities.java.AocUtils;

import java.util.*;

public class Solution {
  private static final String DAY = "18";

  public static void part1(String inputPath) {
    List<String> input = AocUtils.readInputDataToList(inputPath);

    Map<Point, String> map = createMapPart1(input);

    int minRow = Integer.MAX_VALUE;
    int minCol = Integer.MAX_VALUE;
    int maxRow = Integer.MIN_VALUE;
    int maxCol = Integer.MIN_VALUE;
    for (Point p : map.keySet()) {
      minRow = Math.min(p.getRowInt(), minRow);
      minCol = Math.min(p.getColInt(), minCol);
      maxRow = Math.max(p.getRowInt(), maxRow);
      maxCol = Math.max(p.getColInt(), maxCol);
    }
    boolean[][] diggedMap = new boolean[maxRow - minRow + 1][maxCol - minCol + 1];

    int rowDelta = -minRow;
    int colDelta = -minCol;
    for (int row = 0; row < diggedMap.length; row++) {
      for (int col = 0; col < diggedMap[row].length; col++) {
        diggedMap[row][col] = map.containsKey(new Point(row - rowDelta, col - colDelta));
      }
    }

    //printDiggedMap(diggedMap);

    for (int row = 0; row < diggedMap.length; row++) {
      for (int col = 0; col < diggedMap[row].length - 2; col++) {
        if (!diggedMap[row][col] && diggedMap[row][col + 1] && !diggedMap[row][col + 2]) {
          floodFill(diggedMap, row, col + 2);
          row = maxRow;
          col = maxCol;
        }
      }
    }

    //printDiggedMap(diggedMap);

    int result = 0;
    for (int row = 0; row < diggedMap.length; row++) {
      for (int col = 0; col < diggedMap[row].length; col++) {
        if(diggedMap[row][col]) {
          result += 1;
        }
      }
    }

    System.out.println("Day " + DAY + " part 1 result: " + result);
  }

  private static void printDiggedMap(boolean[][] diggedMap) {
    for (int row = 0; row < diggedMap.length; row++) {
      String stringRow = "";
      for (int col = 0; col < diggedMap[row].length; col++) {
        if(diggedMap[row][col]) {
          stringRow += "#";
        } else {
          stringRow += ".";
        }
      }
      System.out.println(stringRow);
    }
    System.out.println();
    System.out.println();
    System.out.println();
  }

  private static void floodFill(boolean[][] diggedMap, int row, int col) {
    Stack<Point> stack = new Stack<>();
    stack.add(new Point(row, col));

    while (!stack.isEmpty()) {
      Point p = stack.pop();
      if (!diggedMap[p.getRowInt()][p.getColInt()]) {
        diggedMap[p.getRowInt()][p.getColInt()] = true;
        stack.add(new Point(p.getRow() + 1, p.getCol()));
        stack.add(new Point(p.getRow() - 1, p.getCol()));
        stack.add(new Point(p.getRow() , p.getCol() + 1));
        stack.add(new Point(p.getRow() , p.getCol() - 1));
      }
    }
  }

  private static Map<Point, String> createMapPart1(List<String> input) {
    Map<Point, String> map = new HashMap<>();

    int currentRow = 0;
    int currentCol = 0;
    for (String inputLine : input) {
      String[] line = inputLine.split(" ");
      for (int i = 0; i < Integer.parseInt(line[1]); i++) {
        switch (line[0]) {
          case "R" -> map.put(new Point(currentRow, ++currentCol), line[2]);
          case "D" -> map.put(new Point(++currentRow, currentCol), line[2]);
          case "L" -> map.put(new Point(currentRow, --currentCol), line[2]);
          case "U" -> map.put(new Point(--currentRow, currentCol), line[2]);
        }
      }

    }
    return map;
  }

  public static void part2(String inputPath) {
    List<String> input = AocUtils.readInputDataToList(inputPath);
    List<Point> map = createMapPart2(input);

    long twoShoeLaceArea = 0;
    for (int i = 0; i < map.size() - 1; i++) {
      long a = map.get(i).getRow() * map.get(i + 1).getCol();
      long b = map.get(i + 1).getRow() * map.get(i).getCol();
      twoShoeLaceArea += a - b;
    }
    long area = Math.abs(twoShoeLaceArea / 2);

    long trenchArea = 0;
    for (int i = 0; i < map.size() - 1; i++) {
      trenchArea += Math.abs(map.get(i).getRow() - map.get(i + 1).getRow());
      trenchArea += Math.abs(map.get(i).getCol() - map.get(i + 1).getCol());
    }
    trenchArea = trenchArea / 2 + 1;

    long result = area + trenchArea;
    System.out.println("Day " + DAY + " part 2 result: " + result);
  }

  private static List<Point> createMapPart2(List<String> input) {
    List<Point> map = new ArrayList<>();

    int currentRow = 0;
    int currentCol = 0;
    map.add(new Point(currentRow, currentCol));
    for (String inputLine : input) {
      String[] line = inputLine.split(" ");
      int dist = Integer.parseInt(line[2].substring(2, 7), 16);
      String dir = line[2].substring(7, 8);
      switch (dir) {
        case "0" -> map.add(new Point(currentRow, currentCol += dist));
        case "1" -> map.add(new Point(currentRow += dist, currentCol));
        case "2" -> map.add(new Point(currentRow, currentCol -= dist));
        case "3" -> map.add(new Point(currentRow -= dist, currentCol));
      }

    }
    return map;
  }

  public static void main(String[] args) {
    part1("solutions/y23/d" + DAY + "/input.txt");
    part2("solutions/y23/d" + DAY + "/input.txt");
  }
}
