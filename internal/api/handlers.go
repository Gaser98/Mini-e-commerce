package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"db-design-project/internal/db"
)

func RegisterRoutes(r *gin.Engine, q *db.Queries) {

	// ---------- LOGIN ----------
	r.POST("/login", func(c *gin.Context) {
		var req struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := c.BindJSON(&req); err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		user, err := q.GetUserByEmail(c, req.Email)
		if err != nil {
			c.Status(http.StatusUnauthorized)
			return
		}
		/*debug logs were hashed after finishing the debuging 
		fmt.Println("DEBUG DB HASH:", user.PasswordHash)
		fmt.Println("DEBUG INPUT PASSWORD:", req.Password)*/

		if err := bcrypt.CompareHashAndPassword(
			[]byte(user.PasswordHash),
			[]byte(req.Password),
		); err != nil {
			c.Status(http.StatusUnauthorized)
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": user.ID,
			"exp": time.Now().Add(15 * time.Minute).Unix(),
		})

		signed, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"access_token": signed,
		})
	})

	// ---------- AUTH GROUP ----------
	auth := r.Group("/")
	auth.Use(AuthMiddleware())

	// ---------- CURRENT USER ----------
	auth.GET("/users/me", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"user_id": c.GetInt("user_id"),
		})
	})

	// ---------- PRODUCTS ----------
	r.GET("/products", func(c *gin.Context) {
		products, err := q.ListAvailableProducts(c, db.ListAvailableProductsParams{
			Limit:  10,
			Offset: 0,
		})
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, products)
	})

	r.GET("/products/:id", func(c *gin.Context) {
		id := mustAtoi(c.Param("id"))

		product, err := q.GetProductByID(c, int32(id))
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		c.JSON(http.StatusOK, product)
	})

	// ---------- CREATE ORDER ----------
	auth.POST("/orders", func(c *gin.Context) {
		var req struct {
			Total float64 `json:"total"`
		}

		if err := c.BindJSON(&req); err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		order, err := q.CreateOrder(c, db.CreateOrderParams{
			UserID: int32(c.GetInt("user_id")),
			Total: sql.NullString{
				String: fmt.Sprintf("%.2f", req.Total),
				Valid:  true,
			},
		})
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusCreated, order)
	})

	// ---------- LIST ORDERS ----------
	auth.GET("/orders", func(c *gin.Context) {
		orders, err := q.ListOrdersByUser(c, int32(c.GetInt("user_id")))
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, orders)
	})
}

// ---------- HELPERS ----------
func mustAtoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
