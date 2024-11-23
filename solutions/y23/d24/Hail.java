package y23.d24;

import org.apache.commons.math3.linear.*;
import utilities.java.AocUtils;

import java.util.List;
import java.util.Objects;

public class Hail {
  private String name;
  private final double s_x; // a e
  private final double s_y; // b f
  private final double start_z;

  private final double v_x; // c g
  private final double v_y; // d h
  private final double v_z;

  public Hail(String input) {
    this.name = input;
    List<Long> numbers = AocUtils.createListOfLongsFromString(input);
    this.s_x = numbers.get(0);
    this.s_y = numbers.get(1);
    this.start_z = numbers.get(2);
    this.v_x = numbers.get(3);
    this.v_y = numbers.get(4);
    this.v_z = numbers.get(5);
  }

  public boolean intersects(Hail other, long lower, long upper) {
    // Create coefficient matrix M
    RealMatrix coefficients = MatrixUtils.createRealMatrix(new double[][]{
            {-this.v_x, other.v_x},
            {-this.v_y, other.v_y}
    });

    // Create constants vector C
    RealVector constants = new ArrayRealVector(new double[]{
            this.s_x - other.s_x,
            this.s_y - other.s_y
    });

    // Solve the system of linear equations
    DecompositionSolver solver = new LUDecomposition(coefficients).getSolver();
    RealVector solution;
    try {
      solution = solver.solve(constants);
    } catch (SingularMatrixException e) {
      return false;
    }

    // Extract t and s values
    double t = solution.getEntry(0);
    double s = solution.getEntry(1);

    if (t <= 0 || s <= 0) {
      return false;
    }

    return getX(t) >= lower && getX(t) <= upper && getY(t) >= lower && getY(t) <= upper;
  }

  public double getX(Double time) {
    return time * this.getV_x() + this.getS_x();
  }

  public double getY(Double time) {
    return time * this.getV_y() + this.getS_y();
  }

  public double getS_x() {
    return s_x;
  }

  public double getS_y() {
    return s_y;
  }

  public double getStart_z() {
    return start_z;
  }

  public double getV_x() {
    return v_x;
  }

  public double getV_y() {
    return v_y;
  }

  public double getV_z() {
    return v_z;
  }

  @Override
  public String toString() {
    return "Hail{" +
            "name='" + name + '\'' +
            '}';
  }
}
