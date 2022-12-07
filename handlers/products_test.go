package handlers

import (
	"fmt"
	"temp/schema"
	"testing"
)

type preget struct {
}


func getMock(pro schema.Product) bool {
	return true
}
func (p preget) get(pro schema.Product) bool {
	return getMock(pro)
}

// func TestCreateproduct(t *testing.T) {
// 	controller := gomock.NewController(t)
// 	defer controller.Finish()
// 	MockInterface := NewMockprods(controller)
// 	MockInterface.EXPECT().get(gomock.Any()).Times(1)

// 	// empty{}.p.Createproduct()
// }

// func TestCreateproduct(t *testing.T) {
// 	controller := gomock.NewController(t)
// 	defer controller.Finish()
// 	MockInterface := NewMockprods(controller)
// 	MockInterface.EXPECT().get(gomock.Any()).Times(1)

// 	// empty{}.p.Createproduct()
// }

var regpro prods

func TestCreateproduct(t *testing.T) {
	regpro = preget{}

	found := regpro.get(schema.Product{Name: "gst"})
	if found {
		fmt.Println("done")
	}
	// empty{}.p.Createproduct()
}
