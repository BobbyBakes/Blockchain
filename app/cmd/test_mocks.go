package cmd

type mockMsgService struct {
	returnValue bool
}

func (mms *mockMsgService) CreatePhrase() bool { return mms.returnValue }
