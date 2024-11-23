package y23.d21;

import java.util.LinkedList;
import java.util.List;
import java.util.Objects;

public class Node {
  private final int row;
  private final int col;
  private final int cost;
  private int distance = Integer.MAX_VALUE;
  private List<Node> shortestPath = new LinkedList<>();

  public Node(int row, int col, int cost) {
    this.row = row;
    this.col = col;
    this.cost = cost;
  }

  public int getRow() {
    return this.row;
  }

  public int getCol() {
    return this.col;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) return true;
    if (o == null || getClass() != o.getClass()) return false;
    Node node = (Node) o;
    return getRow() == node.getRow() && getCol() == node.getCol();
  }

  @Override
  public int hashCode() {
    return Objects.hash(getRow(), getCol());
  }

  public int getCost() {
    return cost;
  }

  public int getDistance() {
    return distance;
  }

  public void setDistance(int distance) {
    this.distance = distance;
  }

  public List<Node> getShortestPath() {
    return shortestPath;
  }

  public void setShortestPath(List<Node> shortestPath) {
    this.shortestPath = shortestPath;
  }

  @Override
  public String toString() {
    return "Node{" +
            "row=" + row +
            ", col=" + col +
            ", distance=" + distance +
            '}';
  }

  protected enum Direction {
    NORTH,
    EAST,
    SOUTH,
    WEST
  }
}
