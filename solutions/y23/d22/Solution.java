package y23.d22;

import utilities.java.AocUtils;

import java.util.*;

public class Solution {
  private static final String DAY = "22";

  public static void part1(String inputPath) {
    List<String> input = AocUtils.readInputDataToList(inputPath);
    Map<String, List<Block>> blocks = createBlocks(input);


    int xMax = Integer.MIN_VALUE;
    int yMax = Integer.MIN_VALUE;
    int zMax = Integer.MIN_VALUE;
    for (List<Block> blockList : blocks.values()) {
      for (Block block : blockList) {
        xMax = Math.max(xMax, block.getX());
        yMax = Math.max(yMax, block.getY());
        zMax = Math.max(zMax, block.getZ());
      }
    }

    Block[][][] state = new Block[xMax + 1][yMax + 1][zMax + 1];
    for (List<Block> blockList : blocks.values()) {
      for (Block block : blockList) {
        state[block.getX()][block.getY()][block.getZ()] = block;
      }
    }

    for (int z = 1; z <= zMax; z++) {
      for (int x = 0; x <= xMax; x++){
        for (int y = 0; y <= yMax; y++) {
          if (state[x][y][z] != null) {
            String label = state[x][y][z].getLabel();
            while (!hasBlockUnder(blocks, state, label)) {
              moveDownOneLevel(blocks, state, label);
            }
          }
        }
      }
    }

    long result = 0;
    for (List<Block> currentBlocks : blocks.values()){
      String label = currentBlocks.getFirst().getLabel();
      if (!hasBlockAbove(blocks, state, label)) {
        result += 1;
      } else {
        boolean blocksAboveHaveSupport = true;
        for (Block block : currentBlocks) {
          if (!blockAboveMoreThenOneSupport(blocks, state, block)) {
            blocksAboveHaveSupport = false;
          }
        }
        if (blocksAboveHaveSupport) {
          result += 1;
        }
      }
    }
    System.out.println("Day " + DAY + " part 1 result: " + result);

    result = 0;
    for (List<Block> currentBlocks : blocks.values()){
      if (canBeRemoved(blocks, state, currentBlocks)) {
        result += 1;
      }
    }

    System.out.println("Day " + DAY + " part 1 result: " + result);
  }

  private static boolean canBeRemoved(Map<String, List<Block>> blocks, Block[][][] state, List<Block> currentBlocks) {
    boolean canBeRemoved = true;
    for (Block block : currentBlocks) {

      Block above = state[block.getX()][block.getY()][block.getZ() + 1];
      if (above != null && !above.getLabel().equals(block.getLabel())) {
        state[block.getX()][block.getY()][block.getZ()] = null;
        if (!hasBlockUnder2(blocks, state, above.getLabel(), block.getLabel())) {
          state[block.getX()][block.getY()][block.getZ()] = block;
          return false;
        }
        state[block.getX()][block.getY()][block.getZ()] = block;
      }
    }
    return canBeRemoved;
  }

  private static boolean blockAboveMoreThenOneSupport(Map<String, List<Block>> blocks, Block[][][] state, Block block) {
    if (state[block.getX()][block.getY()][block.getZ() + 1] == null) {
      return true;
    }

    List<Block> above = blocks.get(state[block.getX()][block.getY()][block.getZ() + 1].getLabel());

    HashMap<String, Boolean> support = new HashMap<>();
    for (Block aboveBlock : above) {
      if (state[aboveBlock.getX()][aboveBlock.getY()][aboveBlock.getZ() - 1] != null) {
        support.put(state[aboveBlock.getX()][aboveBlock.getY()][aboveBlock.getZ() - 1].getLabel(), true);
      }
    }
    support.remove(state[block.getX()][block.getY()][block.getZ() + 1].getLabel());

    return support.size() > 1;
  }

  private static boolean hasBlockAbove(Map<String, List<Block>> allBlocks, Block[][][] state, String label) {
    List<Block> blocks = allBlocks.get(label);
    for (Block block : blocks) {
      if (state[block.getX()][block.getY()][block.getZ() + 1] != null) {
        if (!state[block.getX()][block.getY()][block.getZ() + 1].getLabel().equals(label)) {
          return true;
        }
      }
    }
    return false;
  }

  private static void moveDownOneLevel(Map<String, List<Block>> allBlocks, Block[][][] state, String label) {
    List<Block> blocks = allBlocks.get(label);
    for (Block block : blocks) {
      state[block.getX()][block.getY()][block.getZ()] = null;
      block.setZ(block.getZ() - 1);
    }

    for (Block block : blocks) {
      state[block.getX()][block.getY()][block.getZ()] = block;
    }
  }

  private static boolean hasBlockUnder2(Map<String, List<Block>> allBlocks, Block[][][] state, String label, String dontCountLabel) {
    List<Block> blocks = allBlocks.get(label);
    for (Block block : blocks) {
      if (block.getZ() == 1) {
        return true;
      }
      if (state[block.getX()][block.getY()][block.getZ() - 1] != null) {
        if (!state[block.getX()][block.getY()][block.getZ() - 1].getLabel().equals(label)) {
          if (!state[block.getX()][block.getY()][block.getZ() - 1].getLabel().equals(dontCountLabel)) {
            return true;
          }
        }
      }
    }
    return false;
  }

  private static boolean hasBlockUnder(Map<String, List<Block>> allBlocks, Block[][][] state, String label) {
    List<Block> blocks = allBlocks.get(label);
    for (Block block : blocks) {
      if (block.getZ() == 1) {
        return true;
      }
      if (state[block.getX()][block.getY()][block.getZ() - 1] != null) {
        if (!state[block.getX()][block.getY()][block.getZ() - 1].getLabel().equals(label)) {
          return true;
        }
      }
    }
    return false;
  }


  private static Map<String, List<Block>> createBlocks(List<String> input) {
    Map<String, List<Block>> blocks = new HashMap<>();

    int counter = 0;
    for (String line : input) {
      String label = getLabel(counter++);
      List<Integer> start = AocUtils.createListOfIntegersFromString(line.split("~")[0]);
      List<Integer> end = AocUtils.createListOfIntegersFromString(line.split("~")[1]);
      List<Block> block = new ArrayList<>();
      if (!Objects.equals(start.get(0), end.get(0))) {
        for (int i = start.get(0); i <= end.get(0); i++) {
          block.add(new Block(label, i, start.get(1), start.get(2)));
        }
      } else if (!Objects.equals(start.get(1), end.get(1))) {
        for (int i = start.get(1); i <= end.get(1); i++) {
          block.add(new Block(label, start.get(0), i, start.get(2)));
        }
      } else if (!Objects.equals(start.get(2), end.get(2))) {
        for (int i = start.get(2); i <= end.get(2); i++) {
          block.add(new Block(label, start.get(0), start.get(1), i));
        }
      } else {
        block.add(new Block(label, start.get(0), start.get(1), start.get(2)));
      }
      blocks.put(label, block);
    }

    return blocks;
  }

  private static String getLabel(int counter) {
    return String.format("%02x", counter);
  }

  public static void part2(String inputPath) {
    List<String> input = AocUtils.readInputDataToList(inputPath);
    Map<String, List<Block>> blocks = createBlocks(input);


    int xMax = Integer.MIN_VALUE;
    int yMax = Integer.MIN_VALUE;
    int zMax = Integer.MIN_VALUE;
    for (List<Block> blockList : blocks.values()) {
      for (Block block : blockList) {
        xMax = Math.max(xMax, block.getX());
        yMax = Math.max(yMax, block.getY());
        zMax = Math.max(zMax, block.getZ());
      }
    }

    Block[][][] state = new Block[xMax + 1][yMax + 1][zMax + 1];
    for (List<Block> blockList : blocks.values()) {
      for (Block block : blockList) {
        state[block.getX()][block.getY()][block.getZ()] = block;
      }
    }

    for (int z = 1; z <= zMax; z++) {
      for (int x = 0; x <= xMax; x++){
        for (int y = 0; y <= yMax; y++) {
          if (state[x][y][z] != null) {
            String label = state[x][y][z].getLabel();
            while (!hasBlockUnder(blocks, state, label)) {
              moveDownOneLevel(blocks, state, label);
            }
          }
        }
      }
    }


    int result = 0;
    for (List<Block> block : blocks.values()){
      Map<String, List<Block>> blocksCopy = copyBlocks(blocks);
      Block[][][] stateCopy = copyState(blocksCopy, xMax, yMax, zMax);
      removeBlockFromState(stateCopy, block);

      Map<String, Boolean> fallenBlocks = new HashMap<>();
      for (int z = 1; z <= zMax; z++) {
        for (int x = 0; x <= xMax; x++){
          for (int y = 0; y <= yMax; y++) {
            if (stateCopy[x][y][z] != null) {
              String label = stateCopy[x][y][z].getLabel();
              while (!hasBlockUnder(blocksCopy, stateCopy, label)) {
                moveDownOneLevel(blocksCopy, stateCopy, label);
                fallenBlocks.put(label, true);
              }
            }
          }
        }
      }
      result += fallenBlocks.size();
    }

    System.out.println("Day " + DAY + " part 2 result: " + result);
  }

  private static Map<String, List<Block>> copyBlocks(Map<String, List<Block>> blocks) {
    Map<String, List<Block>> blocksCopy = new HashMap<>();
    for (String copy : blocks.keySet()) {
      List<Block> parts = new ArrayList<>();
      for (Block block : blocks.get(copy)) {
        parts.add(block.clone());
      }
      blocksCopy.put(copy, parts);
    }
    return blocksCopy;
  }

  private static void removeBlockFromState(Block[][][] stateCopy, List<Block> blocks) {
    for (Block block : blocks) {
      stateCopy[block.getX()][block.getY()][block.getZ()] = null;
    }
  }

  private static Block[][][] copyState(Map<String, List<Block>> blocks, int xMax, int yMax, int zMax) {
    Block[][][] copy = new Block[xMax + 1][yMax + 1][zMax + 1];
    for (List<Block> blockList : blocks.values()) {
      for (Block block : blockList) {
        copy[block.getX()][block.getY()][block.getZ()] = block;
      }
    }
    return copy;
  }

  public static void main(String[] args) {
    part1("solutions/y23/d" + DAY + "/input.txt");
    part2("solutions/y23/d" + DAY + "/input.txt");
  }
}

// 516 low
