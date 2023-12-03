package day03_gearratios;

import java.util.HashSet;
import java.util.Set;
import java.util.stream.Collectors;

interface CoordinatesInterface {

  long getRow();

  long getColumn();

  default String getAbbreviation() {
    return String.format("%d,%d", getRow(), getColumn());
  }

  default Set<String> getAdjoiningCoordinates(long maxRow, long maxColumn) {
    Set<Coordinates> adjoining = new HashSet<>();
    if (getColumn() > 0) {
      adjoining.add(new Coordinates(getRow(), getColumn() - 1));
    }
    if (getColumn() < maxColumn) {
      adjoining.add(new Coordinates(getRow(), getColumn() + 1));
    }
    if (getRow() > 0) {
      adjoining.add(new Coordinates(getRow() - 1, getColumn()));
    }
    if (getRow() < maxRow) {
      adjoining.add(new Coordinates(getRow() + 1, getColumn()));
    }
    if (getRow() < maxRow && getColumn() > 0) {
      adjoining.add(new Coordinates(getRow() + 1, getColumn() - 1));
    }
    if (getRow() < maxRow && getColumn() < maxColumn) {
      adjoining.add(new Coordinates(getRow() + 1, getColumn() + 1));
    }
    if (getRow() > 0 && getColumn() > 0) {
      adjoining.add(new Coordinates(getRow() - 1, getColumn() - 1));
    }
    if (getRow() > 0 && getColumn() < maxColumn) {
      adjoining.add(new Coordinates(getRow() - 1, getColumn() + 1));
    }
    return adjoining.stream().map(Coordinates::getAbbreviation).collect(Collectors.toSet());
  }


}
