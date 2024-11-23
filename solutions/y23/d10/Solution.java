package y23.d10;

import utilities.java.AocUtils;

import java.util.List;
import java.util.Objects;

public class Solution {
  private static final String DAY = "10";
  private static Pipe[][] map = null;

  public static void part1(String inputPath) {
    List<String> input = AocUtils.readInputDataToList(inputPath);

    map = new Pipe[input.size()][input.get(0).length()];
    int currentRow = 0;
    int currentCol = 0;
    for (int row = 0; row < input.size(); row++) {
      for (int col = 0; col < map[row].length; col++) {
        map[row][col] = getPipeFromChar(input.get(row).charAt(col));
        if (map[row][col].isStart()) {
          currentRow = row;
          currentCol = col;
        }
      }
    }

    int counter = 0;
    do {
      for (Pipe.Direction direction : map[currentRow][currentCol].getConnections()) {
        if (isConnectionInDirection(currentRow, currentCol, direction) || isStartInDirection(currentRow, currentCol, direction)) {
          switch (direction) {
            case NORTH -> currentRow -= 1;
            case EAST -> currentCol += 1;
            case SOUTH -> currentRow += 1;
            case WEST -> currentCol -= 1;
          }
          counter++;
          map[currentRow][currentCol].connectFrom(direction);
          break;
        }
      }
    } while (!map[currentRow][currentCol].isStart());

    long result = counter / 2;
    System.out.println("Day " + DAY + " part 1 result: " + result);
  }

  private static boolean isConnectionInDirection(int row, int col, Pipe.Direction direction) {
    if (direction.equals(map[row][col].getConnectedWith())) {
      return false;
    }

    try {
      return switch (direction) {
        case NORTH -> map[row - 1][col].contains(Pipe.Direction.SOUTH);
        case EAST -> map[row][col + 1].contains(Pipe.Direction.WEST);
        case SOUTH -> map[row + 1][col].contains(Pipe.Direction.NORTH);
        case WEST -> map[row][col - 1].contains(Pipe.Direction.EAST);
      };
    } catch (Exception e) {
      return false;
    }
  }

  private static boolean isStartInDirection(int row, int col, Pipe.Direction direction) {
    try {
      return switch (direction) {
        case NORTH -> map[row - 1][col].isStart();
        case EAST -> map[row][col + 1].isStart();
        case SOUTH -> map[row + 1][col].isStart();
        case WEST -> map[row][col - 1].isStart();
      };
    } catch (Exception e) {
      return false;
    }
  }

  private static Pipe getPipeFromChar(char s) {
    return switch (s) {
      case '|' -> new Pipe(Pipe.Direction.NORTH, Pipe.Direction.SOUTH);
      case '-' -> new Pipe(Pipe.Direction.WEST, Pipe.Direction.EAST);
      case 'L' -> new Pipe(Pipe.Direction.NORTH, Pipe.Direction.EAST);
      case 'J' -> new Pipe(Pipe.Direction.NORTH, Pipe.Direction.WEST);
      case '7' -> new Pipe(Pipe.Direction.WEST, Pipe.Direction.SOUTH);
      case 'F' -> new Pipe(Pipe.Direction.SOUTH, Pipe.Direction.EAST);
      case 'S' -> new Pipe(true);
      case '.' -> new Pipe(false);
      default -> null;
    };
  }

  // REQUIRES MAP FROM PART 1!!!
  public static void part2(String inputPath) {
    List<String> input = AocUtils.readInputDataToList(inputPath);
    for (int row = 0; row < input.size(); row++) {
      for (int col = 0; col < map[row].length; col++) {
        if (map[row][col].isStart()) {
          Pipe.Direction correctDirection = null;
          for (Pipe.Direction direction : map[row][col].getConnections()) {
            try {
              switch (direction) {
                case NORTH -> correctDirection = map[row - 1][col].getConnectedWith() == Pipe.Direction.SOUTH ? Pipe.Direction.NORTH : correctDirection;
                case EAST -> correctDirection = map[row][col + 1].getConnectedWith() == Pipe.Direction.WEST ? Pipe.Direction.EAST : correctDirection;
                case SOUTH -> correctDirection = map[row + 1][col].getConnectedWith() == Pipe.Direction.NORTH ? Pipe.Direction.SOUTH : correctDirection;
                case WEST -> correctDirection = map[row][col - 1].getConnectedWith() == Pipe.Direction.EAST ? Pipe.Direction.WEST : correctDirection;
              }
            } catch (Exception e) {
            }
          }
          assert correctDirection != null;
          map[row][col].setConnections(List.of(correctDirection));
        }
      }
    }


    long result = 0;
    for (int row = 0; row < input.size(); row++) {
      for (int col = 0; col < map[row].length; col++) {
        if (Objects.isNull(map[row][col].getConnectedWith())) { //  && map[row][col].getConnections().size() == 0
          int counterIn = 0;

          boolean passedNorth = false;
          boolean passedSouth = false;

          for (int i = 0; i < col; i++) {

            if (map[row][i].getConnectedWith() == null) {
              passedNorth = false;
              passedSouth = false;
            } else {
              Pipe.Direction connectedWith = map[row][i].getConnectedWith();
              if (connectedWith == Pipe.Direction.NORTH) {
                passedNorth = passedNorth ? false : true;
              } else if (connectedWith == Pipe.Direction.SOUTH) {
                passedSouth = passedSouth ? false : true;
              }
              if (map[row][i].getConnections().contains(Pipe.Direction.NORTH)) {
                passedNorth = passedNorth ? false : true;
              } else if (map[row][i].getConnections().contains(Pipe.Direction.SOUTH)) {
                passedSouth = passedSouth ? false : true;
              }
            }

            if (map[row][i].getConnectedWith() != null && passedSouth && passedNorth) {
              counterIn++;
              passedNorth = false;
              passedSouth = false;
            }

          }
          if (counterIn % 2 != 0) {
            result += 1;
          }
        }

      }
    }

    System.out.println("Day " + DAY + " part 2 result: " + result);
  }


  public static void main(String[] args) {
    part1("solutions/y23/d" + DAY + "/input.txt");
    part2("solutions/y23/d" + DAY + "/input.txt");
  }
}
