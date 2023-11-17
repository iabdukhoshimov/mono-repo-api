package alif

type OnboardingHandlers interface {
}

type ApplicationHandlers interface {
}

type Handlers interface {
	OnboardingHandlers
	ApplicationHandlers
}

