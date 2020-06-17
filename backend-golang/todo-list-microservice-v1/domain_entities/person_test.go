package domain_test

import (
	"testing"

	domain "github.com/dotm/clean-architecture/backend-golang/todo-list-microservice-v1/domain_entities"
	"github.com/dotm/clean-architecture/backend-golang/utilities/assert"
)

func TestAddEntryToList(t *testing.T) {
	t.Run("positive case", func(t *testing.T) {
		//given
		I := domain.Person{Name: "Yoshua"}
		workList := I.CreateNewList("work")

		//when
		err := I.AddEntryToList("check email", &workList)

		//then
		assert.True(t, len(workList.Entries) == 1)
		assert.NoError(t, err)
	})

	t.Run("failed when we exceed maximum allowed entry length", func(t *testing.T) {
		//given
		I := domain.Person{Name: "Yoshua"}
		workList := I.CreateNewList("work")

		//when
		err := I.AddEntryToList("a very long string that we use to trigger a business logic requirement", &workList)

		//then
		assert.True(t, len(workList.Entries) == 0)
		assert.ErrorRaised(t, err)
	})
}
