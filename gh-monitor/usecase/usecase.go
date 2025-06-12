package usecase

type GHMonitorUsecase struct{
	gitHubAPI gitHubAPI
}

func NewGHMonitorUsecase(gitHubAPI gitHubAPI) *GHMonitorUsecase {
	return &GHMonitorUsecase{
		gitHubAPI: gitHubAPI,
	}
}

func (uc *GHMonitorUsecase) InitFetching() {
	
}