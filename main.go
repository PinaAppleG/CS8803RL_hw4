package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type Transition struct {
	Id int `json:"id"`
	Probability float32 `json:"probability"`
	Reward int `json:"reward"`
	To int `json:"to"`
}

type Action struct {
	Id int `json:"id"`
	Transitions []Transition `json:"transitions"`
}

type State struct {
	Id int `json:"id"`
	Actions []Action `json:"actions"`
}

type MDP struct {
	Gamma float32 `json:"gamma"`
	States []State `json:"states"`
}

func main() {
	rand.Seed(time.Now().Unix())
	
	var numStates = 30
	var numActions = 2

	mdp := &MDP{
		Gamma: 0.99,
		States: make([]State, numStates),
	}

	for i := 0; i < numStates; i++ {
		state := &State{
			Id: i,
			Actions: make([]Action, numActions),
		}
		mdp.States[i] = *state

		for j := 0; j < numActions; j++ {
			probability1 := rand.Float32()
			probability2 := 1.0 - probability1

			transition1 := &Transition{
				Id: 0,
				Probability: probability1,
				Reward: rand.Int()%100,
				To: rand.Int() % numStates,
			}

			transition2 := &Transition{
				Id: 1,
				Probability: probability2,
				Reward: -1 * rand.Int()%100,
				To: rand.Int() % numStates,
			}

			mdp.States[i].Actions[j] = Action{
				Id: j,
				Transitions: make([]Transition, 2),
			}

			mdp.States[i].Actions[j].Transitions[0] = *transition1
			mdp.States[i].Actions[j].Transitions[1] = *transition2
			
		}
		
	}

	

	ret, _ := json.Marshal(mdp)
	
	fmt.Println(string(ret))
}
