package day01_trebuchet;

import java.io.BufferedReader;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Map;
import java.util.TreeMap;
import java.util.function.Function;

public class Main {

    private static final String INPUT = "inputs/day01_trebuchet.txt";

    public static void main(String[] args) {
        testPartOneExample();
        solvePartOne();

        testPartTwoExample();
        solvePartTwo();
    }

    private static void solvePartOne() {
        List<String> lines = readFile(INPUT);
        assert lines.size() == 1000;
        long sum = lines.stream().map(l -> decipherLineValue(l, Main::findDigitCharacters)).reduce(0L, Long::sum);

//        System.out.println("Part 1 Input Sum: " + sum);
        assert sum == 53334;
    }

    private static void solvePartTwo() {
        List<String> lines = readFile(INPUT);
        assert lines.size() == 1000;
        long sum = lines.stream().map(l -> decipherLineValue(l, Main::findDigitAndSpelledNumbers)).reduce(0L, Long::sum);

//        System.out.println("Part 2 Input Sum: " + sum);
        assert sum == 52834;
    }

    private static List<String> readFile(String filepath) {
        InputStream inputStream = Main.class.getClassLoader().getResourceAsStream(filepath);
        assert inputStream != null;
        return new BufferedReader(new InputStreamReader(inputStream)).lines().toList();
    }

    private static void testPartOneExample() {
        List<String> lines = splitLines(PART_ONE_EXAMPLE_INPUT);
        assert lines.size() == 4;
        long sum = lines.stream().map(l -> decipherLineValue(l, Main::findDigitCharacters)).reduce(0L, Long::sum);

        assert sum == 142;
    }

    private static void testPartTwoExample() {
        List<String> lines = splitLines(PART_TWO_EXAMPLE_INPUT);
        assert lines.size() == 7;
        long sum = lines.stream().map(l -> decipherLineValue(l, Main::findDigitAndSpelledNumbers)).reduce(0L, Long::sum);

        assert sum == 281;
    }

    private static List<String> splitLines(String input) {
        return Arrays.asList(input.split("\\R"));
    }

    private static long decipherLineValue(String line, Function<String, List<Long>> digitExtractor) {
        List<Long> numbers = digitExtractor.apply(line);
        if (numbers.isEmpty()) {
            throw new IllegalArgumentException("No numbers found in line: " + line);
        }

        if (numbers.size() == 1) {
            return numbers.get(0) * 10 + numbers.get(0);
        }

        return numbers.get(0) * 10 + numbers.get(numbers.size() - 1);
    }

    private static List<Long> findDigitCharacters(String line) {
        List<Long> numbers = new ArrayList<>();
        for (int i = 0; i < line.length(); i++) {
            char c = line.charAt(i);
            if (Character.isDigit(c)) {
                numbers.add(Long.parseLong(String.valueOf(c)));
            }
        }
        return numbers;
    }

    private static List<Long> findDigitAndSpelledNumbers(String line) {
        Map<Integer, Long> digitsByIndex = new TreeMap<>();
        for (int i = 0; i < line.length(); i++) {
            char c = line.charAt(i);
            if (Character.isDigit(c)) {
                digitsByIndex.put(i, Long.parseLong(String.valueOf(c)));
            }
        }
        for (String word : spelledDigits.keySet()) {
            int index = line.indexOf(word);
            while (index >= 0) {
                digitsByIndex.put(index, (long) spelledDigits.get(word));
                index = line.indexOf(word, index + 1);
            }
        }
        return new ArrayList<>(digitsByIndex.values());
    }

    private static final String PART_ONE_EXAMPLE_INPUT = """
            1abc2
            pqr3stu8vwx
            a1b2c3d4e5f
            treb7uchet
            """;

    private static final String PART_TWO_EXAMPLE_INPUT = """
            two1nine
            eightwothree
            abcone2threexyz
            xtwone3four
            4nineeightseven2
            zoneight234
            7pqrstsixteen
            """;

    private static final Map<String, Integer> spelledDigits = Map.of(
            "zero", 0,
            "one", 1,
            "two", 2,
            "three", 3,
            "four", 4,
            "five", 5,
            "six", 6,
            "seven", 7,
            "eight", 8,
            "nine", 9
    );
}
