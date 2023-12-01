import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.function.Function;

import helpers.FileUtils;

public class day11_1 {
    private static List<Monkey> monkeys;
    private static List<Long> monkeyActivity;
    private static Long monkeyBusiness;
    
    public static void main(String[] args) {
        monkeys = new ArrayList<>();
        monkeyActivity = new ArrayList<>();
        monkeyBusiness = 0l;

        Monkey.Builder current = new Monkey.Builder();
        for (String line : FileUtils.readLines("2022/day11.txt")) {
            if(line.isBlank()) {
                monkeys.add(current.build());
                current = null;
                continue;
            }
            String[] commands = line.strip().replace(":", "").split(" ");
            current = handleInput(commands, current);
        }
        if (current != null) {
            monkeys.add(current.build());
        }
        for (int i = 0; i < 20; i++) {
            runARound(false);
        }
        for (Monkey monkey : monkeys) {
            monkeyActivity.add(monkey.inspectCount);
        }
        Collections.sort(monkeyActivity);
        Collections.reverse(monkeyActivity);

        monkeyBusiness = monkeyActivity.get(0) * monkeyActivity.get(1);
        System.out.println(monkeyBusiness);
        System.out.println(monkeyActivity);
    }

    private static void runARound(boolean debug) {
        Long tmpCount = 0l;

        for (Monkey monkey : monkeys) {
            for (Long item : monkey.currentItems) {
                monkey.inspect();
                item = monkey.modifyWorry(item);
                item = Math.floorDiv(item, 3);
                Long newMonkeyId = Math.floorMod(item, monkey.testDivisor)==0 ? monkey.testTrueThrowId : monkey.testFalseThrowId;
                monkeys.get((int)(long)newMonkeyId).addItem(item);
            }
            monkey.currentItems.clear();
            if (debug) {
                System.out.println(tmpCount+" Monkey step");
                monkeys.forEach(System.out::println);
            }
            tmpCount++;
        }
    }

    private static Monkey.Builder handleInput(String[] commands, Monkey.Builder currentBuilder) {
        switch (commands[0].toLowerCase()) {
            case "monkey": 
                currentBuilder = new Monkey.Builder();
                break;
            case "starting":
                for (String command : commands) {
                    String value = command.split(",")[0];
                    Long numberVal;
                    try {
                        numberVal = Long.valueOf(value);
                    } catch (NumberFormatException e) {
                        continue;
                    }
                    currentBuilder.addItem(numberVal);
                }
                break;
            case "operation":
                try {
                    Long value = Long.valueOf(commands[5]);
                    switch (commands[4]) {
                        case "+": currentBuilder.setOperation((x) -> x + value); break;
                        case "*": currentBuilder.setOperation((x) -> x * value); break;
                        default: throw new IllegalArgumentException("Operation wasn't * or +");
                    }
                } catch (NumberFormatException e) {
                    switch (commands[4]) {
                        case "+": currentBuilder.setOperation((x) -> x+x); break;
                        case "*": currentBuilder.setOperation((x) -> x*x); break;
                        default: throw new IllegalArgumentException("Operation wasn't * or +");
                    }
                }
                break;
            case "test":
                Long divisor = Long.valueOf(commands[3]);
                currentBuilder.setTest(divisor);
                break;
            case "if":
                switch (commands[1].toLowerCase()) {
                    case "true": currentBuilder.setTrueTestThrow(Long.valueOf(commands[5])); break;
                    case "false": currentBuilder.setFalseTestThrow(Long.valueOf(commands[5])); break;
                }
                break;
            default:
                System.out.println("Unknown input type: "+commands[0]);
                break;
        }
        return currentBuilder;

    }

    public static class Monkey {
        List<Long> currentItems;
        Function<Long, Long> operation;
        Long testDivisor;
        Long testTrueThrowId;
        Long testFalseThrowId;
        Long inspectCount;

        public Monkey(List<Long> currentItems, Function<Long, Long> operation, Long test, Long testTrueThrowId, Long testFalseThrowId) {
            this.currentItems = currentItems;
            this.operation = operation;
            this.testDivisor = test;
            this.testTrueThrowId = testTrueThrowId;
            this.testFalseThrowId = testFalseThrowId;
            this.inspectCount = 0l;
        }

        public Long modifyWorry(Long originalWorry) {
            return operation.apply(originalWorry);
        }

        public void addItem(Long itemWorry) {
            currentItems.add(itemWorry);
        }

        public void inspect() {
            inspectCount++;
        }

        @Override
        public String toString() {
            return "Monkey Items: %s, TrueThrow: %s, FalseThrow: %s, InspectCount: %s, Divisor: %s"
                .formatted(currentItems, testTrueThrowId, testFalseThrowId, inspectCount, testDivisor);
        }
        
        
        public static class Builder {
            private List<Long> items;
            private Function<Long, Long> op;
            private Long test;
            private Long trueTest;
            private Long falseTest;

            public Builder() {
                this.items = new ArrayList<>();
                this.op = x -> x;
            }

            public Builder addItem(Long worryValue) {
                this.items.add(worryValue);
                return this;
            }

            public Builder setOperation(Function<Long, Long> operation) {
                this.op = operation;
                return this;
            }

            public Builder setTest(Long test) {
                this.test = test;
                return this;
            }

            public Builder setTrueTestThrow(Long trueTestId) {
                this.trueTest = trueTestId;
                return this;
            }

            public Builder setFalseTestThrow(Long falseTestId) {
                this.falseTest = falseTestId;
                return this;
            }

            public Monkey build() {
                if (this.test == null || this.falseTest == null || this.trueTest == null) {
                    throw new NullPointerException("Can't build monkey as the throwing decisions haven't been defined.");
                }

                return new Monkey(items, op, test, trueTest, falseTest);
            }
            
        }
    }
}
