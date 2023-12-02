package day02_cubeconundrum;

import lombok.AccessLevel;
import lombok.Getter;
import lombok.RequiredArgsConstructor;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

@Getter
@RequiredArgsConstructor(access = AccessLevel.PRIVATE)
public class Game {

    private final long id;
    private final List<Map<String, Integer>> rounds = new ArrayList<>();
    private final Map<String, Integer> maxCountByColor = new HashMap<>();
    private final long power;

    Game(String rawGame) {
        String[] colonSplit = rawGame.split(":");
        this.id = Long.parseLong(colonSplit[0].split("Game ")[1].trim());

        for (String round : colonSplit[1].split(";")) {
            String[] commaSplit = round.trim().split(",");
            Map<String, Integer> outcomeMap = new HashMap<>();
            for (String outcome : commaSplit) {
                String[] spaceSplit = outcome.trim().split(" ");
                int count = Integer.parseInt(spaceSplit[0]);
                String color = spaceSplit[1];
                outcomeMap.put(color, count);

                maxCountByColor.merge(color, count, Math::max);
            }
            this.rounds.add(outcomeMap);
        }

        this.power = maxCountByColor.values().stream().reduce(1, (a, b) -> a * b).longValue();
    }

    boolean isValid(Map<String, Integer> limits) {
        for (Map.Entry<String, Integer> limit : limits.entrySet()) {
            if (maxCountByColor.containsKey(limit.getKey()) && maxCountByColor.get(limit.getKey()) > limit.getValue()) {
                return false;
            }
        }
        return true;
    }

}
