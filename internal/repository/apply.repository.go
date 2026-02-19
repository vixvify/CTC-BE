package repository

type ApplyRepository interface {
	Do(fn func(team TeamRepository, quiz QuizRepository) error) error
}
