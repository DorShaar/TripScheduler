package schedule

type ItemIdCreator struct {
	nextID int
}

func (creator *ItemIdCreator) NextId() (id int) {
	id = creator.nextID
	creator.nextID += 1
	return id
}
