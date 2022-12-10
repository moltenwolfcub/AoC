import java.util.ArrayList;
import java.util.List;

import Helpers.FileUtils;

public class day8_2 {
	private static Integer highestScenic;
    
    public static void main(String[] args) {
        highestScenic = 0;

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
                highestScenic = Math.max(highestScenic, calculateScenic(trees, xPos, yPos));
            }
        }
        System.out.println(highestScenic);
    }

    private static Integer calculateScenic(List<List<Integer>> treeTable, Integer x, Integer y) {
        Integer thisTreeHeight = treeTable.get(y).get(x);

        List<Integer> xRow = new ArrayList<>();
		for (Integer tree : treeTable.get(y)) {
			xRow.add(tree);
		}
		List<Integer> yRow = new ArrayList<>();
		for (List<Integer> treeRow : treeTable) {
			yRow.add(treeRow.get(x));
		}

        Integer up = calculateRowScenic(thisTreeHeight, y, yRow, false);
        Integer down = calculateRowScenic(thisTreeHeight, y, yRow, true);
        Integer left = calculateRowScenic(thisTreeHeight, x, xRow, false);
		Integer right = calculateRowScenic(thisTreeHeight, x, xRow, true);

        return right*up*left*down;
    }

	private static Integer calculateRowScenic(Integer thisHeight, Integer treeIdx, List<Integer> row, Boolean forwards) {
		Integer visible = 0;
        if (forwards) {
            for (int tree = treeIdx+1; tree <= row.size()-1; tree++) {
                visible++;
                if (row.get(tree) >= thisHeight) {
                    break;
                }
            }
        } else {
            for (int tree = treeIdx-1; tree >= 0; tree--) {
                visible++;
                if (row.get(tree) >= thisHeight) {
                    break;
                }
            }
        }
		return visible;
	}
}
