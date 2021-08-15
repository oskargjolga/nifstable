package nifs

type ByPointsAndGoalDiff []TableEntry

func (a ByPointsAndGoalDiff) Len() int { return len(a) }
func (a ByPointsAndGoalDiff) Less(i, j int) bool {
	if a[i].Points == a[j].Points {
		return a[i].GoalDiff() > a[j].GoalDiff()
	}
	return a[i].Points > a[j].Points
}
func (a ByPointsAndGoalDiff) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
