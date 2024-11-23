package utilities.java;

import java.io.File;
import java.io.IOException;
import java.nio.charset.StandardCharsets;
import java.nio.file.Files;
import java.util.List;
import java.util.regex.MatchResult;
import java.util.regex.Pattern;
import java.util.stream.Stream;

public class AocUtils {
  public static Pattern numberPattern = Pattern.compile("-?\\d+");
  public static List<String> readInputDataToList(String inputPath) {
    File inputFile = new File(inputPath);
    try {
      return Files.readAllLines(inputFile.toPath(), StandardCharsets.UTF_8);
    } catch (IOException e) {
      throw new RuntimeException(e);
    }
  }

  public static Stream<String> readInputDataToStream(String inputPath) {
    File inputFile = new File(inputPath);
    try {
      return Files.lines(inputFile.toPath(), StandardCharsets.UTF_8);
    } catch (IOException e) {
      throw new RuntimeException(e);
    }
  }

  public static List<Integer> createListOfIntegersFromString(String s) {
    return numberPattern.matcher(s).results().map(MatchResult::group).map(Integer::parseInt).toList();
  }

  public static List<Long> createListOfLongsFromString(String s) {
    return numberPattern.matcher(s).results().map(MatchResult::group).map(Long::parseLong).toList();
  }

}
