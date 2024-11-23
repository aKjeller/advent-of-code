package y23.d11;

public class Point {
  private final long a;
  private final long b;

  public Point(long a, long b) {
    this.a = a;
    this.b = b;
  }

  public long getA() {
    return a;
  }

  public long getB() {
    return b;
  }

  public long distanceToPoint(Point other) {
    return Math.abs(this.getA() - other.getA()) + Math.abs(this.getB() - other.getB());
  }

  @Override
  public String toString() {
    return "Point{" +
            "a=" + a +
            ", b=" + b +
            '}';
  }
}
