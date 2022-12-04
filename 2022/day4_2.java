import java.util.Set;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

import Helpers.FileUtils;

public class day4_2 {
    private static Integer runningTotal;
    
    public static void main(String[] args) {
        runningTotal = 0;

        for (String line : FileUtils.readLines("2022/day4.txt")) {
            String[] elves = line.split(",");

            IntStream firstAssignedSegs = IntStream.range(Integer.valueOf(elves[0].split("-")[0]), Integer.valueOf(elves[0].split("-")[1])+1);
            IntStream secondAssignedSegs = IntStream.range(Integer.valueOf(elves[1].split("-")[0]), Integer.valueOf(elves[1].split("-")[1])+1);

            Set<Integer> firstSegs = firstAssignedSegs.boxed().collect(Collectors.toSet());
            Set<Integer> intersection = secondAssignedSegs.boxed().filter(firstSegs::contains).collect(Collectors.toSet());
            if (!intersection.isEmpty()) {
                runningTotal++;
            }
        }
        System.out.println(runningTotal);
    }
}
