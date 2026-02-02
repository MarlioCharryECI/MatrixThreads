package game

import "time"

const (
	phoneWeight = 4
	agentWeight = 3
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func manhattan(a, b Position) int {
	return abs(a.X-b.X) + abs(a.Y-b.Y)
}

func possibleMoves(pos Position) []Position {
	return []Position{
		{pos.X - 1, pos.Y},
		{pos.X + 1, pos.Y},
		{pos.X, pos.Y - 1},
		{pos.X, pos.Y + 1},
	}
}

func moveTowards(w *World, from, target Position) Position {
	best := from
	bestDist := manhattan(from, target)

	for _, next := range possibleMoves(from) {
		if !IsValid(w, next) {
			continue
		}
		if d := manhattan(next, target); d < bestDist {
			bestDist = d
			best = next
		}
	}
	return best
}

func closestAgentDistance(pos Position, agents []Position) int {
	min := 1 << 30
	for _, a := range agents {
		if d := manhattan(pos, a); d < min {
			min = d
		}
	}
	return min
}

func closestPhone(pos Position, phones []Position) Position {
	best := phones[0]
	bestDist := manhattan(pos, phones[0])

	for _, p := range phones[1:] {
		if d := manhattan(pos, p); d < bestDist {
			bestDist = d
			best = p
		}
	}
	return best
}

func moveNeoSmart(w *World) Position {
	target := closestPhone(w.Neo, w.Phones)

	best := w.Neo
	bestScore := 1 << 30

	for _, next := range possibleMoves(w.Neo) {
		if !IsValid(w, next) {
			continue
		}

		score :=
			manhattan(next, target)*phoneWeight -
				closestAgentDistance(next, w.Agents)*agentWeight

		if score < bestScore {
			bestScore = score
			best = next
		}
	}
	return best
}

func NeoRoutine(w *World, done chan string) {
	for {
		time.Sleep(700 * time.Millisecond)

		w.Mutex.Lock()
		w.Neo = moveNeoSmart(w)

		for i, p := range w.Phones {
			if w.Neo == p {
				w.Phones = append(w.Phones[:i], w.Phones[i+1:]...)
				if len(w.Phones) == 0 {
					w.Mutex.Unlock()
					done <- "¡Victoria! Neo ha conseguido todos los teléfonos."
					return
				}
				break
			}
		}

		for _, a := range w.Agents {
			if w.Neo == a {
				w.Mutex.Unlock()
				done <- "¡Derrota! Los agentes han atrapado a Neo."
				return
			}
		}

		w.placeEntities()
		w.Mutex.Unlock()
	}
}

func AgentRoutine(w *World, index int, done chan string) {
	for {
		time.Sleep(900 * time.Millisecond)

		w.Mutex.Lock()
		w.Agents[index] = moveTowards(w, w.Agents[index], w.Neo)

		if w.Agents[index] == w.Neo {
			w.Mutex.Unlock()
			done <- "¡Derrota! Un agente ha atrapado a Neo."
			return
		}

		w.placeEntities()
		w.Mutex.Unlock()
	}
}
