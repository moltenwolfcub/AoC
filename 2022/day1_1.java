import java.util.List;

import Helpers.FileUtils;

public class day1_1 {
    private static Integer mostCalories = 0;
    private static List<String> elfCalorieData;

    public static void main(String[] args) {
        elfCalorieData = FileUtils.readLines("2022/day1.txt");

        Integer currentTotal = 0;
        for (String calorie : elfCalorieData) {
            if (calorie.isBlank()) {
                mostCalories = Math.max(currentTotal, mostCalories);
                currentTotal = 0;
            } else {
                currentTotal += Integer.valueOf(calorie);
            }
        }
        System.out.println(mostCalories);
    }
}
