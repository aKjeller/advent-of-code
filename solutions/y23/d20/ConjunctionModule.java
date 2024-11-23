package y23.d20;

import java.util.*;

public class ConjunctionModule extends Module {

  private final Map<String, Pulse> inputStates = new LinkedHashMap<>();
  @Override
  public List<Signal> getSignalsForPulse(String source, Pulse pulse) {
    List<Signal> signals = new ArrayList<>();

    inputStates.put(source, pulse);

    boolean allHigh = inputStates.values().stream().allMatch(p -> p.equals(Pulse.HIGH));
    for (String output : this.getOutputs()) {
      signals.add(new Signal(this.getId(), output, allHigh ? Pulse.LOW : Pulse.HIGH));
    }

    return signals;
  }

  public void addInput(String input) {
    this.inputStates.put(input, Pulse.LOW);
  }

  public Map<String, Pulse> getInputStates() {
    return this.inputStates;
  }

  public ConjunctionModule(String id, List<String> outputs) {
    super(id.substring(1), outputs);
  }

  @Override
  public String toString() {
    return "ConjunctionModule{" +
            "inputStates=" + inputStates +
            '}';
  }
}
