package mock

import "github.com/YuChaoGithub/YARC/backend/app/models"

// SubredditModel is a mock model for subreddit.
type SubredditModel struct{}

// List returns ["meirl", "dankmeme"].
func (m *SubredditModel) List() []string {
	return []string{"meirl", "dankmeme"}
}

// Insert always returns nil.
func (m *SubredditModel) Insert(name, description string) error {
	return nil
}

// Get only returns a result when name is "radiohead".
func (m *SubredditModel) Get(name string) (models.SubredditInfo, error) {
	if name == "radiohead" {
		return models.SubredditInfo{
			Name:        "radiohead",
			Members:     3,
			Description: "Dedicated to all human beings.",
		}, nil
	}
	return models.SubredditInfo{}, models.ErrSubredditNotExist
}

// IncrVisitCount does nothing.
func (m *SubredditModel) IncrVisitCount(name string) {
}

// GetTrending returns a result if limit=2.
func (m *SubredditModel) GetTrending(limit int) ([]models.SubredditInfo, error) {
	return []models.SubredditInfo{
		{
			Name:        "radiohead",
			Members:     0,
			Description: "Dedicated to all human beings.",
		},
		{
			Name:        "underrated",
			Members:     4,
			Description: "Let down and hanging around.",
		},
	}, nil
}
