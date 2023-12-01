import java.util.ArrayList;
import java.util.List;
import java.util.regex.MatchResult;
import java.util.regex.Pattern;
import java.util.stream.Collectors;

import helpers.FileUtils;
import helpers.Position2D;

public class day15_1 {
    private static List<Sensor> sensors;
    private static Integer noBeaconPositions;
    
    public static void main(String[] args) {
        sensors = new ArrayList<>();
        noBeaconPositions = 0;
        Integer minX = Integer.MAX_VALUE;
        Integer maxX = 0;
        Integer maxBeaconDist = 0;
        
        for (String line : FileUtils.readLines("2022/day15.txt")) {
            List<MatchResult> lineData = Pattern.compile("x=(-?\\d+), y=(-?\\d+)").matcher(line).results().collect(Collectors.toList());

            Integer beaconX = Integer.valueOf(lineData.get(1).group(1));
            Integer beaconY = Integer.valueOf(lineData.get(1).group(2));
            Position2D beacon = new Position2D(beaconX, beaconY);
            Integer sensorX = Integer.valueOf(lineData.get(0).group(1));
            Integer sensorY = Integer.valueOf(lineData.get(0).group(2));
            Sensor sensor = new Sensor(sensorX, sensorY, beacon);

            minX = Math.min(minX, sensorX);
            maxX = Math.max(maxX, sensorX);
            maxBeaconDist = Math.max(maxBeaconDist, sensor.getBeaconDistance());

            sensors.add(sensor);
        }

        Integer rangeBottom = minX-maxBeaconDist;
        Integer rangeTop = maxX+maxBeaconDist;
        for (int i = rangeBottom; i < rangeTop-rangeBottom; i++) {
            Position2D testPos = new Position2D(i, 2000000);
            Boolean posChecked = false;
            for (Sensor sensor : sensors) {
                if (sensor.posIsMyBeacon(testPos)) {
                    posChecked = false;
                    break;
                }
                posChecked = posChecked || sensor.checkedForBeacon(testPos);
            }
            if (posChecked) {
                noBeaconPositions++;
            }
        }
        System.out.println(noBeaconPositions);
    }

    private static class Sensor extends Position2D {
        private Position2D beaconPos;
        private Integer beaconDistance;

        public Sensor(Integer x, Integer y, Position2D closestBeacon) {
            super(x, y);
            beaconPos = closestBeacon;
            beaconDistance = this.distanceFromManhattan(closestBeacon);
        }
    
        public Boolean checkedForBeacon(Position2D pos) {
            Integer dist = this.distanceFromManhattan(pos);
            return dist <= beaconDistance;
        }

        public Boolean posIsMyBeacon(Position2D pos) {
            return beaconPos.equals(pos);
        }

        public Integer getBeaconDistance() {
            return beaconDistance;
        }
        
    }
}
