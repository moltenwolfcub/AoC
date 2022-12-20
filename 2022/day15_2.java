import java.util.ArrayList;
import java.util.Iterator;
import java.util.List;
import java.util.regex.MatchResult;
import java.util.regex.Pattern;
import java.util.stream.Collectors;

import Helpers.FileUtils;
import Helpers.Position2D;

public class day15_2 {
    private static List<Sensor> sensors;
    private static Long tuningFrequency;
    
    public static void main(String[] args) {
        sensors = new ArrayList<>();
        tuningFrequency = 0l;
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

        Position2D beaconPos = findPossibleBeacon(4000000); 

        tuningFrequency = ((long)(int)beaconPos.x)*4000000l+((long)(int)beaconPos.y);
        System.out.println(tuningFrequency);
        System.out.println(beaconPos.x+", "+beaconPos.y);
    }

    private static Position2D findPossibleBeacon(Integer max) {
        Iterator<Sensor> sensorsIter = sensors.iterator();
        while (sensorsIter.hasNext()) {
            Sensor currentSensor = sensorsIter.next();
            for (Position2D testPos : currentSensor.getPerimeter()) {
                if (testPos.x < 0 || testPos.y < 0 || testPos.x > max || testPos.y > max) {
                    continue;
                }

                Boolean inRange = false;
                Iterator<Sensor> innerSensorsIter = sensors.iterator();
                while (innerSensorsIter.hasNext()) {
                    Sensor sensor = innerSensorsIter.next();
                    if (currentSensor.equals(sensor)) {
                        continue;
                    }
                    if (sensor.posInRange(testPos)) {
                        inRange = true;
                        break;
                    }
                }
                if (!inRange) {
                    System.out.println("Returning Test Pos: %s".formatted(testPos));
                    return testPos;
                }
            }
        }
        throw new RuntimeException("No beacon was found.");

    }

    private static class Sensor extends Position2D {
        private Position2D beaconPos;
        private Integer beaconDistance;

        public Sensor(Integer x, Integer y, Position2D closestBeacon) {
            super(x, y);
            beaconPos = closestBeacon;
            beaconDistance = this.distanceFromManhattan(closestBeacon);
        }
    
        public Boolean posInRange(Position2D pos) {
            Integer dist = this.distanceFromManhattan(pos);
            return dist <= beaconDistance;
        }

        public Integer getBeaconDistance() {
            return beaconDistance;
        }

        public List<Position2D> getPerimeter() {
            List<Position2D> perimeterPositions = new ArrayList<>();

            Position2D tmp = this.clone();
            tmp.y+= beaconDistance+1;
            while (tmp.y >= this.y) {
                perimeterPositions.add(tmp);
                tmp = new Position2D(tmp.x+1, tmp.y-1);
            }
            
            tmp = this.clone();
            tmp.x+= beaconDistance+1;
            while (tmp.x >= this.x) {
                perimeterPositions.add(tmp);
                tmp = new Position2D(tmp.x-1, tmp.y-1);
            }

            tmp = this.clone();
            tmp.y-= beaconDistance+1;
            while (tmp.y <= this.y) {
                perimeterPositions.add(tmp);
                tmp = new Position2D(tmp.x-1, tmp.y+1);
            }
            
            tmp = this.clone();
            tmp.x-= beaconDistance+1;
            while (tmp.x <= this.x) {
                perimeterPositions.add(tmp);
                tmp = new Position2D(tmp.x+1, tmp.y+1);
            }

            return perimeterPositions;
        }
        
        public Position2D getBeaconPos() {
            return beaconPos;
        }
        
    }
}
