package main

import (
	"fmt"
	"github.com/arelate/gog_types"
	"github.com/arelate/vangogh_types"
	"github.com/boggydigital/vangogh_values"
	"log"
)

func main() {
	vrStoreProducts, err := vangogh_values.NewReader(vangogh_types.StoreProducts, gog_types.Game)
	if err != nil {
		log.Fatal(err)
	}

	sp, err := vrStoreProducts.StoreProduct("1073954123")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sp.Title)
}
