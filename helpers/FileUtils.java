package helpers;

import java.io.File;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;
import java.util.stream.Collectors;

public class FileUtils {
    public static List<String> readLines(String fileName) {
        List<String> lines = new ArrayList<>();

        try {
            File file = new File(fileName);

            Scanner fileReader = new Scanner(file);
            while (fileReader.hasNextLine()) {
                lines.add(fileReader.nextLine());
            }
            fileReader.close();

        } catch (IOException e) {
            System.out.println("An Error Occured while reading in the file.");
            e.printStackTrace();
            System.out.println("The data being returned will be incomplete or incorrect.");
        }
        
        return lines;
    }
    public static List<Character> readChars(String fileName) {
        List<String> lines = readLines(fileName);
        List<Character> chars = new ArrayList<>();
        for (String line : lines) {
            chars.addAll(line.chars().mapToObj(i -> (char)i).collect(Collectors.toList()));
        }
        return chars;
    }
    public static List<String> splitCSV(String input) {
        String str = input.strip();
        String[] split = str.split(",");

        List<String> splitList = new ArrayList<>();
        for (String splitValues : split) {
            splitValues = splitValues.strip();
            splitList.add(splitValues);
        }
        return splitList;
    }

	public static String removeLastDir(String inputDir) {
		String[] splitDirs = inputDir.split("/");
		String returnVal = "";
		for (int i = 1; i < splitDirs.length-1; i++) {
			returnVal = returnVal.concat("/");
			returnVal = returnVal.concat(splitDirs[i]);
		}

		if (returnVal != "/") {
			returnVal = returnVal.concat("/");
		}

		return returnVal;
		
	}
}
