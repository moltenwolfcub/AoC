import java.util.ArrayList;
import java.util.List;

import helpers.FileUtils;

public class day10_1 {
    private static Integer xRegister;
    private static List<Integer> significantSignalStrengths;
    private static Integer output;

    public static void main(String[] args) {
        xRegister = 1;
        significantSignalStrengths = new ArrayList<>();
        output = 0;

        List<Integer> xRegChanges = new ArrayList<>();

        for (String fullCommandString : FileUtils.readLines("2022/day10.txt")) {
            String[] fullCommand = fullCommandString.split(" ");
            switch (fullCommand[0]) {
                case "noop":
                    xRegChanges.add(0);
                    break;
                case "addx":
                    xRegChanges.add(0);
                    xRegChanges.add(Integer.valueOf(fullCommand[1]));
                    break;
                default: 
                    break;
            }
        }

        for (int clockCycle = 0; clockCycle < xRegChanges.size(); clockCycle++) {
            if (Math.floorMod(clockCycle+1-20, 40) == 0) {
                significantSignalStrengths.add(xRegister*(clockCycle+1));
            }

            System.out.println("During the "+(clockCycle+1)+"th clock cycle the Xregister has the value: "+ xRegister);
            xRegister+= xRegChanges.get(clockCycle);
        }
        System.out.println("After all commands have ran the Xregister value is: "+xRegister);

        for (Integer signalStrenth : significantSignalStrengths) {
            output+=signalStrenth;
        }

        System.out.println(output);
    }
}
