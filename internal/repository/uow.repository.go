package repository

type UnitOfWork interface {
	Do(fn func(team TeamRepository, quiz QuizRepository) error) error
}
