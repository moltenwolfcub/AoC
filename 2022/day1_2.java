import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;

import helpers.FileUtils;

public class day1_2 {
    private static List<String> elfCalorieData;
    private static List<Integer> allCalorieSums;
    private static Integer topCombinedCaloriesSum;

    public static void main(String[] args) {
        elfCalorieData = FileUtils.readLines("2022/day1.txt");
        allCalorieSums = new ArrayList<Integer>(Arrays.asList(0, 0, 0));
        topCombinedCaloriesSum = 0;

        Integer currentTotal = 0;
        for (String calorie : elfCalorieData) {
            if (calorie.isBlank()) {
                allCalorieSums.add(currentTotal);
                currentTotal = 0;

            } else {
                currentTotal += Integer.valueOf(calorie);
            }
        }

        Collections.sort(allCalorieSums);
        Collections.reverse(allCalorieSums);
        for (int i = 0; i < 3; i++) {
            topCombinedCaloriesSum += allCalorieSums.get(i);
        }
        System.out.println(topCombinedCaloriesSum);
    }
}
