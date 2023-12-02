import org.junit.Test;

import java.util.Map;
import java.util.TreeMap;

public class ScratchTests {

    @Test
    public void testTreeMapOrdering() {
        Map<Integer, Integer> map = new TreeMap<>();
        map.put(1, 5);
        map.put(3, 4);
        map.put(2, 6);

        System.out.println(map.values());
    }

}
