package usecase

// import (
// 	"reflect"
// 	"testing"

// 	"github.com/hieuphq/califit/src/domain/model"
// 	"github.com/hieuphq/califit/src/interfaces/repository"
// 	"github.com/hieuphq/califit/src/interfaces/repository/service"
// )

// func Test_userUsecase_ListUser(t *testing.T) {
// 	type fields struct {
// 		repo     repository.DBRepo
// 		userRepo repository.UserRepository
// 		service  *service.UserService
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		want    []model.User
// 		wantErr bool
// 	}{
// 		{
// 			name: "get all user NOT Ân",
// 			fields: fields{
// 				userRepo: &repository.UserRepositoryMock{
// 					FindAllFunc: func(repo repository.DBRepo, param repository.QueryParam) ([]model.User, error) {
// 						return []model.User{
// 							model.User{
// 								Name: "Ân",
// 							},
// 						}, nil
// 					},
// 				},
// 			},
// 			want: []model.User{
// 				model.User{
// 					Name: "Ân",
// 				},
// 			},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			u := &userUsecase{
// 				repo:     tt.fields.repo,
// 				userRepo: tt.fields.userRepo,
// 				service:  tt.fields.service,
// 			}
// 			got, err := u.ListUser()
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("userUsecase.ListUser() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("userUsecase.ListUser() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
