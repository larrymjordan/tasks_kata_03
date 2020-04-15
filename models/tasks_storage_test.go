package models

func (ms *ModelSuite) Test_DatabaseTaskStorage_List() {
	ts := DatabaseTaskStorage{Tx: ms.DB}
	ms.Len(ts.List(), 0)

	ms.NoError(ms.DB.Create(&Task{Description: "Do this!"}))
	ms.Len(ts.List(), 1)
}

func (ms *ModelSuite) Test_DatabaseTaskStorage_Create() {
	ts := DatabaseTaskStorage{Tx: ms.DB}
	ts.Add(&Task{Description: "Do this!"})
	ms.Len(ts.List(), 1)
}

func (ms *ModelSuite) Test_DatabaseTaskStorage_Clear() {
	ts := DatabaseTaskStorage{Tx: ms.DB}
	ts.Add(&Task{Description: "Do this!"})
	ms.Len(ts.List(), 1)

	ts.Clear()
	ms.Len(ts.List(), 0)
}
