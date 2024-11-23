package y23.d16;

import javax.swing.*;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

public class Beam {
  private int row;
  private int col;
  private Direction entryDirection;
  private final List<Beam> next = new ArrayList<>();

  public Beam(int row, int col) {
    this.row = row;
    this.col = col;
  }

  public Beam(int row, int col, Direction direction) {
    switch (direction) {
      case NORTH -> {
        this.row = row - 1;
        this.col = col;
      }
      case EAST -> {
        this.row = row;
        this.col = col + 1;
      }
      case SOUTH -> {
        this.row = row + 1;
        this.col = col;
      }
      case WEST -> {
        this.row = row;
        this.col = col - 1;
      }
    }
  }

  public int getRow() {
    return row;
  }

  public int getCol() {
    return col;
  }

  public Direction getEntryDirection() {
    return this.entryDirection;
  }

  public void setEntryDirection(Direction entryDirection) {
    this.entryDirection = entryDirection;
  }

  public List<Beam> getNext() {
    return next;
  }

  public void addNext(Beam next) {
    this.next.add(next);
  }

  public boolean hasNext() {
    return !this.next.isEmpty();
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) return true;
    if (o == null || getClass() != o.getClass()) return false;
    Beam beam = (Beam) o;
    return getRow() == beam.getRow() && getCol() == beam.getCol() && getEntryDirection() == beam.getEntryDirection();
  }

  @Override
  public int hashCode() {
    return Objects.hash(getRow(), getCol(), getEntryDirection());
  }

  public enum Direction {
    NORTH,
    EAST,
    SOUTH,
    WEST
  }
}
