package y23.d19;

import java.util.ArrayList;
import java.util.List;

public class Node {
  private final String label;
  private final Part lowerBound;
  private final Part upperBound;
  private final List<Node> children = new ArrayList<>();

  public Node(String label, Part lowerBound, Part upperBound) {
    this.label = label;
    this.lowerBound = lowerBound;
    this.upperBound = upperBound;
  }

  public void addChildren(Node node) {
    this.children.add(node);
  }

  public String getLabel() {
    return label;
  }

  public Part getLowerBound() {
    return lowerBound;
  }

  public Part getUpperBound() {
    return upperBound;
  }

  public List<Node> getChildren() {
    return children;
  }
}
