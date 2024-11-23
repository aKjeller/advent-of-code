package y23.d17;

import java.util.LinkedList;
import java.util.List;
import java.util.Objects;

public class Node {
  private final int row;
  private final int col;
  private final int cost;
  private int distance = Integer.MAX_VALUE;
  private Direction direction;
  private int nTimesInDirection;
  private List<Node> shortestPath = new LinkedList<>();

  public Node(int row, int col, int cost) {
    this.row = row;
    this.col = col;
    this.cost = cost;
  }

  public Node(int row, int col, int cost, Direction direction, int nTimesInDirection) {
    this.row = row;
    this.col = col;
    this.cost = cost;
    this.direction = direction;
    this.nTimesInDirection = nTimesInDirection;
  }

  public int getRow() {
    return this.row;
  }

  public int getCol() {
    return this.col;
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

  public Direction getDirection() {
    return direction;
  }

  public void setDirection(Direction direction) {
    this.direction = direction;
  }

  public int getnTimesInDirection() {
    return nTimesInDirection;
  }

  public void setnTimesInDirection(int nTimesInDirection) {
    this.nTimesInDirection = nTimesInDirection;
  }

  public List<Node> getShortestPath() {
    return shortestPath;
  }

  public void setShortestPath(List<Node> shortestPath) {
    this.shortestPath = shortestPath;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) return true;
    if (o == null || getClass() != o.getClass()) return false;
    Node node = (Node) o;
    return getRow() == node.getRow() && getCol() == node.getCol() && getnTimesInDirection() == node.getnTimesInDirection() && getDirection() == node.getDirection();
  }

  @Override
  public int hashCode() {
    return Objects.hash(getRow(), getCol(), getDirection(), getnTimesInDirection());
  }

  @Override
  public String toString() {
    return "Node{" +
            "row=" + row +
            ", col=" + col +
            ", cost=" + cost +
            ", distance=" + distance +
            ", direction=" + direction +
            ", nTimesInDirection=" + nTimesInDirection +
            '}';
  }

  protected enum Direction {
    NORTH,
    EAST,
    SOUTH,
    WEST
  }
}
