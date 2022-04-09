package databases

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"testing"
)

type FakeAlbum struct {
	ID     int64   `faker:"oneof: 1, 2, 3"`
	Title  string  `faker:"title"`
	Artist string  `faker:"name"`
	Price  float32 `faker:"oneof: 4.95, 9.99, 31997.97"`
}

func TestInsertAlbum(t *testing.T) {
	a := FakeAlbum{}
	err := faker.FakeData(&a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", a)
}
