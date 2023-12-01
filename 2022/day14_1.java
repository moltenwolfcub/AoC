import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;

import helpers.Axis;
import helpers.FileUtils;
import helpers.Position2D;

public class day14_1 {
    private static Map<Position2D, Boolean> solidityMap;
    private static Integer maxY;
    private static Integer settledSand;
    
    public static void main(String[] args) {
        solidityMap = new HashMap<>();
        settledSand = 0;
        maxY = 0;

        for (String line : FileUtils.readLines("2022/day14.txt")) {
            List<Character> inputChars = line.chars().mapToObj(i -> (char)i).collect(Collectors.toList());
            List<Position2D> points = parseLine(inputChars);
            fillSolidityList(points);
        }

        while (true) {
            Position2D sandPiece = new Position2D(500, 0);
            while (true) {
                if (sandPiece.y > maxY) {
                    break;
                } else if (!solidityMap.containsKey(new Position2D(sandPiece.x, sandPiece.y+1))) {
                    sandPiece.y += 1;
                } else if (!solidityMap.containsKey(new Position2D(sandPiece.x-1, sandPiece.y+1))) {
                    sandPiece.x -= 1;
                    sandPiece.y += 1;
                } else if (!solidityMap.containsKey(new Position2D(sandPiece.x+1, sandPiece.y+1))) {
                    sandPiece.x += 1;
                    sandPiece.y += 1;
                } else {
                    solidityMap.put(sandPiece, true);
                    settledSand++;
                    break;
                }
            }
            if (sandPiece.y > maxY) {
                break;
            }
        }
        System.out.println("%s sand units could settle before falling to infinity.".formatted(settledSand));
    }

    /**
     * Converts a {@code List} of {@code Character}s from the input file
     * to a List of {@code Position2D}s ordered by their connections.
     * 
     * @param inputLine             A list of characters to convert
     * @return                      An ordered list of positions by their
     *                              connection order
     */
    public static List<Position2D> parseLine(List<Character> inputLine) {
        List<Position2D> returnPositions = new ArrayList<>();

        String currentInteger = "";
        Integer currentX = null;
        Integer currentY = null;

        while (inputLine.size() > 0) {
            Character currentChar = inputLine.get(0);
            inputLine.remove(0);

            if (Character.isDigit(currentChar)) {
                currentInteger = currentInteger.concat(currentChar.toString());
                continue;
            }

            switch (currentChar) {
                case ',':
                    currentX = Integer.valueOf(currentInteger);
                    currentInteger = "";
                    continue;
                case '>':
                    currentY = Integer.valueOf(currentInteger);
                    if (currentX==null || currentY==null) {
                        throw new IllegalArgumentException("Didn't read in X and Y values before trying to write to a Position2D");
                    }
                    maxY = Math.max(maxY, currentY);
                    Position2D currentPosition = new Position2D(currentX, currentY);
                    returnPositions.add(currentPosition);

                    currentInteger = "";
                    currentX = null;
                    currentY = null;
                    currentPosition = null;
                    continue;
            }
        }
        currentY = Integer.valueOf(currentInteger);
        Position2D currentPosition = new Position2D(currentX, currentY);
        returnPositions.add(currentPosition);

        return returnPositions;
    }

    private static void fillSolidityList(List<Position2D> points) {
        for (int i = 0; i < points.size()-1; i++) {
            Position2D start = points.get(i);
            Position2D end = points.get(i+1);
            Axis differentAxis = null;
            if (start.x.equals(end.x)) {
                differentAxis = Axis.Y;
            } else if (start.y.equals(end.y)) {
                differentAxis = Axis.X;
            } else {
                throw new IllegalArgumentException("The two points shared no axis. There is no current support for diagonal lines.");
            }

            Integer difference = differentAxis.getPositionAxisValue(end)-differentAxis.getPositionAxisValue(start);
            Integer dir = (int)Math.signum(difference);

            Position2D midPos = null;
            for (int j = 0; j < Math.abs(difference); j++) {
                midPos = differentAxis.movePosition(start, dir*j);
                solidityMap.put(midPos, true);
            }
        }
        solidityMap.put(points.get(points.size()-1), true);
    }

}
