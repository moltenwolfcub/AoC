import java.util.HashMap;
import java.util.Map;

import Helpers.FileUtils;

public class day7_1 {
	private static Map<String, Map<String, Integer>> fileStructure;
	private static Map<String, Integer> dirSizes;
	private static Integer totalSize;

    public static void main(String[] args) {
		fileStructure = new HashMap<>();
		dirSizes = new HashMap<>();
		totalSize = 0;

		//read in file structure
        String currentDir = "/";
        for (String terminalLine : FileUtils.readLines("2022/day7.txt")) {
			String[] fullCommand = terminalLine.split(" ");

            if (fullCommand[0].equals("$")) {
				String command = fullCommand[1];
				if (command.contains("cd")) {
					currentDir = handleDirChange(currentDir, fullCommand[2]);
				}

            } else {
				if (!(fullCommand[0].equals("dir"))) {
					String filePath = fullCommand[1];
					Integer size = Integer.valueOf(fullCommand[0]);

					Map<String, Integer> newMap = fileStructure.getOrDefault(currentDir, new HashMap<String, Integer>());
					newMap.put(filePath, size);
					fileStructure.put(currentDir, newMap);
				}
            }
        }

		//calculate directory totals sizes
		for (String directory : fileStructure.keySet()) {
			Map<String, Integer> files = fileStructure.get(directory);
			for (String file : files.keySet()) {
				Integer size = files.get(file);

				String tmpDir = directory;
				while (true) {
					Integer prevValue = dirSizes.getOrDefault(tmpDir, 0);
					dirSizes.put(tmpDir, prevValue+size);
					if (tmpDir.equals("/")) {
						break;
					}
					tmpDir = removeLastDir(tmpDir);
				}
			}
		}

		//sum all the dirs that are less than or equal to 100,000 in size
		for (Integer fileSize : dirSizes.values()) {
			if (fileSize <= 100000) {
				totalSize += fileSize;
			}
		}

		System.out.println(totalSize);
    }

	private static String handleDirChange(String startingDir, String directoryParam) {
		String returnVal = startingDir;
		switch (directoryParam) {
			case "..": 
				returnVal = removeLastDir(startingDir);
				break;
			case "/": 
				returnVal = "/";
				break;
			default: 
				returnVal = returnVal.concat(directoryParam);
				returnVal = returnVal.concat("/");
				break;
		}
		return returnVal;
	}

	private static String removeLastDir(String input) {
		String[] splitDirs = input.split("/");
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
