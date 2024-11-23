package y23.d21;

import utilities.java.AocUtils;

import java.util.*;

public class Solution {
  private static final String DAY = "21";

  public static void part1(String inputPath) {
    List<String> input = AocUtils.readInputDataToList(inputPath);
    Node[][] map = createMap(input);

    Node startNode = null;
    for (int row = 0; row < input.size(); row++) {
      for (int col = 0; col < input.get(row).length(); col++) {
        if (input.get(row).charAt(col) == 'S'){
          startNode = new Node(row, col, 1);
          map[row][col] = startNode;
        }
      }
    }

    Map<Node, List<Node>> graph = createGraphPart(map);

    dijkstra(graph, startNode);

    long steps = 65;
    long result = 0;
    for (Node node : graph.keySet()) {
      if (node.getDistance() % 2 != 0 && node.getDistance() <= steps) {
        result += 1;
      }
    }

    System.out.println("Day " + DAY + " part 1 result: " + result);
  }

  private static void dijkstra(Map<Node, List<Node>> graph, Node startNode) {
    Set<Node> finishedNodes = new HashSet<>();
    Set<Node> unFinishedNodes = new HashSet<>();

    startNode.setDistance(0);
    unFinishedNodes.add(startNode);

    while (unFinishedNodes.size() != 0) {
      Node currentNode = getLowestDistance(unFinishedNodes);
      unFinishedNodes.remove(currentNode);
      for (Node node : graph.get(currentNode)) {
        if (!finishedNodes.contains(node)) {
          calculateMinimumDistance(node, node.getCost(), currentNode);
          unFinishedNodes.add(node);
        }
      }
      finishedNodes.add(currentNode);
    }
  }

  private static void calculateMinimumDistance(Node evaluationNode, int cost, Node sourceNode) {
    int sourceDistance = sourceNode.getDistance();
    if (sourceDistance + cost < evaluationNode.getDistance()) {
      evaluationNode.setDistance(sourceDistance + cost);
      LinkedList<Node> shortestPath = new LinkedList<>(sourceNode.getShortestPath());
      shortestPath.add(sourceNode);
      evaluationNode.setShortestPath(shortestPath);
    }
  }
  private static Node getLowestDistance(Set<Node> nodes) {
    Node lowestDistanceNode = null;
    int lowestDistance = Integer.MAX_VALUE;
    for (Node node : nodes) {
      if (node.getDistance() < lowestDistance) {
        lowestDistanceNode = node;
        lowestDistance = node.getDistance();
      }
    }
    return lowestDistanceNode;
  }


  private static Map<Node, List<Node>> createGraphPart(Node[][] map) {
    Map<Node, List<Node>> graph = new HashMap<>();
    for (int row = 0; row < map.length; row++) {
      for (int col = 0; col < map[row].length; col++) {
        if (map[row][col] != null) {
          graph.put(map[row][col], getAdjacencyNodes(map, map[row][col]));
        }
      }
    }
    return graph;
  }

  private static List<Node> getAdjacencyNodes(Node[][] map, Node node) {
    List<Node> adjacency = new ArrayList<>();
    addNodeFromMap(adjacency, map, node.getRow() + 1, node.getCol());
    addNodeFromMap(adjacency, map, node.getRow() - 1, node.getCol());
    addNodeFromMap(adjacency, map, node.getRow(), node.getCol() + 1);
    addNodeFromMap(adjacency, map, node.getRow(), node.getCol() - 1);
    return adjacency;
  }

  private static void addNodeFromMap(List<Node> list, Node[][] map, int row, int col) {
    if ((row >= 0) && (row < map.length) && (col >= 0) && (col < map[row].length)) {
      if (map[row][col] != null){
        list.add(map[row][col]);
      }
    }
  }

  private static Node[][] createMap(List<String> input) {
    Node[][] map = new Node[input.size()][input.get(0).length()];
    for (int row = 0; row < input.size(); row++) {
      for (int col = 0; col < input.get(row).length(); col++) {
        if (input.get(row).charAt(col) == '.'){
          map[row][col] = new Node(row, col, 1);
        } else {
          map[row][col] = null;
        }
      }
    }
    return map;
  }

  public static void part2(String inputPath) {
    List<String> oldInput = AocUtils.readInputDataToList(inputPath);

    List<String> input = new ArrayList<>();
    for (int i = 0; i < 7; i++) {
      for (String line : oldInput){
        input.add(line.repeat(7));
      }
    }

    Node[][] map = createMap(input);

    List<Node> startNodes = new ArrayList<>();
    for (int row = 0; row < input.size(); row++) {
      for (int col = 0; col < input.get(row).length(); col++) {
        if (input.get(row).charAt(col) == 'S'){
          Node startNode = new Node(row, col, 1);
          map[row][col] = startNode;
          startNodes.add(startNode);
        }
      }
    }

    Map<Node, List<Node>> graph = createGraphPart(map);
    dijkstra(graph, startNodes.get(startNodes.size() / 2));

    for (long steps = 0; steps <= 4; steps++) {
      long result = 0;
      long n = 131 * steps + 65;
      for (Node node : graph.keySet()) { // node.getDistance() % 2 == 0 &&
        if (n % 2 == 0) {
          if (node.getDistance() % 2 == 0 && node.getDistance() <= n) {
            result += 1;
          }
        } else {
          if (node.getDistance() % 2 == 1 && node.getDistance() <= n) {
            result += 1;
          }
        }
      }

      System.out.println("Day " + DAY + " part 2 result: " + "(" + steps + "," + result + ")");
      // wolfram: quadratic equation(0,3911),(1,34786),(2,96435) -> 3911 + 15488 x + 15387 x^2
      // (26501365 - 65) / 131 = 202300
      // 3911 + 15488 x + 15387 x^2 where x = 202300
    }

  }

  public static void main(String[] args) {
    part1("solutions/y23/d" + DAY + "/input.txt");
    part2("solutions/y23/d" + DAY + "/input.txt");
  }
}
