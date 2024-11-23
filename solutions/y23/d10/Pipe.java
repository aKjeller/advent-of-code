package y23.d10;

import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;
import java.util.stream.Stream;

public class Pipe {
  private List<Direction> connections = new ArrayList<>();
  private boolean isStart = false;
  private Direction connectedWith = null;
  private boolean done = false;

  public Pipe(Direction a, Direction b) {
    connections = Stream.of(a, b).collect(Collectors.toCollection(ArrayList::new));
  }

  public Pipe(boolean isStart) {
    if (isStart) {
      connections = Stream.of(Direction.NORTH, Direction.EAST, Direction.SOUTH, Direction.WEST).collect(Collectors.toCollection(ArrayList::new));
    }
    this.isStart = isStart;
  }

  public boolean isStart() {
    return this.isStart;
  }

  public List<Direction> getConnections() {
    return this.connections;
  }

  public void setConnections(List<Direction> connections) {
    this.connections = connections;
  }

  public Direction getConnectedWith() {
    return this.connectedWith;
  }

  public boolean contains(Direction direction) {
    if (isStart || done) {
      return false;
    }

    return connections.contains(direction);
  }

  public void connectFrom(Direction direction) {
    this.done = true;
    switch (direction) {
      case NORTH -> this.connectedWith = Direction.SOUTH;
      case EAST -> this.connectedWith = Direction.WEST;
      case SOUTH -> this.connectedWith = Direction.NORTH;
      case WEST -> this.connectedWith = Direction.EAST;
    }

    switch (direction) {
      case NORTH -> this.connections.remove(Direction.SOUTH);
      case EAST -> this.connections.remove(Direction.WEST);
      case SOUTH -> this.connections.remove(Direction.NORTH);
      case WEST -> this.connections.remove(Direction.EAST);
    }
  }

  public enum Direction {
    NORTH,
    EAST,
    SOUTH,
    WEST
  }
}
