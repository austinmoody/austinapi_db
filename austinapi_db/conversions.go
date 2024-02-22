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

//func GetSqid(sqids *sqids.Sqids, id int64) string {
//	sqid, _ := sqids.Encode([]uint64{uint64(id)})
//	return sqid
//}
//
//func (sr *SleepsRow) MarshalJSON(sqids *sqids.Sqids) ([]byte, error) {
//	return json.Marshal(struct {
//		ID               string    `json:"id"`
//		Date             time.Time `json:"date"`
//		Rating           int64     `json:"rating"`
//		TotalSleep       int       `json:"total_sleep"`
//		DeepSleep        int       `json:"deep_sleep"`
//		LightSleep       int       `json:"light_sleep"`
//		RemSleep         int       `json:"rem_sleep"`
//		CreatedTimestamp time.Time `json:"created_timestamp"`
//		UpdatedTimestamp time.Time `json:"updated_timestamp"`
//		PreviousID       int64     `json:"previous_id"`
//		NextID           int64     `json:"next_id"`
//	}{
//		ID:               GetSqid(sqids, sr.ID),
//		Date:             sr.Date,
//		Rating:           sr.Rating,
//		TotalSleep:       sr.TotalSleep,
//		DeepSleep:        sr.DeepSleep,
//		LightSleep:       sr.LightSleep,
//		RemSleep:         sr.RemSleep,
//		CreatedTimestamp: sr.CreatedTimestamp,
//		UpdatedTimestamp: sr.UpdatedTimestamp,
//		PreviousID:       sr.PreviousID,
//		NextID:           sr.NextID,
//	})
//}
