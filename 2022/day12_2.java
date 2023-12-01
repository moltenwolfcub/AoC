import java.util.ArrayList;
import java.util.Iterator;
import java.util.List;
import java.util.NoSuchElementException;

import helpers.FileUtils;
import helpers.Position2D;

public class day12_2 {
    private static HeightGrid heightmap;
    private static Position2D endPos;
    private static Integer shortestDistance;
    
    public static void main(String[] args) {
        heightmap = new HeightGrid();

        String fileName = "2022/day12.txt";

        Integer rowLength = FileUtils.readLines(fileName).get(0).length();


        Integer currentX = 0;
        Integer currentY = 0;
        List<HeightTile> currentList = new ArrayList<>();
        for (char readValue : FileUtils.readChars(fileName)) {
            Integer height = readValue - '`';
            if (height<1) {
                switch (readValue) {
                    case 'S':
                        height = 1;
                        break;
                    case 'E':
                        height = 26;
                        endPos = new Position2D(currentX, currentY);
                        break;
                    default:
                        throw new IllegalArgumentException("Character in heightmap was before the lowercase letters and wasn't a start or end pos.");
                }
            }
            currentList.add(new HeightTile(currentX, currentY, height));

            currentX++;
            if (currentX>=rowLength) {
                currentX = 0;
                currentY++;
                heightmap.addRow(currentList);
                currentList = new ArrayList<>();
            }
        }

        shortestDistance = Integer.MAX_VALUE;
        for (HeightTile tile : heightmap) {
            if(tile.height!=1) {
                continue;
            }
            shortestDistance = Math.min(shortestDistance, shortestPath(tile));
            heightmap.resetVisited();
        }
        System.out.println(shortestDistance);
    }

    private static Integer shortestPath(HeightTile origin) {
        List<HeightTile> current = new ArrayList<>();
        List<HeightTile> next = new ArrayList<>();
        current.add(origin);
        origin.setDistance(0);

        while (!current.isEmpty()) {
            for (HeightTile focusedTile : current) {
                for (HeightTile neighbour : focusedTile.getPossibleMoves(heightmap)) {
                    if(neighbour.visited()) {
                        continue;
                    }
                    neighbour.setDistance(focusedTile.getDistanceFromStart()+1);
                    if (endPos.equals(neighbour)) {
                        return neighbour.getDistanceFromStart();
                    }
                    next.add(neighbour);
                }
            }
            current.clear();
            current.addAll(next);
            next = new ArrayList<>();
        }
        return Integer.MAX_VALUE;
    }

    private static void debugHeightMap() {
        System.out.println("EndPos: "+endPos);
        for (HeightTile tile : heightmap) {
            System.out.println(tile);
        }
    }

    private static class HeightTile extends Position2D {
        protected Integer height;
        protected Integer distanceFromStart;

        public HeightTile(Integer x, Integer y, Integer height) {
            super(x, y);
            this.height = height;
            this.distanceFromStart = -1;
        }

        public void setDistance(Integer dist) {
            this.distanceFromStart = dist;
        }

        public Integer getHeight() {
            return height;
        }

        public Integer getDistanceFromStart() {
            return distanceFromStart;
        }

        public boolean visited() {
            return getDistanceFromStart()>-1;
        }

        public void resetVisited() {
            distanceFromStart = -1;
        }

        public List<HeightTile> getPossibleMoves(HeightGrid grid) {
            List<HeightTile> possibleMoves = new ArrayList<>();

            List<HeightTile> adjacentTiles = new ArrayList<>();
            adjacentTiles.add(grid.getTile(x, y+1));
            adjacentTiles.add(grid.getTile(x, y-1));
            adjacentTiles.add(grid.getTile(x+1, y));
            adjacentTiles.add(grid.getTile(x-1, y));

            for (HeightTile tile : adjacentTiles) {
                if (tile == null) {
                    continue;
                }
                if (tile.getHeight() <= this.getHeight()+1) {
                    possibleMoves.add(tile);
                }
            };

            return possibleMoves;
        }

        @Override
        public String toString() {
            return "Height Tile at %s, %s with height %s.".formatted(x, y, height);
        }
    }

    private static class HeightGrid implements Iterable<HeightTile> {
        private List<List<HeightTile>> grid;
        
        public HeightGrid() {
            grid = new ArrayList<>();
        }

        public void addRow(List<HeightTile> row) {
            grid.add(row);
        }

        public HeightTile getTile(Integer x, Integer y) {
            try {
                List<HeightTile> row = grid.get(y);
                return row.get(x);
            } catch (IndexOutOfBoundsException e) {
                return null;
            }
        }

        public void resetVisited() {
            forEach(t -> t.resetVisited());
        }


        @Override
        public Iterator<HeightTile> iterator() {
            return new Itr();
        }

        private class Itr implements Iterator<HeightTile> {
            Integer cursor = 0;
            Integer size = 0;
            List<HeightTile> list = new ArrayList<>();

            private Itr() {
                for (List<HeightTile> row : grid) {
                    for(HeightTile tile : row) {
                        list.add(tile);
                    }
                }
                size = list.size();
            }

            @Override
            public boolean hasNext() {
                return cursor < size;
            }

            @Override
            public HeightTile next() {
                try {
                    HeightTile next = list.get(cursor);
                    cursor++;
                    return next;
                } catch (IndexOutOfBoundsException e) {
                    throw new NoSuchElementException(e);
                }
            }
            
        }
    }
}
