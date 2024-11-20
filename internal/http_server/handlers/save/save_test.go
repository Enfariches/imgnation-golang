package save

import (
	"img/internal/http_server/handlers/save/mocks"
	"img/internal/lib/random"
	"mime/multipart"
	"os"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestSaveImage(t *testing.T) {
	cases := []struct {
		Name string
		File func(path string) multipart.File
	}{
		{
			Name: "Happy Path",
			File: func(path string) multipart.File {
				file, _ := os.Open(path)
				var mpltFile multipart.File = file
				return mpltFile
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			mockSaver := mocks.NewS3SaverImage(t) //Создание мока по интерфейсу хендлера

			file := tc.File("../testing_files/picture3.png")

			mockSaver.On("Save", file, mock.AnythingOfType("string")).Return(nil) //Что ожидает Мок, когда мы его вызовим
			err := mockSaver.Save(file, random.RandStringByte(5))                 // Тот самый вызов мока, возвращает то, что мы указали выше
			if err != nil {
				t.Error(err)
			}
				//to be continued:)
		})
	}
}
