import java.util.ArrayList;
import java.util.List;
import java.util.stream.IntStream;

import Helpers.FileUtils;

public class day10_2 {
    private static Integer xRegister;
    private static List<List<Character>> output;

    public static void main(String[] args) {
        xRegister = 1;
        output = new ArrayList<>();

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

        List<Character> tmpSpriteRowData = new ArrayList<>();
        for (int clockCycle = 0; clockCycle < xRegChanges.size(); clockCycle++) {

            List<Integer> spriteCoverPos = IntStream.rangeClosed(xRegister-1, xRegister+1).boxed().toList();

            if (spriteCoverPos.contains(Math.floorMod(clockCycle, 40))) {
                tmpSpriteRowData.add('#');
            } else {
                tmpSpriteRowData.add('.');
            }
            
            if (Math.floorMod(clockCycle+1, 40) == 0) {
                output.add(tmpSpriteRowData);
                tmpSpriteRowData = new ArrayList<>();
            }
            xRegister+= xRegChanges.get(clockCycle);
        }

        for (List<Character> line : output) {
            for (Character character : line) {
                System.out.print(character);
            }
            System.out.println();
        }
    }
    
}
