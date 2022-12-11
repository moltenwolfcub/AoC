import java.util.HashSet;
import java.util.List;
import java.util.Set;

import Helpers.FileUtils;
import Helpers.Position2D;

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
}
