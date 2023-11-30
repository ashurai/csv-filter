package pkg

import "sort"

// Comparer struct is used to compare Installation instances for distinct filtering
type Comparer struct{}

// Equals method is to cpmaire all components
func (c Comparer) Equals(i, j Installation) bool {
	return i.ComputerID == j.ComputerID &&
		i.UserID == j.UserID &&
		i.ApplicationID == j.ApplicationID &&
		i.ComputerType == j.ComputerType &&
		i.Comment == j.Comment
}

// ContainsComputerType checks if the group contains a specific computer type
func ContainsComputerType(group []Installation, computerType string) bool {
	for _, installation := range group {
		if installation.ComputerType == computerType {
			return true
		}
	}
	return false
}

// Filter to indetify conditions
func Filter(rows []Installation) int64 {
	// Filter and process data
	var applicationIDQuery []Installation
	for _, row := range rows {
		if row.ApplicationID == 374 {
			applicationIDQuery = append(applicationIDQuery, row)
		}
	}

	sort.Slice(applicationIDQuery, func(i, j int) bool {
		return applicationIDQuery[i].UserID < applicationIDQuery[j].UserID
	})

	// Group/list by UserID
	groupUsers := make(map[int][]Installation)
	for _, installation := range applicationIDQuery {
		groupUsers[installation.UserID] = append(groupUsers[installation.UserID], installation)
	}

	var copiesNeeded int64

	for _, group := range groupUsers {
		groupSize := len(group)
		if groupSize == 1 {
			copiesNeeded++
		} else if groupSize == 2 {
			if ContainsComputerType(group, "LAPTOP") && ContainsComputerType(group, "DESKTOP") {
				copiesNeeded++
			} else if ContainsComputerType(group, "LAPTOP") && ContainsComputerType(group, "LAPTOP") {
				copiesNeeded++
			} else {
				copiesNeeded += 2
			}
		}
	}

	return copiesNeeded
}
