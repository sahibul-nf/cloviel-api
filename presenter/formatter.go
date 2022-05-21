package presenter

type PresenterFormatter struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	AvatarURL        string `json:"avatar_url"`
}

func FormatPresenter(presenter Presenter) PresenterFormatter {

	formatter := PresenterFormatter{}
	formatter.ID = presenter.ID
	formatter.Name = presenter.Name
	formatter.ShortDescription = presenter.ShortDescription
	formatter.AvatarURL = presenter.AvatarURL

	return formatter
}

func FormatPresenters(presenters []Presenter) []PresenterFormatter {

	if len(presenters) == 0 {
		return []PresenterFormatter{}
	}

	var formatted []PresenterFormatter
	for _, presenter := range presenters {
		formatted = append(formatted, FormatPresenter(presenter))
	}

	return formatted
}
