import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;

import Helpers.FileUtils;

public class Day1_2 {
    private static List<String> elfCalorieData;
    private static List<Integer> topCalorieCounts;
    private static Integer topCombinedCaloriesSum;

    public static void main(String[] args) {
        elfCalorieData = FileUtils.readLines("2022/day1.txt");
        topCalorieCounts = new ArrayList<Integer>(Arrays.asList(0, 0, 0));
        topCombinedCaloriesSum = 0;

        Integer currentTotal = 0;
        for (String calorie : elfCalorieData) {
            if (calorie.isBlank()) {

                //add new item
                topCalorieCounts.add(currentTotal);

                //put in descending order
                Collections.sort(topCalorieCounts);
                Collections.reverse(topCalorieCounts);

                //make a new list of the top 3
                List<Integer> newList = new ArrayList<>();
                for (int i = 0; i < 3; i++) {
                    newList.add(topCalorieCounts.get(i));
                }
                //replace the original list with the top 3 list
                topCalorieCounts.clear();
                topCalorieCounts.addAll(newList);

                currentTotal = 0;

            } else {
                currentTotal += Integer.valueOf(calorie);
            }
        }
        
        //sum the calories
        for (Integer calorie : topCalorieCounts) {
            topCombinedCaloriesSum += calorie;
        }
        System.out.println(topCombinedCaloriesSum);
    }
}
