import java.util.HashSet;
import java.util.List;
import java.util.Set;

import Helpers.FileUtils;

public class day9_1 {
	private static Position2D head;
	private static Position2D tail;
	private static Set<List<Integer>> tailVisited = new HashSet<>();

    public static void main(String[] args) {
		head = new Position2D();
		tail = new Position2D();
		tailVisited.add(List.of(tail.x, tail.y));

        for (String input : FileUtils.readLines("2022/day9.txt")) {
            String[] commands = input.split(" ");
			Integer moveDistance = Integer.valueOf(commands[1]);
			moveBy(moveDistance, commands[0]);
        }

		System.out.println(tailVisited.size());
    }

	private static void moveBy(Integer distance, String direction) throws IllegalArgumentException {
		switch (direction) {
			case "U":
				for (int i = 0; i < distance; i++) {
					move(0, 1); 
				}
				break;
			case "D":
				for (int i = 0; i < distance; i++) {
					move(0, -1); 
				}
				break;
			case "L":
				for (int i = 0; i < distance; i++) {
					move(-1, 0); 
				}
				break;
			case "R":
				for (int i = 0; i < distance; i++) {
					move(1, 0); 
				}
				break;
			default: 
				throw new IllegalArgumentException("Parsed value in file wasn't one of U, D, L or R.");
		}
	}
	private static void move(Integer x, Integer y) {
		head.x += x;
		head.y += y;

		if (Math.abs(head.x - tail.x) > 1 || Math.abs(head.y - tail.y) > 1) {
			// the tail is too far and needs to move to catch up

			if (tail.y == head.y) {
				// only needs to move along the x
				tail.x += (int)Math.signum(head.x - tail.x);
			} else if (tail.x == head.x) {
				// only needs to move along the y
				tail.y += (int)Math.signum(head.y - tail.y);
			} else {
				tail.y += (int)Math.signum(head.y - tail.y);
				tail.x += (int)Math.signum(head.x - tail.x);
			}
			tailVisited.add(List.of(tail.x, tail.y));
		}
	}

	public static class Position2D implements Cloneable {
		Integer x;
		Integer y;

		public Position2D() {
			this.x = 0;
			this.y = 0;
		}

		public Position2D(Integer x, Integer y) {
			this.x = x;
			this.y = y;
		}

		@Override
		public Position2D clone() {
			return new Position2D(x, y);
		}

		@Override
		public String toString() {
			return "<x: "+ this.x+", y: "+this.y+">";
		}

		@Override
		public int hashCode() {
			final int prime = 31;
			int hash = 7;
			hash = prime * hash + ((x == null) ? 0 : x.hashCode());
			hash = prime * hash + ((y == null) ? 0 : y.hashCode());
			return hash;
		}

		@Override
		public boolean equals(Object other) {
			if (this == other) {
				return true;
			}
			if (other == null || this == null) {
				return false;
			}
			if (other instanceof Position2D) {
				Position2D pos = (Position2D) other;
				return pos.x == this.x && pos.y == this.y;
			}
			return false;
		}

	}
}
