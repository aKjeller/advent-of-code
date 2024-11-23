package y23.d25;

import utilities.java.AocUtils;

import java.util.*;

public class Solution {

  public record Edge(String a, String b){}
  private static final String DAY = "25";

  public static void part1(String inputPath) {
    List<String> input = AocUtils.readInputDataToList(inputPath);
    Set<Edge> doubleEdges = new HashSet<>();
    Map<String, Set<String>> graph = new HashMap<>();
    for (String line : input) {
      String name = line.split(": ")[0];
      String[] connections = line.split(": ")[1].split( " ");
      if (!graph.containsKey(name)) {
        graph.put(name, new HashSet<>());
      }
      for (String conn : connections) {
        if (!graph.containsKey(conn)) {
          graph.put(conn, new HashSet<>());
        }
        graph.get(name).add(conn);
        graph.get(conn).add(name);
        doubleEdges.add(new Edge(name, conn));
        doubleEdges.add(new Edge(conn, name));
      }
    }

    List<Edge> edges = new ArrayList<>();
    for (Edge edge : doubleEdges) {
      if (!edges.contains(edge)) {
        if (!edges.contains(new Edge(edge.b(), edge.a()))) {
          edges.add(edge);
        }
      }
    }

    long result = 0;



    System.out.println("Day " + DAY + " part 1 result: " + result);
  }

  private static List<Edge> stoerWagner(Map<String, Set<String>> g) {
    Map<String, Set<String>> v = new HashMap<>(g);

    List<Edge> minCuts = new ArrayList<>();
    while (v.size() > 2) {
      List<Edge> cuts = minimumCut(v);

      for (Edge cut : cuts) {
        v.get(cut.a()).remove(cut.b());
        v.get(cut.b()).remove(cut.a());
      }

      if (cuts.size() < minCuts.size()) {
        minCuts = cuts;
      }
    }

    return minCuts;
  }

  private static List<Edge> minimumCut(Map<String, Set<String>> v) {

    return null;
  }

  private static long isGraphDisconnected(Map<String, Set<String>> graph) {
    Set<String> visited = new HashSet<>();

    String outSide = null;
    for (String node : graph.keySet()) {
      if (visited.isEmpty()) {
        visitAllNodes(graph, visited, node);
      } else if (!visited.contains(node)) {
        outSide = node;
        break;
      }
    }
    Set<String> visited2 = new HashSet<>();
    visitAllNodes(graph, visited2, outSide);
    if (visited.size() + visited2.size() == graph.keySet().size()) {
      return visited.size() * visited2.size();
    } else {
      return 0;
    }
  }

  private static void visitAllNodes(Map<String, Set<String>> graph, Set<String> visited, String node) {
    if (graph.containsKey(node) && !visited.contains(node)) {
      visited.add(node);
      for (String adj : graph.get(node)) {
        visitAllNodes(graph, visited, adj);
      }
    }
  }

  public static void part2(String inputPath) {
    List<String> input = AocUtils.readInputDataToList(inputPath);

    long result = 0;
    System.out.println("Day " + DAY + " part 2 result: " + result);
  }

  public static void main(String[] args) {
    part1("solutions/y23/d" + DAY + "/input.txt");
    part2("solutions/y23/d" + DAY + "/input.txt");
  }
}
