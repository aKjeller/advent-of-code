package y23.d19;

import java.util.ArrayList;
import java.util.List;

public class WorkFlow {

  private List<String> rules = new ArrayList<>();
  private String end;

  public WorkFlow(String input) {
    String[] rules = input.split(",");
    for (int i = 0; i < rules.length - 1; i++) {
      String rule = rules[i];
      this.rules.add(rule);
    }
    this.setEnd(rules[rules.length - 1]);
  }

  public String evaluateRules(Part part) {
    for (String rule : this.rules) {
      if (evaluateRule(rule.split(":")[0], part)) {
        return rule.split(":")[1];
      }
    }
    return end;
  }

  private boolean evaluateRule(String rule, Part part) {
    if (rule.contains("<")) {
      int value = switch (rule.substring(0, rule.indexOf('<'))) {
        case "x" -> part.getX();
        case "m" -> part.getM();
        case "a" -> part.getA();
        case "s" -> part.getS();
        default -> throw new IllegalStateException("Unexpected value: " + rule.substring(0, rule.indexOf('<')));
      };
      return value < Integer.parseInt(rule.substring(rule.indexOf('<') + 1));
    } else {
      int value = switch (rule.substring(0, rule.indexOf('>'))) {
        case "x" -> part.getX();
        case "m" -> part.getM();
        case "a" -> part.getA();
        case "s" -> part.getS();
        default -> throw new IllegalStateException("Unexpected value: " + rule.substring(0, rule.indexOf('>')));
      };
      return value > Integer.parseInt(rule.substring(rule.indexOf('>') + 1));
    }
  }

  public List<String> getRules() {
    return rules;
  }

  public void setRules(List<String> rules) {
    this.rules = rules;
  }

  public String getEnd() {
    return end;
  }

  public void setEnd(String end) {
    this.end = end;
  }
}
