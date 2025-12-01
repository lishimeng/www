package menu

type Node[T any, K comparable] interface {
	AddChildren(T)
	GetId() K
	GetParentId() K
}

func contains[T comparable](slice []T, target T) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}
	return false
}

func BuildTree[T Node[T, K], K comparable](flat []T, pruning ...K) []T {
	m := make(map[K]T)
	for _, t := range flat {
		m[t.GetId()] = t
	}
	var roots []T
	var zero K
	for _, t := range flat {
		if contains(pruning, t.GetId()) {
			continue
		}
		if t.GetParentId() == zero {
			roots = append(roots, m[t.GetId()])
		} else {
			parent, exists := m[t.GetParentId()]
			if exists {
				parent.AddChildren(t)
			}
		}
	}

	return roots
}
