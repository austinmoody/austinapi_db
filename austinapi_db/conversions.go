package austinapi_db

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

func ConvertToSleep(rows []ListSleepRow) []Sleep {
	var sleeps []Sleep

	for _, row := range rows {
		sleep := row.Sleep()
		sleeps = append(sleeps, sleep)
	}

	return sleeps
}
