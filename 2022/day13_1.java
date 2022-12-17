import java.security.InvalidParameterException;
import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

import Helpers.FileUtils;

public class day13_1 {
    private static Integer correctPacketIdSum;
    
    public static void main(String[] args) {
        correctPacketIdSum = 0;

        PacketItem otherPacket = null;
        Integer packetPairHalf = 0;
        Integer pairId = 0;
        for (String packet : FileUtils.readLines("2022/day13.txt")) {
            List<Character> inputChars = packet.chars().mapToObj(i -> (char)i).collect(Collectors.toList());
            if (inputChars.size() == 0) {
                packetPairHalf = 0;
                continue;
            }

            inputChars.remove(0);
            PacketItem parsedPacket = parseList(inputChars);

            if (packetPairHalf == 0) {
                otherPacket = parsedPacket;
                packetPairHalf++;
                continue;
            }
            if (packetPairHalf != 1) {
                throw new IndexOutOfBoundsException("PacketHalf wasn't 1 or 0 and there wasn't an empty line.");
            }

            pairId++;
            if (parsedPacket.compareTo(otherPacket)>0) {
                correctPacketIdSum+=pairId;
            }
            otherPacket = null;
            parsedPacket = null;
        }
        System.out.println(correctPacketIdSum);
    }


    /**
     * Parses a string csv list to a {@code PacketItem} object
     * 
     * @param inputString       a csv characterList of Integers or sublists
     *                          of integers. E.G. 1,5,2,[2,4,[4,5],9],1
     * @return                  the read in list as a {@code PacketItem}
     *                          representation
     */
    private static PacketItem parseList(List<Character> inputChars) {
        PacketItem returnPacket = new PacketItem();

        String currentInteger = "";
        while (inputChars.size() > 0) {
            Character character = inputChars.get(0);
            inputChars.remove(0);

            //characters to parse: 1234567890,[]
            if (Character.isDigit(character)) {
                //append digit char to the end of integer string
                currentInteger = currentInteger.concat(character.toString());
            } else if(character == ',') {
                //check there is an integer stored
                if (currentInteger.length() == 0) {
                    continue;
                }
                //add Integer to list
                Integer currentValue = Integer.valueOf(currentInteger);
                PacketItem integerPacketItem = new PacketItem(currentValue);
                returnPacket.addSubItem(integerPacketItem);

                //reset integer build
                currentInteger = "";
            } else if(character == '[') {
                returnPacket.addSubItem(parseList(inputChars));
            } else if (character == ']') {
                //add Integer to list
                if (currentInteger.length() != 0) {
                    Integer currentValue = Integer.valueOf(currentInteger);
                    PacketItem integerPacketItem = new PacketItem(currentValue);
                    returnPacket.addSubItem(integerPacketItem);
    
                    currentInteger = "";
                }

                //return the current packet
                return returnPacket;
            } else {
                throw new IllegalArgumentException("Recieved a character to parser that wasn't a digit or one of ,\\[\\]"); 
            }
        }
        throw new IllegalArgumentException("Never found a ']' to end parsing."); 
    }

    private static class PacketItem implements Comparable<PacketItem> {
        private Integer value;
        private List<PacketItem> list;
        private ItemType type;

        public PacketItem() {
            //outer packet
            type = ItemType.LIST;
            list = new ArrayList<>();
        }
        public PacketItem(Integer val) {
            value = val;
            type = ItemType.VALUE;
        }
        public PacketItem(List<PacketItem> list) {
            this.list = list;
            type = ItemType.LIST;
        }

        public void addSubItem(PacketItem subItem) {
            if (type == ItemType.VALUE) {
                throw new NullPointerException("Tried to add a subItem to a packetItem that wasn't a list.");
            }
            list.add(subItem);
        }

        @Override
        public int compareTo(PacketItem other) {
            /*
             * -1 is less than
             *  0 is equal
             *  1 is greater than
             */

            if(other.type == ItemType.VALUE && this.type == ItemType.VALUE) {
                return this.value.compareTo(other.value);
            } else if (this.type != ItemType.LIST) {
                PacketItem tmp = new PacketItem(List.of(this));
                return tmp.compareTo(other);
            } else if (other.type != ItemType.LIST) {
                PacketItem tmp = new PacketItem(List.of(other));
                return this.compareTo(tmp);
            } else if (other.type == ItemType.LIST && this.type == ItemType.LIST) {
                for (int i = 0; i < this.list.size(); i++) {
                    if (i>= other.list.size()) {
                        return 1;
                    }
                    Integer result = this.list.get(i).compareTo(other.list.get(i));
                    if (result != 0) {
                        return result;
                    }
                }
                if (other.list.size() == this.list.size()) {
                    return 0;
                }
                return -1;
            } else {
                throw new InvalidParameterException("Something is very wrong. The values parsed failed all comparisons.");
            }
        }


        @Override
        public String toString() {
            return "%s packetItem of %s".formatted(type, type==ItemType.VALUE ? value : list.toArray());
        }

        private enum ItemType{
            VALUE,
            LIST;
        }
    }
}
