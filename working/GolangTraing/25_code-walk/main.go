package main

import (
	"fmt"
	"math/rand"
)

const (
	win            = 100
	gamesPerSeries = 10
)

type score struct {
	player, opponent, thisTurn int
}

type action func(current score) (result score, turnIsOver bool)

func roll(s score) (score, bool) {
	outcome := rand.Intn(6) + 1
	if outcome == 1 {
		return score{s.opponent, s.player, 0}, true
	}
	return score{s.player, s.opponent, s.thisTurn + outcome}, false
}
func stay(s score) (score, bool) { //stay函数一个入参s，一个出参score---结构体类型
	return score{s.opponent, s.player + s.thisTurn, 0}, true
}

type strategy func(score) action

func stayAtk(k int) strategy { //stayAtk函数一个入参为k，返回类型为strategy--函数类型
	return func(s score) action {
		if s.thisTurn >= k {
			return stay //返回stay函数
		}
		return roll
	}
}
func play(strategy0, strategy1 strategy) int {
	strategies := []strategy{strategy0, strategy1}
	var s score
	var turnIsOver bool
	currentPlayer := rand.Intn(2)
	for s.player+s.thisTurn < win {
		action := strategies[currentPlayer](s)
		s, turnIsOver = action(s)
		if turnIsOver {
			currentPlayer = (currentPlayer + 1) % 2
		}
	}
	return currentPlayer
}

func roundRobin(strategies []strategy) ([]int, int) {
	wins := make([]int, len(strategies))
	for i := 0; i < len(strategies); i++ {
		for j := i + 1; j < len(strategies); j++ {
			for k := 0; k < gamesPerSeries; k++ {
				winner := play(strategies[i], strategies[j])
				if winner == 0 {
					wins[i]++
				} else {
					wins[j]++
				}
			}
		}
	}
	gamesPerStrategy := gamesPerSeries * (len(strategies) - 1)
	return wins, gamesPerStrategy
}
func ratioString(vals ...int) string {
	total := 0
	for _, val := range vals {
		total += val
	}
	s := ""
	for _, val := range vals {
		if s != "" {
			s += ","
		}
		pct := 100 * float64(val) / float64(total)
		s += fmt.Sprintf("%d/%d(%0.1f%%)", val, total, pct)
	}
	return s
}

func main() {
	strategies := make([]strategy, win) //初始化一个[]strategy，容量为win=100

	fmt.Println("strategies", strategies)

	for k := range strategies { //遍历循环切片 strategies
		strategies[k] = stayAtk(k + 1) //给strategies里的元素循环赋值
	}
	wins, games := roundRobin(strategies)
	for k := range strategies {
		fmt.Printf("wins,losses staying at k=% 4d:%s\n",
			k+1, ratioString(wins[k], games-wins[k]))
	}
}
