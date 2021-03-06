package constraint

import (
	"github.com/elgohr/stop-and-go/wait"
)

// NoOrder is used when the order of the Waiter does not matter to other Waiters
func NoOrder(w wait.Waiter) func(wts []wait.Waiter) []wait.Waiter {
	return func(wts []wait.Waiter) []wait.Waiter {
		return append(wts, w)
	}
}

// Before is used when a Waiter has to be called before another Waiter
func Before(first wait.Waiter, second wait.Waiter) func(wts []wait.Waiter) []wait.Waiter {
	return func(wts []wait.Waiter) []wait.Waiter {
		wts, fi := contains(wts, first)
		wts, si := contains(wts, second)
		return sort(wts, fi, si)
	}
}

func sort(wts []wait.Waiter, fi int, si int) []wait.Waiter {
	if fi > si {
		v := wts[fi]
		wts[fi] = wts[si]
		wts[si] = v
	}
	return wts
}

func contains(wts []wait.Waiter, nw wait.Waiter) ([]wait.Waiter, int) {
	for i, w := range wts {
		if w == nw {
			return wts, i
		}
	}
	wts = append(wts, nw)
	return wts, len(wts) - 1
}
