package vangogh_values

import (
	"encoding/json"
	"fmt"
	"github.com/arelate/gog_types"
	"github.com/arelate/vangogh_types"
	"github.com/arelate/vangogh_urls"
	"github.com/boggydigital/kvas"
)

type ValueReader struct {
	productType vangogh_types.ProductType
	mediaType   gog_types.Media
	valueSet    *kvas.ValueSet
}

func NewReader(pt vangogh_types.ProductType, mt gog_types.Media) (*ValueReader, error) {
	dst, err := vangogh_urls.LocalProductsDir(pt, mt)
	if err != nil {
		return nil, err
	}

	vs, err := kvas.NewJsonLocal(dst)
	if err != nil {
		return nil, err
	}

	vr := &ValueReader{
		productType: pt,
		mediaType:   mt,
		valueSet:    vs,
	}

	return vr, nil
}

func (vr *ValueReader) readValue(id string, pt vangogh_types.ProductType, val interface{}) error {
	if vr.productType != pt {
		return fmt.Errorf("vangogh_types: %s value reader doesn't support %s", vr.productType, pt)
	}

	spReadCloser, err := vr.valueSet.Get(id)
	if err != nil {
		return err
	}

	if spReadCloser == nil {
		return nil
	}

	defer spReadCloser.Close()

	if err := json.NewDecoder(spReadCloser).Decode(val); err != nil {
		return err
	}

	return nil
}

func (vr *ValueReader) All() []string {
	return vr.valueSet.All()
}

func (vr *ValueReader) Contains(id string) bool {
	return vr.valueSet.Contains(id)
}

func (vr *ValueReader) CreatedAfter(timestamp int64) []string {
	return vr.valueSet.CreatedAfter(timestamp)
}

func (vr *ValueReader) ModifiedAfter(timestamp int64) []string {
	return vr.valueSet.ModifiedAfter(timestamp)
}

func (vr *ValueReader) StoreProduct(id string) (storeProduct *gog_types.StoreProduct, err error) {
	err = vr.readValue(id, vangogh_types.StoreProducts, &storeProduct)
	return storeProduct, err
}

func (vr *ValueReader) AccountProduct(id string) (accountProduct *gog_types.AccountProduct, err error) {
	err = vr.readValue(id, vangogh_types.AccountProducts, &accountProduct)
	return accountProduct, err
}

func (vr *ValueReader) WishlistProduct(id string) (wishlistProduct *gog_types.StoreProduct, err error) {
	err = vr.readValue(id, vangogh_types.WishlistProducts, &wishlistProduct)
	return wishlistProduct, err
}

func (vr *ValueReader) Details(id string) (details *gog_types.Details, err error) {
	err = vr.readValue(id, vangogh_types.Details, &details)
	return details, err
}

func (vr *ValueReader) ApiProductV1(id string) (apiProductV1 *gog_types.ApiProductV1, err error) {
	err = vr.readValue(id, vangogh_types.ApiProductsV1, &apiProductV1)
	return apiProductV1, err
}

func (vr *ValueReader) ApiProductV2(id string) (apiProductV2 *gog_types.ApiProductV2, err error) {
	err = vr.readValue(id, vangogh_types.ApiProductsV2, &apiProductV2)
	return apiProductV2, err
}

func (vr *ValueReader) StorePage(page string) (storePage *gog_types.StoreProductsPage, err error) {
	err = vr.readValue(page, vangogh_types.StorePage, &storePage)
	return storePage, err
}

func (vr *ValueReader) AccountStorePage(page string) (accountStorePage *gog_types.AccountProductsPage, err error) {
	err = vr.readValue(page, vangogh_types.AccountPage, &accountStorePage)
	return accountStorePage, err
}

func (vr *ValueReader) WishlistPage(page string) (wishlistPage *gog_types.WishlistPage, err error) {
	err = vr.readValue(page, vangogh_types.WishlistPage, &wishlistPage)
	return wishlistPage, err
}

func (vr *ValueReader) ProductType(key string, pt vangogh_types.ProductType) (interface{}, error) {
	switch pt {
	case vangogh_types.StoreProducts:
		return vr.StoreProduct(key)
	case vangogh_types.AccountProducts:
		return vr.AccountProduct(key)
	case vangogh_types.WishlistProducts:
		return vr.WishlistProduct(key)
	case vangogh_types.Details:
		return vr.Details(key)
	case vangogh_types.ApiProductsV1:
		return vr.ApiProductV1(key)
	case vangogh_types.ApiProductsV2:
		return vr.ApiProductV2(key)
	case vangogh_types.StorePage:
		return vr.StorePage(key)
	case vangogh_types.AccountPage:
		return vr.AccountStorePage(key)
	case vangogh_types.WishlistPage:
		return vr.WishlistPage(key)
	default:
		return nil, fmt.Errorf("vangogh_values: cannot create %s value", pt)
	}
}
