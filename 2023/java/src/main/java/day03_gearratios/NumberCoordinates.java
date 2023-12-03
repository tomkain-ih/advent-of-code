package day03_gearratios;

import java.util.HashSet;
import java.util.Set;
import lombok.AccessLevel;
import lombok.Getter;
import lombok.RequiredArgsConstructor;

@Getter
@RequiredArgsConstructor(access = AccessLevel.PUBLIC)
class NumberCoordinates implements CoordinatesInterface {

  private final long row;
  private final long column;
  private final long partNumber;

  @Override
  public Set<String> getAdjoiningCoordinates(long maxRow, long maxColumn) {
    Set<String> ac = new HashSet<>(
        CoordinatesInterface.super.getAdjoiningCoordinates(maxRow, maxColumn));
    for (int i = 1; i <= String.valueOf(partNumber).length(); i++) {
      ac.addAll(new Coordinates(row, column + i).getAdjoiningCoordinates(maxRow, maxColumn));
    }
    return ac;
  }


}
