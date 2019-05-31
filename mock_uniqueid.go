package boltlogrus

type mockUniqueID struct{}

//NewMockUniqueID will construct fake object
//to generate uniqueid
func NewMockUniqueID() UniqueID {
	return &mockUniqueID{}
}

func (m *mockUniqueID) GenerateID() (string, error) {
	return "UNIK", nil
}
