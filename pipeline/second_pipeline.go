package pipeline

import (
	"math/rand"

	"github.com/jaswdr/faker"
	"github.com/ramadhanalfarisi/go-concurrency-pipeline/model"
)

func CreatePerson(count int) chan model.Person {
	chanOut := make(chan model.Person)
	fake := faker.New().Person()
	go func(person chan model.Person) {
		for i := 0; i < count; i++ {
			name := fake.Name()
			age := rand.Intn(100)
			contact := fake.Contact()
			gender := fake.Gender()
			personObject := model.Person{Name: name, Age: age, Contact: contact.Phone, Gender: gender}
			chanOut <- personObject
		}
		close(chanOut)
	}(chanOut)
	return chanOut
}

