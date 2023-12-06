package day03_gearratios;

import java.io.BufferedReader;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Map.Entry;
import java.util.Set;

class Main {

  private static final String INPUT = "inputs/day03_gear-ratios.txt";

  public static void main(String[] args) {
    testPartOneExample();
    solvePartOne();

//    part 2 == 73074886
  }

  private static void testPartOneExample() {
    List<String> lines = Arrays.asList(PART_ONE_EXAMPLE_INPUT.split("\\R"));
    assert 4361 == sumPartNumbers(lines);
  }

  private static void solvePartOne() {
    List<String> lines = readFile(INPUT);
    long sum = sumPartNumbers(lines);
    if (sum != 527369) {
      throw new AssertionError("Wrong answer: " + sum);
      //TODO fix other asserts
    }
    System.out.println("Part 1 Input Sum: " + sum);
  }

  private static long sumPartNumbers(List<String> lines) {
    Map<String, Coordinates> symbols = new HashMap<>();
    Map<String, NumberCoordinates> numbers = new HashMap<>();
    long maxRow = lines.size() - 1L;
    Long maxColumn = null;
    for (int row = 0; row <= maxRow; row++) {
      String line = lines.get(row);
      if (maxColumn == null) {
        maxColumn = (long) line.length() - 1;
      }
      symbols.putAll(findSymbols(line, row));
      numbers.putAll(findNumbers(line, row));
    }

    List<Long> partNumbers = new ArrayList<>();
    for (Entry<String, NumberCoordinates> numberEntry : numbers.entrySet()) {
      Set<String> ac = numberEntry.getValue().getAdjoiningCoordinates(maxRow, maxColumn);
      for (String s : ac) {
        if (symbols.containsKey(s)) {
          partNumbers.add(numberEntry.getValue().getPartNumber());
          break;
        }
      }
    }
    return partNumbers.stream().reduce(0L, Long::sum);
  }

  private static Map<String, Coordinates> findSymbols(String line, long row) {
    Map<String, Coordinates> symbols = new HashMap<>();
    for (int i = 0; i < line.length(); i++) {
      char c = line.charAt(i);
      if (c != '.' && !Character.isDigit(c)) {
        var coordinates = new Coordinates(row, i);
        symbols.put(coordinates.getAbbreviation(), coordinates);
      }
    }
    return symbols;
  }

  private static Map<String, NumberCoordinates> findNumbers(String line, long row) {
    Map<String, NumberCoordinates> numbers = new HashMap<>();
    List<Integer> digits = new ArrayList<>();
    boolean foundNumber = false;
    for (int i = 0; i < line.length(); i++) {
      char c = line.charAt(i);
      if (Character.isDigit(c)) {
        foundNumber = true;
        digits.add(Integer.parseInt(String.valueOf(c)));
      } else if (foundNumber) {
        var coordinates = new NumberCoordinates(row, (long) i - digits.size(), conjoin(digits));
        numbers.put(coordinates.getAbbreviation(), coordinates);
        digits.clear();
        foundNumber = false;
      }
    }
    return numbers;
  }

  private static long conjoin(List<Integer> digits) {
    long number = 0;
    for (int digit : digits) {
      number = number * 10 + digit;
    }
    return number;
  }

  private static List<String> readFile(String filepath) {
    InputStream inputStream = Main.class.getClassLoader().getResourceAsStream(filepath);
    assert inputStream != null;
    return new BufferedReader(new InputStreamReader(inputStream)).lines().toList();
  }

  private static final String PART_ONE_EXAMPLE_INPUT = """
      467..114..
      ...*......
      ..35..633.
      ......#...
      617*......
      .....+.58.
      ..592.....
      ......755.
      ...$.*....
      .664.598..
      """;

}
