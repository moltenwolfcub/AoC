import java.util.ArrayList;
import java.util.Collection;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.regex.MatchResult;
import java.util.regex.Pattern;
import java.util.stream.Collectors;

import Helpers.FileUtils;

public class day16_1 {
    private static Map<String, Valve> valves;
    private static Integer mostPressure;
    
    public static void main(String[] args) {
        System.out.println();
        valves = new HashMap<>();
        mostPressure = 0;

        parseInput();

        Valve startValve = valves.get("AA");
        if (startValve == null) {
            throw new NullPointerException("Couldn't find 'AA' in the set of valves.");
        }
        NodePoint startNode = new NodePoint(startValve, 30, 0, valves.values());

        List<NodePoint> current = new ArrayList<>();
        List<NodePoint> next = new ArrayList<>();
        current.add(startNode);

        while (!current.isEmpty()) {
            for (NodePoint node : current) {
                for (NodePoint neighbour : node.getPossibleMoves(valves.values())) {
                    mostPressure = Math.max(neighbour.currentPressureReleased, mostPressure);
                    next.add(neighbour);
                }
            }
            current.clear();
            current.addAll(next);
            next = new ArrayList<>();
        }
        System.out.println(mostPressure);
    }

    public static void parseInput() {
        for (String line : FileUtils.readLines("2022/day16.txt")) {
            MatchResult lineData = Pattern.compile("Valve ([A-Z]+) has flow rate=(\\d+); tunnels? leads? to valves? ([A-Z]+(?:, [A-Z]+)*)")
                .matcher(line)
                .results()
                .collect(Collectors.toList())
                .get(0);
            
            String valveName = lineData.group(1);
            Integer flowRate = Integer.valueOf(lineData.group(2));
            List<String> neighbours = FileUtils.splitCSV(lineData.group(3));
            Valve readValve = new Valve(valveName, flowRate, neighbours);
            valves.put(valveName, readValve);
        }
        valves.values().forEach((Valve v) -> v.populateNeighbours(valves));

    }


    private static class NodePoint {
        private Integer timeLeft;
        private Valve currentValve;
        private List<Valve> hasPressure;
        private Integer currentPressureReleased;

        public NodePoint(Valve valve, Integer time, Integer pressure, Collection<Valve> closedValves) {
            this.timeLeft = time;
            this.currentValve = valve;
            this.hasPressure = closedValves.stream().filter(Valve::hasPressure).collect(Collectors.toList());
            this.currentPressureReleased = pressure;
        }

        public List<NodePoint> getPossibleMoves(Collection<Valve> allValves) {
            List<NodePoint> possibleMoves = new ArrayList<>();

            allValves = allValves.stream().filter(Valve::hasPressure).collect(Collectors.toList());

            for (Valve valve : allValves) {
                if (!hasPressure.contains(valve)) {
                    continue;
                }
                NodePoint newPoint = createNewPoint(valve);
                if(newPoint.timeLeft <=0) {
                    continue;
                }
                possibleMoves.add(newPoint);
            }

            return possibleMoves;
        }

        public NodePoint createNewPoint(Valve newValve) {
            Integer timeToMove = currentValve.timeFrom(newValve);
            NodePoint newPoint = new NodePoint(
                newValve,
                this.timeLeft-timeToMove-1,
                this.currentPressureReleased,
                hasPressure
            );
            newPoint.hasPressure.remove(newValve);
            if (newPoint.timeLeft > 0) {
                newPoint.currentPressureReleased += newValve.openValve(newPoint.timeLeft);
            }
            return newPoint;
        }


        @Override
        public String toString() {
            return "\n<%s>\nWith %s minutes left\n%s pressure can be released so far.\n%s can still be opened.\n"
                .formatted(currentValve.toString(true), timeLeft, currentPressureReleased, hasPressure);
        }
    }

    private static class Valve {
        private String label;
        private Integer flowRate;
        private List<String> neighbourIds;
        private List<Valve> neighbours;
        private Map<Valve, Integer> distanceCache;

        public Valve(String label, Integer flowRate, List<String> readNeighbours) {
            this.label = label;
            this.flowRate = flowRate;
            this.neighbourIds = readNeighbours;
            this.neighbours = new ArrayList<>();
        }

        public void populateNeighbours(Map<String, Valve> allValves) {
            for (String neighbour : neighbourIds) {
                Valve neighbouringValve = allValves.get(neighbour);
                if (neighbouringValve == null) {
                    throw new NullPointerException("Couldn't locate valve %s in the provided map.".formatted(neighbour));
                }
                neighbours.add(neighbouringValve);
            }
        }


        public Integer timeFrom(Valve otherValve) {
            if(this.distanceCache == null) {
                this.distanceCache = new HashMap<>();
            }
            Integer dist;
            if((dist = this.distanceCache.get(otherValve)) != null) {
                return dist;
            }

            List<Valve> current = new ArrayList<>();
            List<Valve> next = new ArrayList<>();
            List<Valve> visited = new ArrayList<>();
            current.add(this);
            visited.add(this);
            distanceCache.put(this, 0);
    
            while (!current.isEmpty()) {
                for (Valve valve : current) {
                    for (Valve neighbour : valve.neighbours) {
                        if(visited.contains(neighbour)) {
                            continue;
                        }
                        visited.add(neighbour);
                        if (!this.distanceCache.containsKey(neighbour)) {
                            this.distanceCache.put(neighbour, this.distanceCache.get(valve)+1);
                        }
                        if (otherValve.equals(neighbour)) {
                            return distanceCache.get(neighbour);
                        }
                        next.add(neighbour);
                    }
                }
                current.clear();
                current.addAll(next);
                next = new ArrayList<>();
            }
            
            System.out.println("ERROR: No path found from %s to %s".formatted(this, otherValve));
            return 0;
        }

        public Integer openValve(Integer timeLeft) {
            return flowRate*timeLeft;
        }


        public String getLabel() {
            return label;
        }
        public static boolean hasPressure(Valve v) {
            return v.flowRate > 0;
        }

        @Override
        public String toString() {
            return toString(false);
        }
        public String toString(Boolean debug) {
            if (debug) {
                return "%s valve with flow: %s. Neighbours: %s".formatted(label, flowRate, neighbourIds);
            } else {
                return getLabel();
            }
        }
    }
}
