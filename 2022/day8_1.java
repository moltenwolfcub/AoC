import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

import Helpers.FileUtils;

public class day8_1 {
	private static Integer runningTotal;
    
    public static void main(String[] args) {
		runningTotal = 0;

        List<String> treeRows = FileUtils.readLines("2022/day8.txt");
        List<List<Integer>> trees = new ArrayList<>();

        for (String row : treeRows) {
            List<Integer> tmpTreeRow = new ArrayList<>();
            row.chars().forEach(tree -> tmpTreeRow.add(Integer.valueOf(((Character)(char)tree).toString())));
            trees.add(tmpTreeRow);
        }

        for (int yPos = 0; yPos < trees.size(); yPos++) {
            List<Integer> treeRow = trees.get(yPos);
            for (int xPos = 0; xPos < treeRow.size(); xPos++) {
				if (checkTreeVisiblity(trees, xPos, yPos)) {
					runningTotal++;
				}
            }
        }
		System.out.println(runningTotal);
    }

    private static Boolean checkTreeVisiblity(List<List<Integer>> treeTable, Integer x, Integer y) {
        Integer thisTreeHeight = treeTable.get(x).get(y);

        List<Integer> row = new ArrayList<>();
		for (Integer tree : treeTable.get(x)) {
			row.add(tree);
		}
		if (checkTreeRow(y, row) < thisTreeHeight) {
			return true;
		}
		List<Integer> reversed = new ArrayList<>(row);
		Collections.reverse(reversed);
		if (checkTreeRow(row.size()-y-1, reversed) < thisTreeHeight) {
			return true;
		}

		row = new ArrayList<>();
		for (List<Integer> treeRow : treeTable) {
			row.add(treeRow.get(y));
		}
		if (checkTreeRow(x, row) < thisTreeHeight) {
			return true;
		}
		reversed = new ArrayList<>(row);
		Collections.reverse(reversed);
		if (checkTreeRow(row.size()-x-1, reversed) < thisTreeHeight) {
			return true;
		}

        return false;
    }

	private static Integer checkTreeRow(Integer treeIdx, List<Integer> row) {
		Integer currentHeight = -1;
		for (int tree = 0; tree < treeIdx; tree++) {
			if (row.get(tree) > currentHeight) {
				currentHeight = row.get(tree);
			}
        }
		return currentHeight;
	}
}
