package austinapi_db

func (a *SleepsRow) Sleep() Sleep {
	return Sleep{
		ID:               a.ID,
		Date:             a.Date,
		Rating:           a.Rating,
		TotalSleep:       a.TotalSleep,
		DeepSleep:        a.DeepSleep,
		LightSleep:       a.LightSleep,
		RemSleep:         a.RemSleep,
		CreatedTimestamp: a.CreatedTimestamp,
		UpdatedTimestamp: a.UpdatedTimestamp,
	}
}

func AustinRowToSleep(rows []SleepsRow) []Sleep {
	var sleeps []Sleep

	for _, row := range rows {
		sleeps = append(sleeps, row.Sleep())
	}

	return sleeps
}
