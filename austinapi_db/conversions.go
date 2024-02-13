package austinapi_db

type SleepInterface interface {
	Sleep() Sleep
}

func (lsr *ListSleepRow) Sleep() Sleep {
	return Sleep{
		ID:               lsr.ID,
		Date:             lsr.Date,
		Rating:           lsr.Rating,
		TotalSleep:       lsr.TotalSleep,
		DeepSleep:        lsr.DeepSleep,
		LightSleep:       lsr.LightSleep,
		RemSleep:         lsr.RemSleep,
		CreatedTimestamp: lsr.CreatedTimestamp,
		UpdatedTimestamp: lsr.UpdatedTimestamp,
	}
}

func (lsnr *ListSleepNextRow) Sleep() Sleep {
	return Sleep{
		ID:               lsnr.ID,
		Date:             lsnr.Date,
		Rating:           lsnr.Rating,
		TotalSleep:       lsnr.TotalSleep,
		DeepSleep:        lsnr.DeepSleep,
		LightSleep:       lsnr.LightSleep,
		RemSleep:         lsnr.RemSleep,
		CreatedTimestamp: lsnr.CreatedTimestamp,
		UpdatedTimestamp: lsnr.UpdatedTimestamp,
	}
}

func (lspr *ListSleepPreviousRow) Sleep() Sleep {
	return Sleep{
		ID:               lspr.ID,
		Date:             lspr.Date,
		Rating:           lspr.Rating,
		TotalSleep:       lspr.TotalSleep,
		DeepSleep:        lspr.DeepSleep,
		LightSleep:       lspr.LightSleep,
		RemSleep:         lspr.RemSleep,
		CreatedTimestamp: lspr.CreatedTimestamp,
		UpdatedTimestamp: lspr.UpdatedTimestamp,
	}
}

func SleepInterfaceToSleep(rows []SleepInterface) []Sleep {
	var sleeps []Sleep
	for _, row := range rows {
		sleeps = append(sleeps, row.Sleep())
	}

	return sleeps
}

func ListSleepPreviousRowToSleep(rows []ListSleepPreviousRow) []Sleep {
	var sleeps []Sleep
	for _, row := range rows {
		sleeps = append(sleeps, row.Sleep())
	}
	return sleeps
}

func ListSleepNextRowToSleep(rows []ListSleepNextRow) []Sleep {
	var sleeps []Sleep

	for _, row := range rows {
		sleeps = append(sleeps, row.Sleep())
	}

	return sleeps
}

func ListSleepRowToSleep(rows []ListSleepRow) []Sleep {
	var sleeps []Sleep

	for _, row := range rows {
		sleeps = append(sleeps, row.Sleep())
	}

	return sleeps
}
