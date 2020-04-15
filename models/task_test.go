package models

func (ms *ModelSuite) Test_Task() {
	task := Task{}
	ms.Equal("", task.Description)
	ms.Equal("", task.RequestedBy)
	ms.Equal("", task.ExecutedBy)
	ms.False(task.IsDone)
	ms.Zero(task.ExecutedAt)
}

func (ms *ModelSuite) Test_MemoryTaskStorage_Create() {
	ts := MemoryTaskStorage{}
	ts.Add(&Task{Description: "Do this"})
	ms.Len(ts, 1)
}

func (ms *ModelSuite) Test_MemoryTaskStorage_List() {
	ts := MemoryTaskStorage{}
	ms.Len(ts.List(), 0)

	ts.Add(&Task{Description: "Do this"})
	ms.Len(ts.List(), 1)
}

func (ms *ModelSuite) Test_MemoryTaskStorage_Clear() {
	ts := MemoryTaskStorage{}
	ts.Add(&Task{Description: "Do this"})
	ts.Clear()
	ms.Len(ts.List(), 0)
}
