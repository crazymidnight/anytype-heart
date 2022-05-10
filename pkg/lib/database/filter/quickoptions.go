package filter

import (
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
	timeutil "github.com/anytypeio/go-anytype-middleware/util/time"
	"time"
)

func TransformQuickOption(reqFilters *[]*model.BlockContentDataviewFilter) {
	if reqFilters == nil {
		return
	}

	for _, f := range *reqFilters {
		if f.QuickOption > model.BlockContentDataviewFilter_DateNone {

			d1, d2 := getRange(f)
			switch f.Condition {
			case model.BlockContentDataviewFilter_Equal:
				f.Condition = model.BlockContentDataviewFilter_GreaterOrEqual
				f.Value = pb.ToValue(d1)

				*reqFilters = append(*reqFilters, &model.BlockContentDataviewFilter{
					RelationKey: f.RelationKey,
					Condition:   model.BlockContentDataviewFilter_LessOrEqual,
					Value:       pb.ToValue(d2),
				})
			case model.BlockContentDataviewFilter_Less:
				f.Value = pb.ToValue(d1)
			case model.BlockContentDataviewFilter_Greater:
				f.Value = pb.ToValue(d2)
			case model.BlockContentDataviewFilter_LessOrEqual:
				f.Value = pb.ToValue(d2)
			case model.BlockContentDataviewFilter_GreaterOrEqual:
				f.Value = pb.ToValue(d1)
			case model.BlockContentDataviewFilter_In:
				f.Condition = model.BlockContentDataviewFilter_GreaterOrEqual
				f.Value = pb.ToValue(d1)

				*reqFilters = append(*reqFilters, &model.BlockContentDataviewFilter{
					RelationKey: f.RelationKey,
					Condition:   model.BlockContentDataviewFilter_LessOrEqual,
					Value:       pb.ToValue(d2),
				})
			}
			f.QuickOption = 0
		}
	}
}

func getRange(f *model.BlockContentDataviewFilter) (int64, int64) {
	var d1, d2 time.Time
	switch f.QuickOption {
	case model.BlockContentDataviewFilter_Yesterday:
		d1 = timeutil.DayNumStart(-1)
		d2 = timeutil.DayNumEnd(-1)
	case model.BlockContentDataviewFilter_Today:
		d1 = timeutil.DayNumStart(0)
		d2 = timeutil.DayNumEnd(0)
	case model.BlockContentDataviewFilter_Tomorrow:
		d1 = timeutil.DayNumStart(1)
		d2 = timeutil.DayNumEnd(1)
	case model.BlockContentDataviewFilter_LastWeek:
		d1 = timeutil.WeekNumStart(-1)
		d2 = timeutil.WeekNumEnd(-1)
	case model.BlockContentDataviewFilter_CurrentWeek:
		d1 = timeutil.WeekNumStart(0)
		d2 = timeutil.WeekNumEnd(0)
	case model.BlockContentDataviewFilter_NextWeek:
		d1 = timeutil.WeekNumStart(1)
		d2 = timeutil.WeekNumEnd(1)
	case model.BlockContentDataviewFilter_LastMonth:
		d1 = timeutil.MonthNumStart(-1)
		d2 = timeutil.MonthNumEnd(-1)
	case model.BlockContentDataviewFilter_CurrentMonth:
		d1 = timeutil.MonthNumStart(0)
		d2 = timeutil.MonthNumEnd(0)
	case model.BlockContentDataviewFilter_NextMonth:
		d1 = timeutil.MonthNumStart(1)
		d2 = timeutil.MonthNumEnd(1)
	case model.BlockContentDataviewFilter_NumberOfDaysAgo:
		daysCnt := f.Value.GetNumberValue()
		d1 = timeutil.DayNumStart(-int(daysCnt))
		d2 = timeutil.DayNumStart(0)
	case model.BlockContentDataviewFilter_NumberOfDaysNow:
		daysCnt := f.Value.GetNumberValue()
		d1 = timeutil.DayNumStart(0)
		d2 = timeutil.DayNumEnd(int(daysCnt))
	case model.BlockContentDataviewFilter_ExactDate:
		timestamp := f.GetValue().GetNumberValue()
		t := time.Unix(int64(timestamp), 0)
		d1 = timeutil.DayStart(t)
		d1 = timeutil.DayEnd(t)
	}

	return d1.Unix(), d2.Unix()
}
