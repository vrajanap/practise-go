package main
import ("fmt")

type ans struct { 
	count int
	step int
} 

var m = map[int] ans {} 
func solve (n int) int { 
	if n == 1 { 
		m[1] = ans{0, 0}
		return 0
	} 
	
	if m[n].count != 0 { 
		return m[n].count
	}
	if n % 3 == 0 { 
		a :=  ans{ solve(n/3) + 1, n/3} 
		fmt.Println("updating", n, "with", a, "old", m[n])
		m[n] =a 
	} 
	var s int
	if n % 2 == 0 { 
		s = solve(n/2) + 1 
		if m[n].count == 0 || m[n].count > s {
			a :=  ans{ s, n/2} 
			fmt.Println("updating", n, "with", a, "old", m[n])
			m[n] =a 
		}	
	} 
	s = solve(n-1) + 1 
	if m[n].count  ==0 || m[n].count > s { 
		a :=  ans{ s, n - 1} 
		fmt.Println("updating", n, "with", a, "old", m[n])
		m[n] =a
	}
	return m[n].count
} 


func main() {
	var m1 int
	fmt.Scanf("%d", &m1)
	solve(m1) 
	for m1 > 1 { 
		fmt.Println(m[m1].count, m[m1].step)
		m1 = m[m1].step  
	} 
}


