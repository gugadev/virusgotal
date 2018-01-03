package models

import (
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

/*
ScanResult represents a scan of an antivitus
*/
type ScanResult struct {
	Detected bool
	Version  string
	Result   string
}

/*
Report represents the /report response
*/
type Report struct {
	Code      int `json:"response_code"`
	Positives int
	Total     int
	Scans     map[string]ScanResult
}

/*
Show show the data inside a table
*/
func (m *Report) Show() {
	table := tablewriter.NewWriter(os.Stdout)
	header := []string{
		"AV", "Version", "Result", "Detected",
	}
	var data [][]string
	for key, scan := range m.Scans {
		data = append(data, []string{
			key,
			scan.Version,
			scan.Result,
			strconv.FormatBool(scan.Detected),
		})
	}
	footer := []string{
		"Detected", strconv.Itoa(m.Positives),
		"Total", strconv.Itoa(m.Total),
	}
	table.SetBorder(false)
	table.SetAlignment(1)
	table.SetHeader(header)
	table.SetFooter(footer)
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgRedColor},
	)
	table.SetFooterColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgWhiteColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgWhiteColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgRedColor},
	)
	table.AppendBulk(data)
	table.Render()
}
