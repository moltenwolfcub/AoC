import Helpers.FileUtils;

public class Day2_2 {
	private static int runningScore;
    
    public static void main(String[] args) {
        runningScore = 0;
		for (String line : FileUtils.readLines("2022/day2.txt")) {
            char otherPlay = line.charAt(0);

            char myPlay = ' ';
            switch (line.charAt(2)) {
                case 'X'://lose
                    switch (otherPlay) {
                        case 'A': myPlay='Z'; break;
                        case 'B': myPlay='X'; break;
                        case 'C': myPlay='Y'; break;
                    }
                    break;
                case 'Y'://draw
                    switch (otherPlay) {
                        case 'A': myPlay='X'; break;
                        case 'B': myPlay='Y'; break;
                        case 'C': myPlay='Z'; break;
                    }
                    break;
                case 'Z'://win
                    switch (otherPlay) {
                        case 'A': myPlay='Y'; break;
                        case 'B': myPlay='Z'; break;
                        case 'C': myPlay='X'; break;
                    }
                    break;
                default: break;
            }

			runningScore += getGameResult(myPlay, otherPlay);
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
