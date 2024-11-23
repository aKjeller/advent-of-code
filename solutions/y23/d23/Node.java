package y23.d23;

import java.util.LinkedList;
import java.util.List;
import java.util.Objects;

public class Node implements Cloneable{
  private final int row;
  private final int col;
  private final int cost;
  private int distance = Integer.MAX_VALUE;
  private List<Node> longestPath = new LinkedList<>();
  private boolean isSlope = false;
  private Direction slopeDirection;

  public Node(int row, int col, int cost) {
    this.row = row;
    this.col = col;
    this.cost = cost;
  }

  public Node(int row, int col, int cost, boolean isSlope, Direction slopeDirection) {
    this.row = row;
    this.col = col;
    this.cost = cost;
    this.isSlope = isSlope;
    this.slopeDirection = slopeDirection;
  }

  public boolean isSlope() {
    return isSlope;
  }

  public void setIsSlope(boolean isSlope) {
    this.isSlope = isSlope;
  }

  public Direction getSlopeDirection() {
    return slopeDirection;
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

  public List<Node> getLongestPath() {
    return longestPath;
  }

  public void setLongestPath(List<Node> longestPath) {
    this.longestPath = longestPath;
  }

  @Override
  public String toString() {
    return "Node{" +
            "row=" + row +
            ", col=" + col +
            ", distance=" + distance +
            '}';
  }

  public Direction getOpposite() {
    return switch (this.getSlopeDirection()) {
      case NORTH -> Direction.SOUTH;
      case EAST -> Direction.WEST;
      case SOUTH -> Direction.NORTH;
      case WEST -> Direction.EAST;
    };
  }

  @Override
  public Node clone() {
    return new Node(this.getRow(), this.getCol(), this.getCost());
  }

  protected enum Direction {
    NORTH,
    EAST,
    SOUTH,
    WEST
  }
}
