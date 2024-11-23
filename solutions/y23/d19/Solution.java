package y23.d19;

import com.google.gson.Gson;
import utilities.java.AocUtils;

import java.util.*;
import java.util.List;

public class Solution {
  private static final String DAY = "19";
  private static final Gson gson = new Gson();

  public static void part1(String inputPath) {
    List<String> input = AocUtils.readInputDataToList(inputPath);

    Map<String, WorkFlow> workFlows = new HashMap<>();
    List<Part> parts = new ArrayList<>();

    boolean firstHalf = true;
    for (String line : input) {
      if (line.isBlank()) {
        firstHalf = false;
      } else if (firstHalf) {
        String label = line.substring(0, line.indexOf('{'));
        String rules = line.substring(line.indexOf('{') + 1, line.length() - 1);

        workFlows.put(label, new WorkFlow(rules));

      } else {
        parts.add(gson.fromJson(line, Part.class));
      }
    }

    long result = 0;
    for (Part part : parts) {
      if (evaulatePart(workFlows, "in", part)) {
        result += part.getX();
        result += part.getM();
        result += part.getA();
        result += part.getS();
      }
    }

    System.out.println("Day " + DAY + " part 1 result: " + result);
  }

  private static boolean evaulatePart(Map<String, WorkFlow> workFlows, String in, Part part) {
    String label = workFlows.get(in).evaluateRules(part);
    return switch (label) {
      case "A" -> true;
      case "R" -> false;
      default -> evaulatePart(workFlows, label, part);
    };
  }


  public static void part2(String inputPath) {
    List<String> input = AocUtils.readInputDataToList(inputPath);

    Map<String, WorkFlow> workFlows = new HashMap<>();
    for (String line : input) {
      if (line.isBlank()) {
        break;
      }
      String label = line.substring(0, line.indexOf('{'));
      String rules = line.substring(line.indexOf('{') + 1, line.length() - 1);
      workFlows.put(label, new WorkFlow(rules));
    }

    Part lower = new Part(1, 1, 1, 1);
    Part upper = new Part(4000, 4000, 4000, 4000);
    Node top = new Node("in", lower, upper);
    generateTree(workFlows, top);

    long result = traverseTree(top);
    System.out.println("Day " + DAY + " part 2 result: " + result);

  }


  private static long traverseTree(Node node) {
    if (node.getLabel().equals("A")) {
      long result = node.getUpperBound().getX() - node.getLowerBound().getX() + 1;
      result *= node.getUpperBound().getM() - node.getLowerBound().getM() + 1;
      result *= node.getUpperBound().getA() - node.getLowerBound().getA() + 1;
      result *= node.getUpperBound().getS() - node.getLowerBound().getS() + 1;
      return result;
    } else if (node.getLabel().equals("R")) {
      return 0;
    } else {
      long result = 0;
      for (Node children : node.getChildren()) {
        result += traverseTree(children);
      }
      return result;
    }
  }

  private static void generateTree(Map<String, WorkFlow> workFlows, Node node) {
    if (node.getLabel().equals("A") || node.getLabel().equals("R")) {
      return;
    }

    Part lower = node.getLowerBound().copy();
    Part upper = node.getUpperBound().copy();

    WorkFlow workFlow = workFlows.get(node.getLabel());

    for (String rule : workFlow.getRules()) {
      String condition = rule.split(":")[0];
      String dest = rule.split(":")[1];
      if (condition.contains("<")) {
        int cutOff = Integer.parseInt(condition.substring(condition.indexOf("<") + 1));
        switch (condition.charAt(0)) {
          case 'x' -> {
            int tmp = upper.getX();
            upper.setX(cutOff - 1);
            Node children = new Node(dest, lower.copy(), upper.copy());
            node.addChildren(children);
            generateTree(workFlows, children);
            upper.setX(tmp);
            lower.setX(cutOff);
          }
          case 'm' -> {
            int tmp = upper.getM();
            upper.setM(cutOff - 1);
            Node children = new Node(dest, lower.copy(), upper.copy());
            node.addChildren(children);
            generateTree(workFlows, children);
            upper.setM(tmp);
            lower.setM(cutOff);
          }
          case 'a' -> {
            int tmp = upper.getA();
            upper.setA(cutOff - 1);
            Node children = new Node(dest, lower.copy(), upper.copy());
            node.addChildren(children);
            generateTree(workFlows, children);
            upper.setA(tmp);
            lower.setA(cutOff);
          }
          case 's' -> {
            int tmp = upper.getS();
            upper.setS(cutOff - 1);
            Node children = new Node(dest, lower.copy(), upper.copy());
            node.addChildren(children);
            generateTree(workFlows, children);
            upper.setS(tmp);
            lower.setS(cutOff);
          }
        }
      } else {
        int cutOff = Integer.parseInt(condition.substring(condition.indexOf(">") + 1));
        switch (condition.charAt(0)) {
          case 'x' -> {
            int tmp = lower.getX();
            lower.setX(cutOff + 1);
            Node children = new Node(dest, lower.copy(), upper.copy());
            node.addChildren(children);
            generateTree(workFlows, children);
            lower.setX(tmp);
            upper.setX(cutOff);
          }
          case 'm' -> {
            int tmp = lower.getM();
            lower.setM(cutOff + 1);
            Node children = new Node(dest, lower.copy(), upper.copy());
            node.addChildren(children);
            generateTree(workFlows, children);
            lower.setM(tmp);
            upper.setM(cutOff);
          }
          case 'a' -> {
            int tmp = lower.getA();
            lower.setA(cutOff + 1);
            Node children = new Node(dest, lower.copy(), upper.copy());
            node.addChildren(children);
            generateTree(workFlows, children);
            lower.setA(tmp);
            upper.setA(cutOff);
          }
          case 's' -> {
            int tmp = lower.getS();
            lower.setS(cutOff + 1);
            Node children = new Node(dest, lower.copy(), upper.copy());
            node.addChildren(children);
            generateTree(workFlows, children);
            lower.setS(tmp);
            upper.setS(cutOff);
          }
        }
      }
    }
    Node children = new Node(workFlow.getEnd(), lower.copy(), upper.copy());
    node.addChildren(children);
    generateTree(workFlows, children);
  }

  public static void main(String[] args) {
    part1("solutions/y23/d" + DAY + "/input.txt");
    part2("solutions/y23/d" + DAY + "/input.txt");
  }
}
