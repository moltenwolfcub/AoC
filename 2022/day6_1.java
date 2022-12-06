import java.util.HashMap;
import java.util.List;
import java.util.Map;

import Helpers.FileUtils;

public class day6_1 {
    private static List<Character> dataStream;

    private static int charsRead;
    
    public static void main(String[] args) {
        dataStream = FileUtils.readChars("2022/day6.txt");

        charsRead = 4;

        for (int i = 0; i < dataStream.size()-3; i++) {
            String output = "";
            for (int j = 0; j < 4; j++) {
                output+=String.valueOf(dataStream.get(i+j));
            }

            char[] selection = {0,0,0,0};
            output.getChars(0, 4, selection, 0);
            
            Map<Character, Integer> charOccurences = new HashMap<>();
            for(char data : selection) {
                if(charOccurences.containsKey(data)) {
                    int occurences = charOccurences.get(data);
                    charOccurences.put(data, ++occurences);
                } else {
                    charOccurences.put(data, 1);
                }
            }

            Boolean foundDupe = false;
            for(char data : charOccurences.keySet()) {
                if(charOccurences.get(data) > 1) {
                    foundDupe = true;
                    break;
                }
            }
            if (!foundDupe) {
                break;
            }

            charsRead++;
        }
        System.out.println(charsRead);
    }
}
