package y23.d16;

import utilities.java.AocUtils;

import java.util.*;

public class Solution {
  private static final String DAY = "16";

  private static Map<Beam, Boolean> hashMap = new HashMap<>();

  public static void part1(String inputPath) {
    List<String> input = AocUtils.readInputDataToList(inputPath);
    char[][] map = createMap(input);

    Beam start = new Beam(0, 0);
    start.setEntryDirection(Beam.Direction.EAST);
    hashMap.put(start, true);
    addBeam(map, start, Beam.Direction.EAST);

    long result = getResult(map);
    System.out.println("Day " + DAY + " part 1 result: " + result);
  }

  private static long getResult(char[][] map){
    char[][] resultMap = new char[map.length][map[0].length];
    for (int row = 0; row < map.length; row++) {
      for (int col = 0; col < map[0].length; col++) {
        resultMap[row][col] = '.';
      }
    }

    for (Beam beam : hashMap.keySet()) {
      resultMap[beam.getRow()][beam.getCol()] = '#';
    }

    long result = 0;
    for (int row = 0; row < map.length; row++) {
      for (int col = 0; col < map[0].length; col++) {
        if(resultMap[row][col] == '#') {
          result += 1;
        }
      }
    }
    return result;
  }

  private static void addBeam(char[][] map, Beam start, Beam.Direction lastDirection) {
    char object = map[start.getRow()][start.getCol()];

    switch (object) {
      case '.' -> {
        createAndAddBeam(map, start, lastDirection);
      }
      case '/' -> {
        Beam.Direction newDirection = switch (lastDirection) {
          case NORTH -> Beam.Direction.EAST;
          case EAST -> Beam.Direction.NORTH;
          case SOUTH -> Beam.Direction.WEST;
          case WEST -> Beam.Direction.SOUTH;
        };
        createAndAddBeam(map, start, newDirection);
      }
      case '\\' -> {
        Beam.Direction newDirection = switch (lastDirection) {
          case NORTH -> Beam.Direction.WEST;
          case EAST -> Beam.Direction.SOUTH;
          case SOUTH -> Beam.Direction.EAST;
          case WEST -> Beam.Direction.NORTH;
        };
        createAndAddBeam(map, start, newDirection);
      }
      case '-' -> {
        switch (lastDirection) {
          case NORTH, SOUTH -> {
            createAndAddBeam(map, start, Beam.Direction.WEST);
            createAndAddBeam(map, start, Beam.Direction.EAST);
          }
          case EAST, WEST -> createAndAddBeam(map, start, lastDirection);
        }
      }
      case '|' -> {
        switch (lastDirection) {
          case NORTH, SOUTH -> createAndAddBeam(map, start, lastDirection);
          case EAST, WEST -> {
            createAndAddBeam(map, start, Beam.Direction.NORTH);
            createAndAddBeam(map, start, Beam.Direction.SOUTH);
          }
        }
      }
    }
  }

  private static void createAndAddBeam(char[][] map, Beam start, Beam.Direction newDirection) {
    Beam next = new Beam(start.getRow(), start.getCol(), newDirection);
    next.setEntryDirection(newDirection);
    if (beamIsValid(map, next) && hashMap.put(next, true) == null) {
      start.addNext(next);
      addBeam(map, next, newDirection);
    }
  }

  private static boolean beamIsValid(char[][] map, Beam beam) {
    return beam.getRow() >= 0 && beam.getRow() < map.length && beam.getCol() >= 0 && beam.getCol() < map[0].length;
  }

  private static char[][] createMap(List<String> input) {
    char[][] map = new char[input.size()][input.get(0).length()];
    for (int row = 0; row < input.size(); row++) {
      for (int col = 0; col < input.get(row).length(); col++) {
        map[row][col] = input.get(row).charAt(col);
      }
    }
    return map;
  }

  public static void part2(String inputPath) {
    List<String> input = AocUtils.readInputDataToList(inputPath);
    char[][] map = createMap(input);

    long result = 0;
    for (int row = 0; row < map.length; row++) {
      hashMap = new HashMap<>();
      Beam start = new Beam(row, 0);
      start.setEntryDirection(Beam.Direction.EAST);
      hashMap.put(start, true);
      addBeam(map, start, Beam.Direction.EAST);
      result = Math.max(result, getResult(map));

      hashMap = new HashMap<>();
      start = new Beam(row, map[0].length - 1);
      start.setEntryDirection(Beam.Direction.WEST);
      hashMap.put(start, true);
      addBeam(map, start, Beam.Direction.WEST);
      result = Math.max(result, getResult(map));
    }

    for (int col = 0; col < map.length; col++) {
      hashMap = new HashMap<>();
      Beam start = new Beam(0, col);
      start.setEntryDirection(Beam.Direction.SOUTH);
      hashMap.put(start, true);
      addBeam(map, start, Beam.Direction.SOUTH);
      result = Math.max(result, getResult(map));

      hashMap = new HashMap<>();
      start = new Beam(map.length - 1, col);
      start.setEntryDirection(Beam.Direction.NORTH);
      hashMap.put(start, true);
      addBeam(map, start, Beam.Direction.NORTH);
      result = Math.max(result, getResult(map));
    }

    System.out.println("Day " + DAY + " part 2 result: " + result);
  }

  public static void main(String[] args) {
    part1("solutions/y23/d" + DAY + "/input.txt");
    part2("solutions/y23/d" + DAY + "/input.txt");
  }
}
