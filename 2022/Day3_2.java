import java.util.ArrayList;
import java.util.List;

import Helpers.FileUtils;

public class Day3_2 {

    private static int totalPriority;
    
    public static void main(String[] args) {
        totalPriority = 0;

        int rucksackId = 0;
        List<String> currentGroup = new ArrayList<String>();

        for (String rucksack : FileUtils.readLines("2022/day3.txt")) {
            currentGroup.add(rucksack);

            Character badgeItem = ' ';
            if(Math.floorMod(rucksackId, 3) == 2) {
                for (Character supply : currentGroup.get(0).toCharArray()) {
                    if (currentGroup.get(1).contains(supply.toString()) && currentGroup.get(2).contains(supply.toString())) {
                        badgeItem = supply;
                        break;
                    }
                }
                if (!Character.isLetter(badgeItem)) {
                    System.out.println("Badge Item was found.");
                    continue;
                }
                totalPriority+= getItemPriority(badgeItem);
                currentGroup.clear();
            }

            rucksackId += 1;
        }
        System.out.println(totalPriority);
    }

    private static Integer getItemPriority(Character item) {
        int priority = Character.getNumericValue(item) - 9;
        if (Character.isUpperCase(item)) {
            priority += 26;
        }
        return priority;
    }
}
