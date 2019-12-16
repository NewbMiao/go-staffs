package main

func main() {

}

/*
Some details about iterations over maps are listed here.
For a map, the entry order in an iteration is not guaranteed to be the same as the next iteration, even if the map is not modified between the two iterations. By Go specification, the order is unspecified (kind-of randomized).
If a map entry (a key-element pair) which has not yet been reached is removed during an iteration, then the entry will not iterated in the same iteration for sure.
If a map entry is created during an iteration, that entry may be iterated during the same iteration, or not.
*/
