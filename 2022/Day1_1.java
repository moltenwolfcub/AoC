import java.io.File;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

public class Day1_1 {
    private static Integer mostCalories = 0;
    private static List<List<String>> elfSeperatedCalories = new ArrayList<>();
    private static List<Integer> totalElfCalories = new ArrayList<>();

    //Retrieve the file
    private static List<String> allElfCalories = readFile("2022/day1_1.txt");

    public static void main(String[] args) {
        //seperate out the individual elves by the space delimiter.
        List<String> tmpList = new ArrayList<>();
        for (String calorieData : allElfCalories) {
            if (calorieData.isEmpty()) {
                elfSeperatedCalories.add(tmpList);
                tmpList = new ArrayList<>();
            } else {
                tmpList.add(calorieData);
            }
        }
        elfSeperatedCalories.add(tmpList);

        //sum each elf's individual food items.
        for (List<String> individualElfCalories : elfSeperatedCalories) {
            Integer calorieSum = 0;
            for (String calorie : individualElfCalories) {
                calorieSum += Integer.valueOf(calorie);
            }
            totalElfCalories.add(calorieSum);
        }

        //find the biggest item in the list.
        totalElfCalories.forEach((elfCalorieCount) -> {
            mostCalories = Math.max(mostCalories, elfCalorieCount);
        });

        System.out.println(mostCalories);
    }

    private static List<String> readFile(String fileName) {
        List<String> data = new ArrayList<>();

        try {
            File file = new File(fileName);
            Scanner fileReader = new Scanner(file);
            while (fileReader.hasNextLine()) {
                data.add(fileReader.nextLine());
            }
            fileReader.close();
        } catch (IOException e) {
            System.out.println("An Error Occured while reading in the file.");
            e.printStackTrace();
            System.out.println("The data being returned will be incomplete or incorrect.");
        }
        return data;
    }
}
