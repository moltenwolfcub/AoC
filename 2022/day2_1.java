import helpers.FileUtils;

public class day2_1 {
	
	private static int runningScore;
    public static void main(String[] args) {
        runningScore = 0;
		for (String line : FileUtils.readLines("2022/day2.txt")) {
			runningScore += getGameResult(line.charAt(2), line.charAt(0));
		}
		System.out.println(runningScore);
    }

	private static Integer getGameResult(char me, char other) {
		Integer score = 0;

		switch (me) {
			case 'X':
			score += 1;
				switch (other) {
					case 'A': score += 3; break;
					case 'B': score += 0; break;
					case 'C': score += 6; break;
				
					default: break;
				}
				break;
			case 'Y': 
			score += 2;
				switch (other) {
					case 'A': score += 6; break;
					case 'B': score += 3; break;
					case 'C': score += 0; break;
				
					default: break;
				}
				break;
			case 'Z': 
			score += 3;
				switch (other) {
					case 'A': score += 0; break;
					case 'B': score += 6; break;
					case 'C': score += 3; break;
				
					default: break;
				}
				break;
			default: break;
		}
		return score;
	}
}
