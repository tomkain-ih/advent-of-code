import static junit.framework.TestCase.assertEquals;

import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.TreeMap;
import org.junit.Test;

public class ScratchTests {

  @Test
  public void testTreeMapOrdering() {
    Map<Integer, Integer> map = new TreeMap<>();
    map.put(1, 5);
    map.put(3, 4);
    map.put(2, 6);

    assertEquals(List.of(5, 6, 4), List.copyOf(map.values()));
  }

  @Test
  public void testStringSplit() {
    String input = "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green";
    String[] colonSplit = input.split(":");
//        System.out.println(Long.parseLong(colonSplit[0].split("Game ")[1].trim()));

//        System.out.println(colonSplit[1]);

    for (String round : colonSplit[1].split(";")) {
      String[] commaSplit = round.trim().split(",");
      Map<String, Integer> outcomeMap = new HashMap<>();
      for (String outcome : commaSplit) {
        String[] spaceSplit = outcome.trim().split(" ");
        int count = Integer.parseInt(spaceSplit[0]);
        String color = spaceSplit[1];
        outcomeMap.put(color, count);
      }
      System.out.println(outcomeMap);
    }
  }

}
