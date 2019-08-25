package main

import "fmt"

type Game struct {
	scoreA int	
	scoreB int	
}

func (ga Game) CurrentScore() string {
	
	ps := []string{
			"Love","Fifteen","Thirty","Fourty",
	}

	tmpA := ""
	tmpB := ""
	
	for i:=0 ; i < 3; i++ {
		if ga.scoreA == i {			
			tmpA = ps[i]
		}	
		if ga.scoreB == i {			
			tmpB = ps[i]
		}				
	}

	if tmpA == "" && tmpA == tmpB {
		return "Deuce"
	}

	if tmpA == ""  && (ga.scoreA - ga.scoreB) >= 2 {
		return "A Wins"
	} 

	if tmpA == ""  && (ga.scoreB - ga.scoreA) >= 2 {
		return "B Wins"
	} 

	if tmpA == "" && ga.scoreA > ga.scoreB {
		return "A Advantage"
	}

	if tmpA == "" && ga.scoreB > ga.scoreA {
		return "B Advantage"
	}

	if tmpA == tmpB {
		return tmpA + " Both"
	}

	return tmpA + " - " + tmpB
}

func (ga *Game) PlayerAGetPoint() {
	ga.scoreA++	
}

func (ga *Game) PlayerBGetPoint() {
	ga.scoreB++
}


func main() {
	g := Game{}	
	g.PlayerAGetPoint()
	fmt.Printf("Current Score : %s\n",g.CurrentScore())
	g.PlayerBGetPoint()
	fmt.Printf("Current Score : %s\n",g.CurrentScore())	
	g.PlayerAGetPoint()
	fmt.Printf("Current Score : %s\n",g.CurrentScore())
	g.PlayerAGetPoint()
	fmt.Printf("Current Score : %s\n",g.CurrentScore())
}