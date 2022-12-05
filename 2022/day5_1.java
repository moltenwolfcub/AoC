import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

import Helpers.FileUtils;

public class day5_1 {
    private static List<String> file;

    private static List<String> startingCrates;
    private static List<String> rearangmentProcesses;

    private static List<List<String>> crateStacks;

    private static String output;
    
    public static void main(String[] args) {
        readFileData();
        loadStartingStacks();

        for (String rearangement : rearangmentProcesses) {
            String[] arguments = rearangement.split(" ");
            for (int i = 0; i < Integer.valueOf(arguments[1]); i++) {
                List<String> startStack = crateStacks.get(Integer.valueOf(arguments[3])-1);
                List<String> destStack = crateStacks.get(Integer.valueOf(arguments[5])-1);

                destStack.add(startStack.get(startStack.size()-1));
                startStack.remove(startStack.size()-1);
            }
        }
        
        output = "";
        for (List<String> stack : crateStacks) {
            output+= stack.get(stack.size()-1);
        }
        System.out.println(output);
    }

    private static void readFileData() {
        file = FileUtils.readLines("2022/day5.txt");

        startingCrates = new ArrayList<>();
        rearangmentProcesses = new ArrayList<>();

        List<String> readList = new ArrayList<>();
        for (String line : file) {
            if(line.isBlank()) {
                startingCrates.addAll(readList);
                readList.clear();
                continue;
            }
            readList.add(line);
        }
        rearangmentProcesses.addAll(readList);
        readList.clear();
    }

    private static void loadStartingStacks() {
        Collections.reverse(startingCrates);
        String crateStackIds = startingCrates.get(0).replaceAll("\\s", "");
        Integer crateStackCount = crateStackIds.length();

        crateStacks = new ArrayList<>();
        for (int i = 0; i < crateStackCount; i++) {
            crateStacks.add(new ArrayList<>());
        }
        startingCrates.remove(0);
        for (String crateRow : startingCrates) {
            for (int i = 0; i < crateStackCount; i++) {
                String item = String.valueOf(crateRow.charAt(i*4+1));
                if (!item.isBlank()){
                    crateStacks.get(i).add(item);
                }
            }
        }
        
    }

}
