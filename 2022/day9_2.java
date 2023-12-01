import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

import helpers.FileUtils;
import helpers.Position2D;

public class day9_2 {
	private static Position2D head;
	private static List<Position2D> tail;
	private static Set<List<Integer>> tailVisited = new HashSet<>();

    public static void main(String[] args) {
		head = new Position2D();
		tail = new ArrayList<>();
        for (int i = 0; i < 9; i++) {
            tail.add(new Position2D());
        }

		tailVisited.add(List.of(tail.get(tail.size()-1).x, tail.get(tail.size()-1).y));

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
    private static void move(Integer xDelta, Integer yDelta) {
		head.x += xDelta;
		head.y += yDelta;

        for (int i = 0; i<9; i++) {
            if (!moveTail(i)) {
                break;
            }
        }
    }


	private static Boolean moveTail(Integer idx) {
        Position2D currentPos = tail.get(idx);
        Position2D followingPos;
        if (idx == 0) {
            followingPos = head;
        } else {
            followingPos = tail.get(idx-1);
        }

		if (Math.abs(followingPos.x - currentPos.x) > 1 || Math.abs(followingPos.y - currentPos.y) > 1) {
			// the tail is too far and needs to move to catch up

			if (currentPos.y == followingPos.y) {
				// only needs to move along the x
				currentPos.x += (int)Math.signum(followingPos.x - currentPos.x);
			} else if (currentPos.x == followingPos.x) {
				// only needs to move along the y
				currentPos.y += (int)Math.signum(followingPos.y - currentPos.y);
			} else {
				currentPos.y += (int)Math.signum(followingPos.y - currentPos.y);
				currentPos.x += (int)Math.signum(followingPos.x - currentPos.x);
			}
            tailVisited.add(List.of(tail.get(tail.size()-1).x, tail.get(tail.size()-1).y));

            return true;
		}
        return false;
	}
    
}
