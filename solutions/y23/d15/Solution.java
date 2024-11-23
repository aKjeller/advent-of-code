package y23.d15;

import utilities.java.AocUtils;

import java.util.*;
import java.util.stream.Stream;

public class Solution {
  private static final String DAY = "15";

  public static void part1(String inputPath) {
    List<String> input = AocUtils.readInputDataToList(inputPath);
    String[] lines = input.get(0).split(",");

    int result = 0;
    for (String line :lines) {
      result += getHash(line);
    }

    System.out.println("Day " + DAY + " part 1 result: " + result);
  }

  private static int getHash(String input) {
    int currentValue = 0;
    for (char c : input.toCharArray()){
      currentValue = currentValue + (int) c;
      currentValue = currentValue * 17;
      currentValue = currentValue % 256;
    }
    return currentValue;
  }

  public static void part2(String inputPath) {
    List<String> input = AocUtils.readInputDataToList(inputPath);
    String[] lines = input.get(0).split(",");

    List<Map<String, Integer>> boxes = new ArrayList<>();
    Stream.generate(() -> new LinkedHashMap<String, Integer>()).limit(256).forEach(boxes::add);

    for (String line :lines) {
      if (line.contains("-")) {
        String label = line.substring(0, line.length() - 1);
        int boxNumber = getHash(label);
        boxes.get(boxNumber).remove(label);
      } else {
        String label = line.substring(0, line.length() - 2);
        int boxNumber = getHash(label);
        int focalLength = Integer.parseInt(String.valueOf(line.charAt(line.length() - 1)));
        boxes.get(boxNumber).put(label, focalLength);
      }
    }

    long result = 0;
    for(int i = 0; i < boxes.size(); i++) {
      Map<String, Integer> box = boxes.get(i);
      List<String> keys = new ArrayList<>(box.keySet());
      for (int lens = 0; lens < keys.size(); lens++) {
        result += (long) (1 + i) * (lens + 1) * box.get(keys.get(lens));
      }
    }

    System.out.println("Day " + DAY + " part 2 result: " + result);
  }

  public static void main(String[] args) {
    part1("solutions/y23/d" + DAY + "/input.txt");
    part2("solutions/y23/d" + DAY + "/input.txt");
  }
}
