import Helpers.FileUtils;

public class day3_1 {
    
    private static int totalPriority;

    public static void main(String[] args) {
        totalPriority = 0;

        for (String rucksack : FileUtils.readLines("2022/day3.txt")) {
            String firstCompartment = rucksack.substring(0, rucksack.length()/2);
            String secondCompartment = rucksack.substring(rucksack.length()/2);

            Character sharedItem = ' ';
            for (Character supply : firstCompartment.toCharArray()) {
                if (secondCompartment.contains(supply.toString())) {
                    sharedItem = supply;
                    break;
                }
            }

            if (!Character.isLetter(sharedItem)) {
                System.out.println("No shared Item was found.");
                continue;
            }
            totalPriority+= getItemPriority(sharedItem);
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
