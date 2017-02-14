package main

import "fmt"

func main(){
var tall []int = []int{0,2,1,5,3,4}
fmt.Println("Tallene som skal sorteres:", tall)
bubbleSort(tall)
fmt.Println("Tallene etter sortering:", tall)
}

func bubbleSort(tall []int){
var N = len(tall)
var i int
for i = 0; i < N; i++ {
  if !sweep(tall, i){
    return
  }
  }
}

func sweep(tall []int, prevPasses int) bool{
var N int = len(tall)
var forsteIndex int = 0
var andreIndex int = 1
var didSwap bool = false

for andreIndex < (N - prevPasses) {
  var forsteTall int = tall[forsteIndex]
  var andreTall int = tall[andreIndex]
  didSwap = true
  if forsteTall > andreTall {
  //bytte plassen på numrene
  tall[forsteIndex] = andreTall
  tall[andreIndex] = forsteTall
  }
  //her skal det jobbes
  forsteIndex++
  andreIndex++
  }
  return didSwap
}
