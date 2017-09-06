package greedy;

public class SelectActivity {

	public static void findMaxActivity(int start[], int finish[]){
		int i = 0;
		
		System.out.println(i);
		
		for (int j = 1; j< start.length; j++){
			if(start[j] >= finish[i]){
				System.out.println(j);
				i = j;
			}
		}
	}
	
	public static void main(String[] args) {
		// TODO Auto-generated method stub

		int start_time[] = {1, 3, 0, 5, 8, 5};
		int finish_time[] = {2, 4, 6, 7, 9, 9};
	
		findMaxActivity(start_time, finish_time);	
	}
}
