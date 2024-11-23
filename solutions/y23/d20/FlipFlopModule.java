package y23.d20;

import java.util.ArrayList;
import java.util.List;

public class FlipFlopModule extends Module {
  private boolean ON = false;

  @Override
  public List<Signal> getSignalsForPulse(String source, Pulse pulse) {
    List<Signal> signals = new ArrayList<>();
    if (pulse.equals(Pulse.LOW)) {
      Pulse newPulse;
      if (this.ON) {
        this.ON = false;
        newPulse = Pulse.LOW;
      } else {
        this.ON = true;
        newPulse = Pulse.HIGH;
      }
      for (String output : this.getOutputs()) {
        signals.add(new Signal(this.getId(), output, newPulse));
      }
    }
    return signals;
  }
  public FlipFlopModule(String id, List<String> outputs) {
    super(id.substring(1), outputs);
  }
}
