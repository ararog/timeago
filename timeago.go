package timeago

import (
	"time"
  "fmt"
)

type DateAgoValues int

const (
    SecondsAgo DateAgoValues = iota
    MinutesAgo
    HoursAgo
    DaysAgo
    WeeksAgo
    MonthsAgo
    YearsAgo
)


// Reverse returns its argument string reversed rune-wise left to right.
func TimeAgo(start Time, end Time) string {

  duration := start.Sub(end)

  if duration.Hours() < 24 {
    if duration.Hours() >= 1 {
      return localizedStringFor(HoursAgo, duration.Hours());
    } else if duration.Minutes() >= 1 {
      return localizedStringFor(MinutesAgo, duration.Minutes());
    } else {
      return localizedStringFor(SecondsAgo, duration.Seconds());
    }
  }
  else {
    if duration.Hours() >= 8760 {
        return localizedStringFor(YearsAgo, difference.year);
    } else if duration.Hours() >= 730 {
        return localizedStringFor(MonthsAgo, difference.month);
    } else if duration.Hours() >= 168 {
        return localizedStringFor(WeeksAgo, difference.weekOfYear);
    } else {
        return localizedStringFor(DaysAgo, difference.day);
    }
  }
}

func localizedStringFor(valueType DateAgoValues, value int) string {

    switch valueType {
        case YearsAgo:
            if value >= 2 {
                return fmt.Sprintf("%%d %@years ago", value);
            } else {
                return "Last year";
            }
        case MonthsAgo:
            if value >= 2 {
                return fmt.Sprintf("%%d %@months ago", value);
            } else {
                return "Last month";
            }
        case WeeksAgo:
            if value >= 2 {
                return fmt.Sprintf("%%d %@weeks ago", value);
            } else {
                return "Last week";
            }
        case DaysAgo:
            if value >= 2 {
                return fmt.Sprintf("%%d %@days ago", value);
            } else {
                return "Yesterday";
            }
        case HoursAgo:
            if value >= 2 {
                return fmt.Sprintf("%%d %@hours ago", value);
            } else {
                return "An hour ago";
            }
        case MinutesAgo:
            if value >= 2 {
                return fmt.Sprintf("%%d %@minutes ago", value);
            } else {
                return "A minute ago";
            }
        case SecondsAgo:
            if value >= 2 {
                return fmt.Sprintf("%%d %@seconds ago", value);
            } else {
                return "Just now";
            }
    }
    return nil;
}
