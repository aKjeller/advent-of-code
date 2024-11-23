package y23.d23;

import utilities.java.AocUtils;

import java.util.*;

public class Solution {
  private static final String DAY = "23";

  private static int rowMax = 0;
  private static int colMax = 0;


  public static void part1(String inputPath) {
    List<String> input = AocUtils.readInputDataToList(inputPath);
    Node[][] map = createMap(input);

    Map<Node, List<Node>> graph = createGraph(map);

    Node startNode = map[0][1];

    long result = longestPath(graph, startNode, startNode);
    System.out.println("Day " + DAY + " part 1 result: " + result);
  }

  private static long longestPath(Map<Node, List<Node>> graph, Node earlierNode, Node currNode) {
    long max = 0;
    for (Node node : graph.get(currNode)) {
      if (!node.equals(earlierNode)){
        max = Math.max(max, 1 + longestPath(graph, currNode, node));
      }
    }
    return max;
  }

  private static Map<Node, List<Node>> createGraph(Node[][] map) {
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
    if (node.isSlope()){
      switch (node.getSlopeDirection()) {
        case NORTH -> addNodeFromMap(adjacency, map, node.getRow() - 1, node.getCol(), Node.Direction.NORTH);
        case EAST -> addNodeFromMap(adjacency, map, node.getRow(), node.getCol() + 1, Node.Direction.EAST);
        case SOUTH -> addNodeFromMap(adjacency, map, node.getRow() + 1, node.getCol(), Node.Direction.SOUTH);
        case WEST -> addNodeFromMap(adjacency, map, node.getRow(), node.getCol() - 1, Node.Direction.WEST);
      }
    } else {
      addNodeFromMap(adjacency, map, node.getRow() + 1, node.getCol(), Node.Direction.SOUTH);
      addNodeFromMap(adjacency, map, node.getRow() - 1, node.getCol(), Node.Direction.NORTH);
      addNodeFromMap(adjacency, map, node.getRow(), node.getCol() + 1, Node.Direction.EAST);
      addNodeFromMap(adjacency, map, node.getRow(), node.getCol() - 1, Node.Direction.WEST);
    }
    return adjacency;
  }

  private static void addNodeFromMap(List<Node> list, Node[][] map, int row, int col, Node.Direction direction) {
    if ((row >= 0) && (row < map.length) && (col >= 0) && (col < map[row].length)) {
      if (map[row][col] != null){
        if (map[row][col].isSlope()) {
          if (!direction.equals(map[row][col].getOpposite())) {
            list.add(map[row][col]);
          }
        } else {
          list.add(map[row][col]);
        }
      }
    }
  }

  private static Node[][] createMap(List<String> input) {
    Node[][] map = new Node[input.size()][input.get(0).length()];
    for (int row = 0; row < input.size(); row++) {
      for (int col = 0; col < input.get(row).length(); col++) {
        switch (input.get(row).charAt(col)) {
          case '.' -> map[row][col] = new Node(row, col, 1);
          case '>' -> map[row][col] = new Node(row, col, 1, true, Node.Direction.EAST);
          case 'v' -> map[row][col] = new Node(row, col, 1, true, Node.Direction.SOUTH);
          case '<' -> map[row][col] = new Node(row, col, 1, true, Node.Direction.WEST);
          case '^' -> map[row][col] = new Node(row, col, 1, true, Node.Direction.NORTH);
          default -> map[row][col] = null;
        }
      }
    }
    return map;
  }

  public static void part2(String inputPath) {
    List<String> input = AocUtils.readInputDataToList(inputPath);
    Node[][] map = createMap(input);

    // remove slope from part 1
    for (Node[] nodes : map) {
      for (Node node : nodes) {
        if (node != null) {
          node.setIsSlope(false);
        }
      }
    }

    // create graphs
    Map<Node, List<Node>> graph = createGraph(map);
    Map<Junction, List<Edge>> junctionGraph = createJunctionGraph(graph);


    // add start and end junction
    Junction start = new Junction(0, 1);
    Junction end = new Junction(map.length - 1, map[0].length - 2);
    List<Edge> startEdges = new ArrayList<>();
    List<Edge> endEdges = new ArrayList<>();

    for (Junction junction : junctionGraph.keySet()) {
      List<Edge> edges = junctionGraph.get(junction);
      for (Edge edge : edges){
        if (edge.junction().equals(start)) {
          startEdges.add(new Edge(edge.cost(), junction));
        } else if (edge.junction().equals(end)) {
          endEdges.add(new Edge(edge.cost(), junction));
        }
      }
    }

    junctionGraph.put(start, startEdges);
    junctionGraph.put(end, endEdges);


    // calculate longest path
    long result = longestPath2(junctionGraph, start, start, new HashSet<>(), end);
    System.out.println("Day " + DAY + " part 2 result: " + result);
  }

  private static long longestPath2(Map<Junction, List<Edge>> graph, Junction earlierNode, Junction currNode, Set<Junction> visited, Junction end) {
    visited.add(earlierNode);

    if (currNode.equals(end)) {
      return 0;
    }

    long max = 0;
    for (Edge edge : graph.get(currNode)) {
      if (!edge.junction().equals(earlierNode) && !visited.contains(edge.junction())){
        max = Math.max(max, edge.cost() + longestPath2(graph, currNode, edge.junction(), new HashSet<>(visited), end));
      }
    }

    if (max == 0 && !visited.contains(end)) {
      return Long.MIN_VALUE;
    }

    return max;
  }

  public static Map<Junction, List<Edge>> createJunctionGraph(Map<Node, List<Node>> nodeGraph) {
    Map<Junction, List<Edge>> graph = new HashMap<>();

    for (Node node : nodeGraph.keySet()) {
      List<Node> adjacent = nodeGraph.get(node);
      if (adjacent.size() > 2) {
        Junction junction = new Junction(node.getRow(), node.getCol());
        List<Edge> edges = new ArrayList<>();
        for (Node adjacentNode : adjacent) {
          edges.add(getEdge(nodeGraph, node, adjacentNode));
        }
        graph.put(junction, edges);
      }
    }

    return graph;
  }

  private static Edge getEdge(Map<Node, List<Node>> nodeGraph, Node oldNode, Node currentNode) {
    List<Node> adjacent = new ArrayList<>(nodeGraph.get(currentNode));
    int cost = 1;
    while (adjacent.size() == 2) {
      adjacent.remove(oldNode);
      oldNode = currentNode;
      currentNode = adjacent.getFirst();
      adjacent = new ArrayList<>(nodeGraph.get(currentNode));
      cost += 1;
    }
    return new Edge(cost, new Junction(currentNode.getRow(), currentNode.getCol()));
  }


  public record Junction(int row, int col){}
  public record Edge(int cost, Junction junction){}

  public static void main(String[] args) {
    part1("solutions/y23/d" + DAY + "/input.txt");
    part2("solutions/y23/d" + DAY + "/input.txt");
  }
}
