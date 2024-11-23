package y23.d17;

import utilities.java.AocUtils;

import java.util.*;

public class Solution {
  private static final String DAY = "17";

  public static void part1(String inputPath) {
    List<String> input = AocUtils.readInputDataToList(inputPath);
    Node[][] map = createMap(input);
    Map<Node, List<Node>> graph = createGraphPart1(map);

    Node startNode = new Node(0, 0, map[0][0].getCost(), Node.Direction.NORTH, 0);
    Node endNode = new Node(map.length - 1,  map[0].length - 1, 0);

    int result = dijkstra(graph, startNode, endNode, true);

    System.out.println("Day " + DAY + " part 1 result: " + result);
  }

  private static int dijkstra(Map<Node, List<Node>> graph, Node startNode, Node endNode, boolean part1) {
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

    Node minNode = null;
    int min = Integer.MAX_VALUE;
    for (Node node : finishedNodes) {
      if (node.getRow() == endNode.getRow() && node.getCol() == endNode.getCol()) {
        if (part1) {
          if (min > node.getDistance()) {
            min = node.getDistance();
            minNode = node;
          }
        } else {
          if (min > node.getDistance() && node.getnTimesInDirection() > 3) {
            min = node.getDistance();
            minNode = node;
          }
        }
      }
    }

    return min;
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


  private static Map<Node, List<Node>> createGraphPart1(Node[][] map) {
    Map<Node, List<Node>> graph = new HashMap<>();
    for (int row = 0; row < map.length; row++) {
      for (int col = 0; col < map[row].length; col++) {
        for (int i = 0; i < 5; i++) {
          int newRow = row - i;
          int newCol = col;
          if (isNodeValid(map, newRow, newCol)) {
            Node node = new Node(newRow, newCol, map[newRow][newCol].getCost(), Node.Direction.NORTH, i);
            List<Node> adjacency = getAdjacencyPart1(map, node, Node.Direction.NORTH, i + 1);
            graph.put(node, adjacency);
          }

          newRow = row + i;
          newCol = col;
          if (isNodeValid(map, newRow, newCol)) {
            Node node = new Node(newRow, newCol, map[newRow][newCol].getCost(), Node.Direction.SOUTH, i);
            List<Node> adjacency = getAdjacencyPart1(map, node, Node.Direction.SOUTH, i + 1);
            graph.put(node, adjacency);
          }

          newRow = row;
          newCol = col - 1;
          if (isNodeValid(map, newRow, newCol)) {
            Node node = new Node(newRow, newCol, map[newRow][newCol].getCost(), Node.Direction.WEST, i);
            List<Node> adjacency = getAdjacencyPart1(map, node, Node.Direction.WEST, i + 1);
            graph.put(node, adjacency);
          }

          newRow = row;
          newCol = col + 1;
          if (isNodeValid(map, newRow, newCol)) {
            Node node = new Node(newRow, newCol, map[newRow][newCol].getCost(), Node.Direction.EAST, i);
            List<Node> adjacency = getAdjacencyPart1(map, node, Node.Direction.EAST, i + 1);
            graph.put(node, adjacency);
          }
        }
      }
    }
    return graph;
  }

  private static List<Node> getAdjacencyPart1(Node[][] map, Node node, Node.Direction direction, int times) {
    List<Node> adjacency = new ArrayList<>();
    if (!direction.equals(Node.Direction.SOUTH)) {
      if (direction.equals(Node.Direction.NORTH)) {
        int newRow = node.getRow() - 1;
        int newCol = node.getCol();
        if (isNodeValid(map, newRow, newCol) && times < 4) {
          Node newNode = new Node(newRow, newCol, map[newRow][newCol].getCost(), Node.Direction.NORTH, times);
          addNodeFromMap(adjacency, map, newNode);
        }
      } else {
        int newRow = node.getRow() - 1;
        int newCol = node.getCol();
        if (isNodeValid(map, newRow, newCol)) {
          Node newNode = new Node(newRow, newCol, map[newRow][newCol].getCost(), Node.Direction.NORTH, 1);
          addNodeFromMap(adjacency, map, newNode);
        }
      }
    }

    if (!direction.equals(Node.Direction.NORTH)) {
      if (direction.equals(Node.Direction.SOUTH)) {
        int newRow = node.getRow() + 1;
        int newCol = node.getCol();
        if (isNodeValid(map, newRow, newCol) && times < 4) {
          Node newNode = new Node(newRow, newCol, map[newRow][newCol].getCost(), Node.Direction.SOUTH, times);
          addNodeFromMap(adjacency, map, newNode);
        }
      } else {
        int newRow = node.getRow() + 1;
        int newCol = node.getCol();
        if (isNodeValid(map, newRow, newCol)) {
          Node newNode = new Node(newRow, newCol, map[newRow][newCol].getCost(), Node.Direction.SOUTH, 1);
          addNodeFromMap(adjacency, map, newNode);
        }
      }
    }

    if (!direction.equals(Node.Direction.EAST)) {
      if (direction.equals(Node.Direction.WEST)) {
        int newRow = node.getRow();
        int newCol = node.getCol() - 1;
        if (isNodeValid(map, newRow, newCol) && times < 4) {
          Node newNode = new Node(newRow, newCol, map[newRow][newCol].getCost(), Node.Direction.WEST, times);
          addNodeFromMap(adjacency, map, newNode);
        }
      } else {
        int newRow = node.getRow();
        int newCol = node.getCol() - 1;
        if (isNodeValid(map, newRow, newCol)) {
          Node newNode = new Node(newRow, newCol, map[newRow][newCol].getCost(), Node.Direction.WEST, 1);
          addNodeFromMap(adjacency, map, newNode);
        }
      }
    }

    if (!direction.equals(Node.Direction.WEST)) {
      if (direction.equals(Node.Direction.EAST)) {
        int newRow = node.getRow();
        int newCol = node.getCol() + 1;
        if (isNodeValid(map, newRow, newCol) && times < 4) {
          Node newNode = new Node(newRow, newCol, map[newRow][newCol].getCost(), Node.Direction.EAST, times);
          addNodeFromMap(adjacency, map, newNode);
        }
      } else {
        int newRow = node.getRow();
        int newCol = node.getCol() + 1;
        if (isNodeValid(map, newRow, newCol)) {
          Node newNode = new Node(newRow, newCol, map[newRow][newCol].getCost(), Node.Direction.EAST, 1);
          addNodeFromMap(adjacency, map, newNode);
        }
      }
    }

    return adjacency;
  }

  private static void addNodeFromMap(List<Node> list, Node[][] map, Node node) {
    if ((node.getRow() >= 0) && (node.getRow() < map.length) && (node.getCol() >= 0) && (node.getCol() < map[node.getRow()].length)) {
      list.add(node);
    }
  }

  private static boolean isNodeValid(Node[][] map, int row, int col) {
    return ((row >= 0) && (row < map.length) && (col >= 0) && (col < map[row].length));
  }

  private static Node[][] createMap(List<String> input) {
    Node[][] map = new Node[input.size()][input.get(0).length()];
    for (int row = 0; row < input.size(); row++) {
      for (int col = 0; col < input.get(row).length(); col++) {
        map[row][col] = new Node(row, col, Character.getNumericValue(input.get(row).charAt(col)), null, 0);
      }
    }
    return map;
  }

  public static void part2(String inputPath) {
    List<String> input = AocUtils.readInputDataToList(inputPath);
    Node[][] map = createMap(input);
    Map<Node, List<Node>> graph = createGraphPart2(map);

    Node startNode = new Node(0, 0, map[0][0].getCost(), Node.Direction.NORTH, 0);
    List<Node> startAdjacency = new ArrayList<>();
    startAdjacency.add(new Node(0, 1, map[0][1].getCost(), Node.Direction.EAST, 1));
    startAdjacency.add(new Node(1, 0, map[1][0].getCost(), Node.Direction.SOUTH, 1));
    graph.put(startNode, startAdjacency);
    Node endNode = new Node(map.length - 1,  map[0].length - 1, 0);

    int result = dijkstra(graph, startNode, endNode, false);
    System.out.println("Day " + DAY + " part 2 result: " + result);
  }

  private static Map<Node, List<Node>> createGraphPart2(Node[][] map) {
    Map<Node, List<Node>> graph = new HashMap<>();
    for (int row = 0; row < map.length; row++) {
      for (int col = 0; col < map[row].length; col++) {
        for (int i = 0; i < 12; i++) {
          int newRow = row - i;
          int newCol = col;
          if (isNodeValid(map, newRow, newCol)) {
            Node node = new Node(newRow, newCol, map[newRow][newCol].getCost(), Node.Direction.NORTH, i);
            List<Node> adjacency = getAdjacencyPart2(map, node, Node.Direction.NORTH, i + 1);
            graph.put(node, adjacency);
          }

          newRow = row + i;
          newCol = col;
          if (isNodeValid(map, newRow, newCol)) {
            Node node = new Node(newRow, newCol, map[newRow][newCol].getCost(), Node.Direction.SOUTH, i);
            List<Node> adjacency = getAdjacencyPart2(map, node, Node.Direction.SOUTH, i + 1);
            graph.put(node, adjacency);
          }

          newRow = row;
          newCol = col - 1;
          if (isNodeValid(map, newRow, newCol)) {
            Node node = new Node(newRow, newCol, map[newRow][newCol].getCost(), Node.Direction.WEST, i);
            List<Node> adjacency = getAdjacencyPart2(map, node, Node.Direction.WEST, i + 1);
            graph.put(node, adjacency);
          }

          newRow = row;
          newCol = col + 1;
          if (isNodeValid(map, newRow, newCol)) {
            Node node = new Node(newRow, newCol, map[newRow][newCol].getCost(), Node.Direction.EAST, i);
            List<Node> adjacency = getAdjacencyPart2(map, node, Node.Direction.EAST, i + 1);
            graph.put(node, adjacency);
          }
        }
      }
    }
    return graph;
  }

  private static List<Node> getAdjacencyPart2(Node[][] map, Node node, Node.Direction direction, int times) {
    List<Node> adjacency = new ArrayList<>();
    if (!direction.equals(Node.Direction.SOUTH)) {
      if (direction.equals(Node.Direction.NORTH)) {
        int newRow = node.getRow() - 1;
        int newCol = node.getCol();
        if (isNodeValid(map, newRow, newCol) && times < 11) {
          Node newNode = new Node(newRow, newCol, map[newRow][newCol].getCost(), Node.Direction.NORTH, times);
          addNodeFromMap(adjacency, map, newNode);
        }
      } else if (times > 4) {
        int newRow = node.getRow() - 1;
        int newCol = node.getCol();
        if (isNodeValid(map, newRow, newCol)) {
          Node newNode = new Node(newRow, newCol, map[newRow][newCol].getCost(), Node.Direction.NORTH, 1);
          addNodeFromMap(adjacency, map, newNode);
        }
      }
    }

    if (!direction.equals(Node.Direction.NORTH)) {
      if (direction.equals(Node.Direction.SOUTH)) {
        int newRow = node.getRow() + 1;
        int newCol = node.getCol();
        if (isNodeValid(map, newRow, newCol) && times < 11) {
          Node newNode = new Node(newRow, newCol, map[newRow][newCol].getCost(), Node.Direction.SOUTH, times);
          addNodeFromMap(adjacency, map, newNode);
        }
      } else if (times > 4) {
        int newRow = node.getRow() + 1;
        int newCol = node.getCol();
        if (isNodeValid(map, newRow, newCol)) {
          Node newNode = new Node(newRow, newCol, map[newRow][newCol].getCost(), Node.Direction.SOUTH, 1);
          addNodeFromMap(adjacency, map, newNode);
        }
      }
    }

    if (!direction.equals(Node.Direction.EAST)) {
      if (direction.equals(Node.Direction.WEST)) {
        int newRow = node.getRow();
        int newCol = node.getCol() - 1;
        if (isNodeValid(map, newRow, newCol) && times < 11) {
          Node newNode = new Node(newRow, newCol, map[newRow][newCol].getCost(), Node.Direction.WEST, times);
          addNodeFromMap(adjacency, map, newNode);
        }
      } else if (times > 4) {
        int newRow = node.getRow();
        int newCol = node.getCol() - 1;
        if (isNodeValid(map, newRow, newCol)) {
          Node newNode = new Node(newRow, newCol, map[newRow][newCol].getCost(), Node.Direction.WEST, 1);
          addNodeFromMap(adjacency, map, newNode);
        }
      }
    }

    if (!direction.equals(Node.Direction.WEST)) {
      if (direction.equals(Node.Direction.EAST)) {
        int newRow = node.getRow();
        int newCol = node.getCol() + 1;
        if (isNodeValid(map, newRow, newCol) && times < 11) {
          Node newNode = new Node(newRow, newCol, map[newRow][newCol].getCost(), Node.Direction.EAST, times);
          addNodeFromMap(adjacency, map, newNode);
        }
      } else if (times > 4) {
        int newRow = node.getRow();
        int newCol = node.getCol() + 1;
        if (isNodeValid(map, newRow, newCol)) {
          Node newNode = new Node(newRow, newCol, map[newRow][newCol].getCost(), Node.Direction.EAST, 1);
          addNodeFromMap(adjacency, map, newNode);
        }
      }
    }

    return adjacency;
  }

  public static void main(String[] args) {
    part1("solutions/y23/d" + DAY + "/input.txt");
    part2("solutions/y23/d" + DAY + "/input.txt");
  }
}
