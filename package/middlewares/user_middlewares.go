package middlewares

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"time"
	"unicode"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func Email_verification(email string) bool {
	emailRegex := `(?i)^(?:[a-z0-9!#$%&'*+/=?^_` + "`" + `{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_` + "`" + `{|}~-]+)*|"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*")@[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?(?:\.[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?)*$`
	re := regexp.MustCompile(emailRegex)
	if re.MatchString(email) {
		return true
	} else {
		return false
	}
}

func Password_verification(password string) bool {
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

func Get_dotenv_data() string {
	err := godotenv.Load("/Users/aryangoyal/Desktop/golang/sdsProject/backend/.env") // ! try to give absolute route
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
		return err.Error()
	} else {
		greeting := os.Getenv("secret_key")
		return greeting
	}
}

func Create_token(email string, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"role":  role,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString([]byte(Get_dotenv_data())) // ! the key must be in the form of bytes
	if err != nil {
		log.Fatalf("Error in creating jwt: %s", err)
		return "", err
	}
	return tokenString, nil
}

func Verify_token(tokenString string) (bool, string, string) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(Get_dotenv_data()), nil
	})

	if err != nil {
		fmt.Printf("Error in verifying jwt: %s", err)
		return false, "", ""
	}

	if !token.Valid {
		return false, "", ""
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return true, claims["email"].(string), claims["role"].(string)
	}
	return false, "", ""
}
