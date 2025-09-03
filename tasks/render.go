package tasks

import (
	"fmt"
	"strings"
	"text/tabwriter"
	"time"
)

func RenderTable(t []Task) string {
	sb := strings.Builder{}
	w := tabwriter.NewWriter(&sb, 0, 0, 2, ' ', tabwriter.Debug)
	_, _ = fmt.Fprintln(w, "ID\tTitle\tPriority\tDue Date\tCategory\tStatus")
	_, _ = fmt.Fprintln(w, "----\t-----\t--------\t----------\t--------\t-------")
	for _, task := range t {
		var priority string
		switch task.Priority {
		case Low:
			priority = "LOW"
		case Medium:
			priority = "MEDIUM"
		case High:
			priority = "HIGH"
		}
		dueDate := time.Time(task.DueDate).Format("2006-01-02")
		var status string
		switch task.Status {
		case Pending:
			status = "pending"
		case Completed:
			status = "completed"
		}
		_, _ = fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\t%s\n", task.Id, task.Title, priority, dueDate, task.Category, status)
	}
	_ = w.Flush()
	return sb.String()
}
