package main 

import (
	"fmt"

	"github.com/google/uuid"
)

type Guid struct {
    guid []byte
}

func generateUserID() {
	id := uuid.New()
	fmt.Println(id.String())
	return 
}