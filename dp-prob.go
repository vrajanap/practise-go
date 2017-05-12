package main
import ("fmt")

type ans struct { 
	step int
	count int
} 

var m = map[int] ans {} 
func solve (n int) int { 
	if n == 1 { 
		return 0
	} 
	
	if m[n].count != 0 { 
		return m[n].count
	}
	if n % 3 == 0 { 
		a :=  ans{ solve(n/3) + 1, n/3} 
		m[n] =a 
	} 
	var s int
	if n % 2 == 0 { 
		s = solve(n/2) + 1 
		if m[n].count == 0 || m[n].count > s {
			a :=  ans{ s, n/2} 
			m[n] =a 
		}	
	} 
	s = solve(n-1) + 1 
	if m[n].count  ==0 || m[n].count > s { 
		a :=  ans{ s, n - 1} 
		m[n] =a 
	}
	return m[n].count
} 


func main() {
	var m1 int
	fmt.Scanf("%d", &m1)
	solve(m1) 
	for m1 >= 1 { 
		fmt.Println(m1)
		m1 = m[m1].step  
	} 
}


