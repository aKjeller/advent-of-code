package y23.d22;

public class Block implements Cloneable{
  private String label;
  private int x, y, z;

  public Block(String label, int x, int y, int z) {
    this.label = label;
    this.x = x;
    this.y = y;
    this.z = z;
  }

  public void setLabel(String label) {
    this.label = label;
  }

  public String getLabel() {
    return label;
  }

  public int getX() {
    return x;
  }

  public void setX(int x) {
    this.x = x;
  }

  public int getY() {
    return y;
  }

  public void setY(int y) {
    this.y = y;
  }

  public int getZ() {
    return z;
  }

  public void setZ(int z) {
    this.z = z;
  }

  @Override
  public String toString() {
    return "Block{" +
            "label='" + label + '\'' +
            ", x=" + x +
            ", y=" + y +
            ", z=" + z +
            '}';
  }

  @Override
  public Block clone() {
    return new Block(this.getLabel(), this.getX(), this.getY(), this.getZ());
  }
}
