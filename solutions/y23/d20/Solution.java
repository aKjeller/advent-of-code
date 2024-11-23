package y23.d20;

import utilities.java.AocUtils;

import java.math.BigInteger;
import java.util.*;

public class Solution {
  private static final String DAY = "20";

  private static long lowPulses = 0;
  private static long highPulses = 0;
  private static long[] loopSizes = new long[4];


  public static void part1(String inputPath) {
    List<String> input = AocUtils.readInputDataToList(inputPath);
    Map<String, Module> modules = new HashMap<>();

    for (String line : input) {
      String[] moduleDefinition = line.split(" -> ");
      switch (moduleDefinition[0].charAt(0)) {
        case 'b' -> modules.put(moduleDefinition[0], new BroadcastModule(moduleDefinition[0], Arrays.stream(moduleDefinition[1].split(", ")).toList()));
        case '%' -> modules.put(moduleDefinition[0].substring(1), new FlipFlopModule(moduleDefinition[0], Arrays.stream(moduleDefinition[1].split(", ")).toList()));
        case '&' -> modules.put(moduleDefinition[0].substring(1), new ConjunctionModule(moduleDefinition[0], Arrays.stream(moduleDefinition[1].split(", ")).toList()));

      }
    }

    List<Module> outputs = new ArrayList<>();
    for (Module module : modules.values()) {
      for (String destination : module.getOutputs()) {
        if (modules.get(destination) instanceof ConjunctionModule cm) {
          cm.addInput(module.getId());
        }
        if (!modules.containsKey(destination)) {
          outputs.add(new BroadcastModule(destination, new ArrayList<>()));
        }
      }
    }

    for (Module output : outputs) {
      modules.put(output.getId(), output);
    }


    for (int i = 0; i < 1000; i++) {
      pressButton(modules);
    }

    long result = lowPulses * highPulses;
    System.out.println("Day " + DAY + " part 1 result: " + result);
  }


  private static void pressButton(Map<String, Module> modules) {
    Queue<Signal> queue = new ArrayDeque<>();
    queue.add(new Signal("Button", "broadcaster", Pulse.LOW));


    while (!queue.isEmpty()) {
      Signal signal = queue.poll();
      List<Signal> signals = modules.get(signal.destination()).getSignalsForPulse(signal.source(), signal.pulse());
      for (Signal newSignal : signals) {
        queue.add(newSignal);
        //System.out.println(signal.destination() + " " + newSignal.pulse() + " " + newSignal.destination());
      }
      if (signal.pulse().equals(Pulse.LOW)) {
        lowPulses += 1;
      } else {
        highPulses += 1;
      }
    }
  }

  public static void part2(String inputPath) {
    List<String> input = AocUtils.readInputDataToList(inputPath);
    Map<String, Module> modules = new HashMap<>();

    for (String line : input) {
      String[] moduleDefinition = line.split(" -> ");
      switch (moduleDefinition[0].charAt(0)) {
        case 'b' -> modules.put(moduleDefinition[0], new BroadcastModule(moduleDefinition[0], Arrays.stream(moduleDefinition[1].split(", ")).toList()));
        case '%' -> modules.put(moduleDefinition[0].substring(1), new FlipFlopModule(moduleDefinition[0], Arrays.stream(moduleDefinition[1].split(", ")).toList()));
        case '&' -> modules.put(moduleDefinition[0].substring(1), new ConjunctionModule(moduleDefinition[0], Arrays.stream(moduleDefinition[1].split(", ")).toList()));

      }
    }

    List<Module> outputs = new ArrayList<>();
    for (Module module : modules.values()) {
      for (String destination : module.getOutputs()) {
        if (modules.get(destination) instanceof ConjunctionModule cm) {
          cm.addInput(module.getId());
        }
        if (!modules.containsKey(destination)) {
          outputs.add(new OutputModule(destination, new ArrayList<>()));
        }
      }
    }

    for (Module output : outputs) {
      modules.put(output.getId(), output);
    }

    long loops = 0;
    boolean isFinished = false;
    while (!isFinished) {
      loops += 1;
      pressButton2(modules, loops);
      isFinished = Arrays.stream(loopSizes).allMatch(l -> l > 0);
    }

    long result = Arrays.stream(loopSizes).reduce(1L, (a, b) -> lcm(BigInteger.valueOf(a), BigInteger.valueOf(b)));

    System.out.println("Day " + DAY + " part 2 result: " + result);

  }

  private static void pressButton2(Map<String, Module> modules, long loops) {
    Queue<Signal> queue = new ArrayDeque<>();
    queue.add(new Signal("Button", "broadcaster", Pulse.LOW));

    ConjunctionModule finalA = (ConjunctionModule) modules.get("mf");
    ConjunctionModule finalB = (ConjunctionModule) modules.get("zp");
    ConjunctionModule finalC = (ConjunctionModule) modules.get("jn");
    ConjunctionModule finalD = (ConjunctionModule) modules.get("ph");

    while (!queue.isEmpty()) {
      Signal signal = queue.poll();
      List<Signal> signals = modules.get(signal.destination()).getSignalsForPulse(signal.source(), signal.pulse());
      for (Signal newSignal : signals) {
        queue.add(newSignal);
        //System.out.println(signal.destination() + " " + newSignal.pulse() + " " + newSignal.destination());
      }
      if (finalA.getInputStates().values().stream().allMatch(p -> p.equals(Pulse.HIGH))) {
        if (loopSizes[0] == 0) {
          loopSizes[0] = loops;
        }
      }

      if (finalB.getInputStates().values().stream().allMatch(p -> p.equals(Pulse.HIGH))) {
        if (loopSizes[1] == 0) {
          loopSizes[1] = loops;
        }
      }

      if (finalC.getInputStates().values().stream().allMatch(p -> p.equals(Pulse.HIGH))) {
        if (loopSizes[2] == 0) {
          loopSizes[2] = loops;
        }
      }

      if (finalD.getInputStates().values().stream().allMatch(p -> p.equals(Pulse.HIGH))) {
        if (loopSizes[3] == 0) {
          loopSizes[3] = loops;
        }
      }
    }
  }

  public static long lcm(BigInteger number1, BigInteger number2) {
    BigInteger gcd = number1.gcd(number2);
    BigInteger absProduct = number1.multiply(number2).abs();
    return absProduct.divide(gcd).longValue();
  }

  public static void main(String[] args) {
    part1("solutions/y23/d" + DAY + "/input.txt");
    part2("solutions/y23/d" + DAY + "/input.txt");
  }
}
