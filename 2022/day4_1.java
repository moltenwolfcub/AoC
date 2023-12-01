import java.util.function.Supplier;
import java.util.stream.IntStream;

import helpers.FileUtils;

public class day4_1 {
    private static Integer runningTotal;

    public static void main(String[] args) {
        runningTotal = 0;

        for (String line : FileUtils.readLines("2022/day4.txt")) {
            String[] elves = line.split(",");

            Supplier<IntStream> firstAssignedSegs = ()-> IntStream.range(Integer.valueOf(elves[0].split("-")[0]), Integer.valueOf(elves[0].split("-")[1])+1);
            Supplier<IntStream> secondAssignedSegs = ()-> IntStream.range(Integer.valueOf(elves[1].split("-")[0]), Integer.valueOf(elves[1].split("-")[1])+1);

            Long firstCount = firstAssignedSegs.get().count();
            Long secondCount = secondAssignedSegs.get().count();
            Boolean isFirstBigger = firstCount > secondCount;

            Supplier<IntStream> longer = isFirstBigger ? firstAssignedSegs : secondAssignedSegs;
            Supplier<IntStream> shorter = isFirstBigger ? secondAssignedSegs : firstAssignedSegs;

            if (shorter.get().allMatch(shortSeg-> longer.get().anyMatch(longSeg-> longSeg == shortSeg))) {
                runningTotal++;
            }
        }
        System.out.println(runningTotal);
    }
}
