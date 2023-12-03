package day02_cubeconundrum;

import java.io.BufferedReader;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.util.Arrays;
import java.util.List;
import java.util.Map;

public class Main {

  private static final String INPUT = "inputs/day02_cube-conundrum.txt";

  public static void main(String[] args) {
    testPartOneExample();
    solvePartOne();

    testPartTwoExample();
    solvePartTwo();
  }

  private static void testPartOneExample() {
    Long idSum = computePartOne(readExample());
    assert idSum == 8;
  }

  private static void solvePartOne() {
    Long idSum = computePartOne(readInput());
//        System.out.println("Part 1 Input Sum: " + idSum);
    assert idSum == 2439;
  }

  private static Long computePartOne(List<String> lines) {
    return lines.stream()
                .map(Game::new)
                .filter(g -> g.isValid(PART_ONE_LIMITS))
                .map(Game::getId)
                .reduce(0L, Long::sum);
  }

  private static void testPartTwoExample() {
    List<Game> games = readExample().stream().map(Game::new).toList();
    games.forEach(game -> {
      assert game.getPower() == PART_TWO_EXPECTED_POWERS.get(game.getId());
    });
    var powerSum = games.stream().map(Game::getPower).reduce(0L, Long::sum);
    assert powerSum == 2286;
  }

  private static void solvePartTwo() {
    Long powerSum = readInput().stream().map(Game::new).map(Game::getPower).reduce(0L, Long::sum);
//        System.out.println("Part 2 Input Sum: " + powerSum);
    assert powerSum == 63711;
  }

  private static List<String> readExample() {
    return Arrays.asList(EXAMPLE_INPUT.split("\\R"));
  }

  private static List<String> readInput() {
    InputStream inputStream = day01_trebuchet.Main.class.getClassLoader()
                                                        .getResourceAsStream(INPUT);
    assert inputStream != null;
    return new BufferedReader(new InputStreamReader(inputStream)).lines().toList();
  }

  private static final Map<String, Integer> PART_ONE_LIMITS = Map.of("red", 12, "green", 13, "blue",
      14);

  private static final Map<Long, Long> PART_TWO_EXPECTED_POWERS = Map.of(1L, 48L, 2L, 12L, 3L,
      1560L, 4L, 630L, 5L, 36L);

  private static final String EXAMPLE_INPUT = """
      Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
      Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
      Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
      Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
      Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
      """;
}
