package y23.d18;

import java.util.Objects;

public class Point {
  private final long row;
  private final long col;

  protected Point(long row, long col) {
    this.row = row;
    this.col = col;
  }

  public long getRow() {
    return row;
  }

  public long getCol() {
    return col;
  }

  public int getRowInt() {
    return (int) row;
  }

  public int getColInt() {
    return (int) col;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) return true;
    if (o == null || getClass() != o.getClass()) return false;
    Point point = (Point) o;
    return getRow() == point.getRow() && getCol() == point.getCol();
  }

  @Override
  public int hashCode() {
    return Objects.hash(getRow(), getCol());
  }

  @Override
  public String toString() {
    return "Point{" +
            "row=" + row +
            ", col=" + col +
            '}';
  }
}
