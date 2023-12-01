import java.util.HashMap;
import java.util.LinkedHashMap;
import java.util.Map;
import java.util.stream.Collectors;

import helpers.FileUtils;

public class day7_2 {
	private static Map<String, Map<String, Integer>> fileStructure;
	private static Map<String, Integer> dirSizes;

    public static void main(String[] args) {
		fileStructure = new HashMap<>();
		dirSizes = new HashMap<>();

		//region: read in file structure
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
        //endregion

		//region: calculate directory totals sizes
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
					tmpDir = FileUtils.removeLastDir(tmpDir);
				}
			}
		}
        //endregion

		//region: find the space required
        Integer totalUsed = dirSizes.get("/");
        Integer totalFree = 70000000 - totalUsed;
        Integer spaceReq = 30000000 - totalFree;
        //endregion

        //region: locate smallest file that's big enough
        Map<String, Integer> orderedDirSizes = dirSizes.entrySet().stream()
            .sorted(Map.Entry.comparingByValue())
            .collect(Collectors.toMap(Map.Entry::getKey, Map.Entry::getValue, (e1, e2) -> e1, LinkedHashMap::new));

        Integer smallestDirSize = 0;
        for (Map.Entry<String, Integer> dir : orderedDirSizes.entrySet()) {
            if (dir.getValue() >= spaceReq) {
                smallestDirSize = dir.getValue();
                break;
            }
        }
        //endregion

		System.out.println(smallestDirSize);
    }

	private static String handleDirChange(String startingDir, String directoryParam) {
		String returnVal = startingDir;
		switch (directoryParam) {
			case "..": 
				returnVal = FileUtils.removeLastDir(startingDir);
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
    
}
