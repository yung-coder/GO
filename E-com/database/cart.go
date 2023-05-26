package database

import ()

var (
	ErrCartFindProduct    = errors.New("can't find product")
	ErrCartDecodeProducts = errors.New("cant find the producy")
	ErrUserIdIsNotValid   = errors.New("this user is not valid")
	ErrCantUpdateUser     = errors.New("cannot add this product to cart")
	ErrCantRemoveItemCart = errors.New("cannot remove this item from cart")
	ErrCantGetItem        = errors.New("was unable to get item from the cart")
	ErrCantBuyCartItem    = errors.New("cannot update the purchase")
)

func AddProductToCart() {

}

func RemovecartItem() {

}

func BuyItemFromCart() {

}

func InstantBuyer() {

}
