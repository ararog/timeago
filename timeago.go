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

func TimeAgo(start time.Time, end time.Time) string {

  duration := start.Sub(end)

  if duration.Hours() < 24 {
    if duration.Hours() >= 1 {
      return localizedStringFor(HoursAgo, int(duration.Hours()));
    } else if duration.Minutes() >= 1 {
      return localizedStringFor(MinutesAgo, int(duration.Minutes()));
    } else {
      return localizedStringFor(SecondsAgo, int(duration.Seconds()));
    }
  } else {
    if duration.Hours() >= 8760 {
      years := duration.Hours() / 8760
      return localizedStringFor(YearsAgo, int(years));
    } else if duration.Hours() >= 730 {
      months := duration.Hours() / 730
      return localizedStringFor(MonthsAgo, int(months));
    } else if duration.Hours() >= 168 {
      weeks := duration.Hours() / 168
      return localizedStringFor(WeeksAgo, int(weeks));
    } else {
      days := duration.Hours() / 24
      return localizedStringFor(DaysAgo, int(days));
    }
  }
}

func localizedStringFor(valueType DateAgoValues, value int) string {

    switch valueType {
        case YearsAgo:
            if value >= 2 {
                return fmt.Sprintf("%d years ago", value);
            } else {
                return "Last year";
            }
        case MonthsAgo:
            if value >= 2 {
                return fmt.Sprintf("%d months ago", value);
            } else {
                return "Last month";
            }
        case WeeksAgo:
            if value >= 2 {
                return fmt.Sprintf("%d weeks ago", value);
            } else {
                return "Last week";
            }
        case DaysAgo:
            if value >= 2 {
                return fmt.Sprintf("%d days ago", value);
            } else {
                return "Yesterday";
            }
        case HoursAgo:
            if value >= 2 {
                return fmt.Sprintf("%d hours ago", value);
            } else {
                return "An hour ago";
            }
        case MinutesAgo:
            if value >= 2 {
                return fmt.Sprintf("%d minutes ago", value);
            } else {
                return "A minute ago";
            }
        case SecondsAgo:
            if value >= 2 {
                return fmt.Sprintf("%d seconds ago", value);
            } else {
                return "Just now";
            }
    }
    return "";
}
