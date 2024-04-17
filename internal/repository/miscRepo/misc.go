package miscrepo

import (
	"fmt"
	"time"

	"github.com/jcocozza/cassidy-wails/internal/database"
	"github.com/jcocozza/cassidy-wails/internal/model"
	"github.com/jcocozza/cassidy-wails/internal/sqlcode"
	"github.com/jcocozza/cassidy-wails/internal/utils"
	"github.com/jcocozza/cassidy-wails/internal/utils/dateutil"
	"github.com/jcocozza/cassidy-wails/internal/utils/measurement"
)

// Misc db queries
type MiscRepository interface {
	ReadNCycleSummary(startDate, endDate time.Time, userUuid string) (*model.NCycleSummary, error)
}

// Represents a database connection
type IMiscRepository struct {
	DB database.DbOperations
}

func NewIMiscRepository(db database.DbOperations) *IMiscRepository {
	return &IMiscRepository{
		DB: db,
	}
}

// Create the 12 cycle summary
//
// TODO: Generalize this function for n cycles
func (db *IMiscRepository) ReadNCycleSummary(startDate, endDate time.Time, userUuid string) (*model.NCycleSummary, error) {
	numCycles := 12
	cycleList := dateutil.GetPreviousNCycles(startDate, endDate, numCycles-1)

	cycleLength := len(cycleList[0])

	startDateList := []time.Time{}
	startDateList = append(startDateList, startDate)
	for _, cycle := range cycleList {
		startDateList = append(startDateList, cycle[0])
	}

	sql := sqlcode.SQLReader(sqlcode.N_cycle_summary)
	rows, err2 := db.DB.Query(sql, cycleList[10][0].Format(dateutil.Layout), cycleList[10][cycleLength-1].Format(dateutil.Layout),
		cycleList[9][0].Format(dateutil.Layout), cycleList[9][cycleLength-1].Format(dateutil.Layout),
		cycleList[8][0].Format(dateutil.Layout), cycleList[8][cycleLength-1].Format(dateutil.Layout),
		cycleList[7][0].Format(dateutil.Layout), cycleList[7][cycleLength-1].Format(dateutil.Layout),
		cycleList[6][0].Format(dateutil.Layout), cycleList[6][cycleLength-1].Format(dateutil.Layout),
		cycleList[5][0].Format(dateutil.Layout), cycleList[5][cycleLength-1].Format(dateutil.Layout),
		cycleList[4][0].Format(dateutil.Layout), cycleList[4][cycleLength-1].Format(dateutil.Layout),
		cycleList[3][0].Format(dateutil.Layout), cycleList[3][cycleLength-1].Format(dateutil.Layout),
		cycleList[2][0].Format(dateutil.Layout), cycleList[2][cycleLength-1].Format(dateutil.Layout),
		cycleList[1][0].Format(dateutil.Layout), cycleList[1][cycleLength-1].Format(dateutil.Layout),
		cycleList[0][0].Format(dateutil.Layout), cycleList[0][cycleLength-1].Format(dateutil.Layout),
		startDate.Format(dateutil.Layout), endDate.Format(dateutil.Layout),
		cycleList[10][0].Format(dateutil.Layout), endDate.Format(dateutil.Layout), userUuid)
	if err2 != nil {
		return nil, fmt.Errorf("failed to query n cycle summary: %w", err2)
	}
	defer rows.Close()

	binList := []int{}

	plannedDistances := []*measurement.Measurement{}
	plannedDurations := []float64{}
	plannedVerticals := []*measurement.Measurement{}

	completedDistances := []*measurement.Measurement{}
	completedDurations := []float64{}
	completedVerticals := []*measurement.Measurement{}

	var tmpBin int
	var tmpPlannedDistance float64
	var tmpPlannedDuration float64
	var tmpPlannedVertical float64
	var tmpCompletedDistance float64
	var tmpCompletedDuration float64
	var tmpCompletedVertical float64

	for rows.Next() {
		err3 := rows.Scan(&tmpBin,
			&tmpPlannedDistance, &tmpPlannedDuration, &tmpPlannedVertical,
			&tmpCompletedDistance, &tmpCompletedDuration, &tmpCompletedVertical)
		if err3 != nil {
			return nil, fmt.Errorf("failed scanning n cycle summary rows: %w", err3)
		}

		binList = append(binList, tmpBin)

		plannedDistances = append(plannedDistances, measurement.StandardMeasurement(tmpPlannedDistance))
		plannedDurations = append(plannedDurations, tmpPlannedDuration)
		plannedVerticals = append(plannedVerticals, measurement.StandardMeasurement(tmpPlannedVertical))

		completedDistances = append(completedDistances, measurement.StandardMeasurement(tmpCompletedDistance))
		completedDurations = append(completedDurations, tmpCompletedDuration)
		completedVerticals = append(completedVerticals, measurement.StandardMeasurement(tmpCompletedVertical))
	}
	finalPlannedDistances := createFinalList[*measurement.Measurement](plannedDistances, binList, numCycles, measurement.ZeroLength(measurement.StandardUnit))
	finalPlannedDurations := createFinalList[float64](plannedDurations, binList, numCycles, 0)
	finalPlannedVerticals := createFinalList[*measurement.Measurement](plannedVerticals, binList, numCycles, measurement.ZeroLength(measurement.StandardUnit))

	finalCompletedDistances := createFinalList[*measurement.Measurement](completedDistances, binList, numCycles, measurement.ZeroLength(measurement.StandardUnit))
	finalCompletedDurations := createFinalList[float64](completedDurations, binList, numCycles, 0)
	finalCompletedVerticals := createFinalList[*measurement.Measurement](completedVerticals, binList, numCycles, measurement.ZeroLength(measurement.StandardUnit))

	sdl := utils.ReverseList[time.Time](startDateList)

	return &model.NCycleSummary{
		StartDateList:      sdl,
		PlannedDistances:   finalPlannedDistances,
		PlannedDurations:   finalPlannedDurations,
		PlannedVerticals:   finalPlannedVerticals,
		CompletedDistances: finalCompletedDistances,
		CompletedDurations: finalCompletedDurations,
		CompletedVerticals: finalCompletedVerticals,
	}, nil
}

// fill in 0's where the database doesn't return anything for a given bin
func createFinalList[T any](initialList []T, indexList []int, totalLen int, zeroElm T) []T {
	finalList := []T{}
	for i := 0; i < totalLen; i++ {
		didAppend := false
		for j, idx := range indexList {
			if i == idx {
				finalList = append(finalList, initialList[j])
				didAppend = true
			}
			if didAppend {
				break
			}
		}
		if !didAppend {
			finalList = append(finalList, zeroElm)
		}
	}
	return finalList
}
