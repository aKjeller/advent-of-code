package y23.d20;

import java.util.ArrayList;
import java.util.List;

public class OutputModule extends Module {

  private boolean finished = false;
  public OutputModule(String id, List<String> outputs) {
    super(id, outputs);
  }

  public boolean isFinished() {
    return finished;
  }

  @Override
  public List<Signal> getSignalsForPulse(String source, Pulse pulse) {
    if (pulse.equals(Pulse.LOW)) {
      this.finished = true;
    }
    return new ArrayList<>();
  }
}
