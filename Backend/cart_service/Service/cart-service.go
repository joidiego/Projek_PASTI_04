package service

import (
	"cart_service/dto"
	"cart_service/entity"
	"cart_service/repository"
	"log"

	"github.com/mashingan/smapping"
)

// CartService is a contract about something that this service can do
type CartService interface {
	Insert(b dto.CartCreateDTO) entity.Cart
	Update(b dto.CartUpdateDTO) entity.Cart
	Delete(b entity.Cart)
	All(userId uint64) []entity.Cart
}

type cartService struct {
	cartRepository repository.CartRepository
}

// NewCartService creates a new instance of CartService
func NewCartService(cartRepository repository.CartRepository) CartService {
	return &cartService{
		cartRepository: cartRepository,
	}
}

func (service *cartService) All(userID uint64) []entity.Cart {
	return service.cartRepository.All(userID)
}

func (service *cartService) Insert(b dto.CartCreateDTO) entity.Cart {
	cart := entity.Cart{}
	err := smapping.FillStruct(&cart, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}

	res := service.cartRepository.InsertCart(cart)
	return res
}

func (service *cartService) Update(b dto.CartUpdateDTO) entity.Cart {
	cart := entity.Cart{}
	err := smapping.FillStruct(&cart, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}

	res := service.cartRepository.UpdateCart(cart)
	return res
}

func (service *cartService) Delete(b entity.Cart) {
	service.cartRepository.DeleteCart(b)
}
