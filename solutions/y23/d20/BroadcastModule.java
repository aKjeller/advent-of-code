package y23.d20;

import java.util.ArrayList;
import java.util.List;

public class BroadcastModule extends Module {
  @Override
  public List<Signal> getSignalsForPulse(String source, Pulse pulse) {
    List<Signal> signals = new ArrayList<>();
    for (String output : this.getOutputs()) {
      signals.add(new Signal(this.getId(), output, pulse));
    }
    return signals;
  }

  public BroadcastModule(String id, List<String> outputs) {
    super(id, outputs);
  }
}
