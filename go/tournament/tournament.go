package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

// Team represents matches played, wins, draws, losses, and points of one particular team.
type Team struct {
	name                                 string
	matches, wins, draws, losses, points int
}

// Tally writes table of team results based on individual matches results.
func Tally(reader io.Reader, writer io.Writer) error {
	var err error
	scanner := bufio.NewScanner(reader)
	teams := make(map[string]Team)
	for scanner.Scan() {
		row := strings.TrimSpace(scanner.Text())
		// Ignore empty entries and comments.
		if row == "" || strings.HasPrefix(row, "#") {
			continue
		}
		rowSlices := strings.Split(row, ";")
		if len(rowSlices) != 3 {
			return errors.New("invalid number of fields in input data")
		}
		err = makeTeams(rowSlices, teams)
		if err != nil {
			return err
		}
	}
	makeTable(teams, writer)
	return err
}

// makeTeams makes map of team results based on input textual data.
func makeTeams(line []string, teams map[string]Team) error {
	var err error
	a, b := teams[line[0]], teams[line[1]]
	a.name = line[0]
	b.name = line[1]
	a.matches++
	b.matches++
	switch line[2] {
	case "win":
		a.wins++
		a.points += 3
		b.losses++
	case "loss":
		b.wins++
		b.points += 3
		a.losses++
	case "draw":
		a.draws++
		a.points++
		b.draws++
		b.points++
	default:
		err = errors.New("invalid result of match")
	}
	teams[line[0]], teams[line[1]] = a, b
	return err
}

// makeTable makes sorted table of Team from provided map.
func makeTable(teams map[string]Team, writer io.Writer) {
	fmt.Fprintf(writer, "Team                           | MP |  W |  D |  L |  P\n")
	// Copy data from map into slice to sort it.
	table := make([]Team, 0, len(teams))
	for _, team := range teams {
		table = append(table, team)
	}
	sortSlice(table)
	for _, row := range table {
		fmt.Fprintf(writer,
			"%-30s | %2d | %2d | %2d | %2d | %2d\n", row.name, row.matches, row.wins, row.draws, row.losses, row.points)
	}
}

// sortSlice sorts by points from highest to lowest, then by name alphabetically.
func sortSlice(teamSlice []Team) {
	sort.Slice(teamSlice, func(i, j int) bool {
		ti, tj := teamSlice[i], teamSlice[j]
		if ti.points != tj.points {
			return ti.points > tj.points
		}
		// If points are equal, sort by names alphabetically.
		return ti.name < tj.name
	})
}
