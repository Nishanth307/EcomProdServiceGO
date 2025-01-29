package handlers // Incoming http request from the client and return responses
import (
	// Go Internal Packages
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	// Local Packages
	models "products/models"

	// External Packages
	"github.com/gorilla/mux"
)

type ProductService interface {
	GetProductById(ctx context.Context, id int) (*models.Product, error)
	GetAllProducts(ctx context.Context) ([]models.Product, error)
	CreateProduct(ctx context.Context, product models.Product) error
	UpdateProduct(ctx context.Context, id int, product models.Product) error
	DeleteProductById(ctx context.Context, id int) error
}
type ProductHandler struct {
	service ProductService
}

func NewProductHandler(service ProductService) *ProductHandler { //Constructor
	return &ProductHandler{service: service} //
}

func (h *ProductHandler) GetProductById(w http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)           // get the params from the request
	id, err := strconv.Atoi(params["id"]) // convert the id to integer
	if err != nil {
		fmt.Println("Error converting to integer")
		return
	}
	product, err := h.service.GetProductById(request.Context(), id) // get the product from the service
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // return the error if there is any
		return
	}
	// json.NewEncoder(w).Encode(product) // encode the product into json and return it
	if err := json.NewEncoder(w).Encode(product); err != nil {
		http.Error(w, "Failed to encode product to JSON", http.StatusInternalServerError)
		return
	  }
}

func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, request *http.Request) {
	products, err := h.service.GetAllProducts(request.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := h.service.CreateProduct(r.Context(), product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	// w.Write([]byte("Product created successfully"))
	if _, err := w.Write([]byte("Product created successfully")); err != nil {
		log.Println("Failed to write response:", err)
	  }
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var product models.Product
	if err := json.NewDecoder(request.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.UpdateProduct(request.Context(), id, product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	// w.Write([]byte("Product updated successfully"))
	if _, err := w.Write([]byte("Product updated successfully")); err != nil {
		log.Println("Failed to write response:", err)
	  }
}

func (h *ProductHandler) DeleteProductById(w http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.DeleteProductById(request.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	// w.Write([]byte("Product deleted successfully"))
	if _, err := w.Write([]byte("Product deleted successfully")); err != nil {
		log.Println("Failed to write response:", err)
	  }
}
