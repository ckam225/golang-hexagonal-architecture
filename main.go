package cleanarchhexagonal

//go:generate mockgen -destination=./internal/db/mocks/mock.go "clean-arch-hex/internal/db" Database
////go:generate mockgen -source=./internal/domain/repository/user.go -destination=./internal/domain/repository/mocks/mock_user.go

//go:generate mockgen -destination=./internal/domain/repository/mocks/mock.go "clean-arch-hex/internal/domain/repository" PostRepository,UserRepository
