package main
import ("fmt" )

var m = map[int] int {} 
func solve (n int) int { 
	if m[n] != 0 { 
		return m[n]
	}
	if n % 3 == 0 { 
		m[n] = solve(n/3) + 1 
	} 
	var s int
	if n % 2 == 0 { 
		s = solve(n/2) + 1 
		if m[n] == 0 || m[n] > s {
			m[n] = s
		}	
	} 
	s = solve(n-1) + 1 
	if m[n] ==0 || m[n] > s { 
		m[n] = s
	}
	return m[n]
} 


func main() {
	
	fmt.Println("%v\n", solve(9))
	
}


