package model

import (
  "time"
)

// TimeEntry describes the task being tracked:
// Category (meetings, Internal, Customer#, ...)
// Jira/Case/Ticket #
// What? (short description)
// Start Time
// End Time
type TimeEntry struct {
  ID int64
  Category  string
  IssueNumber string
  Description string
  StartTime time.Time
  EndTime time.Time
}
