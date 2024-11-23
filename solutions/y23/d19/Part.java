package y23.d19;

import java.util.Objects;

public class Part {
  private int x;
  private int m;
  private int a;
  private int s;

  public Part(int x, int m, int a, int s) {
    this.x = x;
    this.m = m;
    this.a = a;
    this.s = s;
  }

  public Part copy(){
    return new Part(this.x, this.m, this.a, this.s);
  }

  public int getX() {
    return x;
  }

  public void setX(int x) {
    this.x = x;
  }

  public int getM() {
    return m;
  }

  public void setM(int m) {
    this.m = m;
  }

  public int getA() {
    return a;
  }

  public void setA(int a) {
    this.a = a;
  }

  public int getS() {
    return s;
  }

  public void setS(int s) {
    this.s = s;
  }

}
