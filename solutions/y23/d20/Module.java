package y23.d20;

import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

public abstract class Module {

  public abstract List<Signal> getSignalsForPulse(String source, Pulse pulse);
  private final String id;
  private final List<String> outputs;

  public Module(String id, List<String> outputs) {
    this.id = id;
    this.outputs = outputs;
  }

  public String getId() {
    return id;
  }

  public List<String> getOutputs() {
    return outputs;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) return true;
    if (o == null || getClass() != o.getClass()) return false;
    Module module = (Module) o;
    return Objects.equals(id, module.id);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id);
  }

}
