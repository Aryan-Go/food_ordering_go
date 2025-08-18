package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
	"unicode"

	backend "github/aryan-go/food_ordering_go"
	"github/aryan-go/food_ordering_go/package/structures"

	"github.com/golang-jwt/jwt/v5"
)

func EmailVerification(email string) bool {
	emailRegex := `(?i)^(?:[a-z0-9!#$%&'*+/=?^_` + "`" + `{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_` + "`" + `{|}~-]+)*|"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*")@[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?(?:\.[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?)*$`
	re := regexp.MustCompile(emailRegex)
	if re.MatchString(email) {
		return true
	} else {
		return false
	}
}

func PasswordVerification(password string) bool {
	score := 0
	for _, char := range password {
		if unicode.IsUpper(char) {
			score = score + 1
		} else if unicode.IsLower(char) {
			continue
		} else if unicode.IsDigit(char) {
			score = score + 1
		} else {
			score = score + 1
		}
	}
	if score >= 5 {
		return true
	} else {
		return false
	}
}

func GetDotenvData() string {
	config, err := backend.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
		return ""
	}
	greeting := config.Secret_key
	return greeting
}

func CreateToken(email string, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"role":  role,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString([]byte(GetDotenvData())) // ! the key must be in the form of bytes
	if err != nil {
		log.Fatalf("Error in creating jwt: %s", err)
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			fmt.Println("Malformed token cus")
			var err structures.Error
			err.Code = http.StatusUnauthorized
			err.Message = "Malformed Token"
			json.NewEncoder(w).Encode(err)
			return
		}
		jwtToken := authHeader[1]
		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			config, err := backend.LoadConfig(".")
			if err != nil {
				log.Fatal(err)
			}
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.Secret_key), nil
		})
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx := context.WithValue(r.Context(), "props", claims)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}
		fmt.Println(err)
		var err2 structures.Error
		err2.Code = http.StatusUnauthorized
		err2.Message = "Unauthorized"
		json.NewEncoder(w).Encode(err2)
	})
}

func JWTAuthMiddlewareCustomer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			fmt.Println("Malformed token cus")
			var err structures.Error
			err.Code = http.StatusUnauthorized
			err.Message = "Malformed Token"
			json.NewEncoder(w).Encode(err)
			return
		}
		jwtToken := authHeader[1]
		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			config, err := backend.LoadConfig(".")
			if err != nil {
				log.Fatal(err)
			}
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.Secret_key), nil
		})
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["role"] != "customer" && claims["role"] != "admin" {
				var err structures.Error
				err.Code = http.StatusUnauthorized
				err.Message = "This is a protected route where only customer is allowed"
				json.NewEncoder(w).Encode(err)
				return
			} else if !ok {
				var err structures.Error
				err.Code = http.StatusBadRequest
				err.Message = "Some error in jwt"
				json.NewEncoder(w).Encode(err)
				return
			}
			ctx := context.WithValue(r.Context(), "props", claims)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}
		fmt.Println(err)
		var err2 structures.Error
		err2.Code = http.StatusUnauthorized
		err2.Message = "Unauthorized"
		json.NewEncoder(w).Encode(err2)
	})
}

func JWTAuthMiddlewareChef(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			fmt.Println("Malformed token chef")
			var err structures.Error
			err.Code = http.StatusUnauthorized
			err.Message = "Malformed Token"
			json.NewEncoder(w).Encode(err)
			return
		}
		jwtToken := authHeader[1]
		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			config, _ := backend.LoadConfig(".")
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.Secret_key), nil
		})
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["role"] != "chef" {
				var err structures.Error
				err.Code = http.StatusUnauthorized
				err.Message = "This is a protected route where only chef is allowed"
				json.NewEncoder(w).Encode(err)
				return
			}
			ctx := context.WithValue(r.Context(), "props", claims)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}
		fmt.Println(err)
		var err2 structures.Error
		err2.Code = http.StatusUnauthorized
		err2.Message = "Unauthorized"
		json.NewEncoder(w).Encode(err2)
	})
}

func JWTAuthMiddlewareAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			fmt.Println("Malformed token adm")
			var err structures.Error
			err.Code = http.StatusUnauthorized
			err.Message = "Malformed Token"
			json.NewEncoder(w).Encode(err)
			return
		}
		jwtToken := authHeader[1]
		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			config, _ := backend.LoadConfig(".")
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.Secret_key), nil
		})
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["role"] != "admin" {
				var err structures.Error
				err.Code = http.StatusUnauthorized
				err.Message = "This is a protected route where only admin is allowed"
				json.NewEncoder(w).Encode(err)
				return
			}
			if claims["email"] != "admin@gmail.com" {
				var err structures.Error
				err.Code = http.StatusUnauthorized
				err.Message = "This is a protected route and you cannot just put a role admin and enter this route"
				json.NewEncoder(w).Encode(err)
				return
			}
			ctx := context.WithValue(r.Context(), "props", claims)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}
		fmt.Println(err)
		var err2 structures.Error
		err2.Code = http.StatusUnauthorized
		err2.Message = "Unauthorized"
		json.NewEncoder(w).Encode(err2)
	})
}
func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
