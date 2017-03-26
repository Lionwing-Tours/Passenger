package jwt

import (
	//"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("secret")

type Claims struct {
	Admin    bool   `json:"isAdmin`
	Username string `json:"username"`
	jwt.StandardClaims
}

func SetToken(w http.ResponseWriter, r *http.Request) {
	// Expires the token and cookie in 1 hour
	expireToken := time.Now().Add(time.Hour * 1).Unix()
	issuedAt := time.Now().Unix()

	// Manually assign the claims but in production, insert values from a database
	claims := Claims{
		false,
		"superman",
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			IssuedAt:  issuedAt,
			Issuer:    "localhost:3000",
		},
	}

	// Create the token using your claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Signs the token with a secret.
	signedToken, _ := token.SignedString(mySigningKey)

	log.Println(token.Claims)
	log.Println(signedToken)
	w.Write([]byte(signedToken))
}

func Show(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")

	if len(authHeader) < 10 || !strings.HasPrefix(authHeader, "Bearer") || len(strings.Split(authHeader, " ")) != 2 {
		//w.Write(http.StatusUnauthorized)
		log.Println("unauth")
		return
	}

	tokenString := strings.Split(authHeader, " ")[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	//	if token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token)); err == nil {
	//		claims := token.Claims.(*Claims)
	//		log.Println("Token for user %v expires %v", claims.User, claims.StandardClaims.ExpiresAt)
	//	}

	if err != nil {
		log.Println("not found")
		http.NotFound(w, r)
		return
	}
	log.Println("Validatd token claims: ", token.Claims)
}

//func validate(w http.ResponseWriter, r *http.Request) {
//	authHeader := r.Header.Get("Authorization")
//	authClaims := r.Header.Get("Claims")
//	//authClaims := Claims{}
//	//passengerId := c.PostForm("passengerId")

//	if len(authHeader) < 10 || !strings.HasPrefix(authHeader, "Bearer") || len(strings.Split(authHeader, " ")) != 2 {
//		w.Write(http.StatusUnauthorized)
//		return
//	}

//	if len(authClaims) == 0 || !strings.HasPrefix(authClaims, "Claims") || len(strings.Split(authClaims, " ")) != 2 {
//		w.Write(http.StatusUnauthorized)
//		return
//	}

//	tokenString := strings.Split(authHeader, " ")[1]
//	token, err := jwt.Parse(tokenString, func(token *jwt.Token) ([]byte, error) {
//		return mySigningKey, nil
//	})
//	if err != nil {
//		http.NotFound(w, r)
//		return
//	}

//	//	if token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, keyLookupFunc); err == nil {
//	//		claims := token.Claims.(*MyCustomClaims)
//	//		fmt.Printf("Token for user %v expires %v", claims.User, claims.StandardClaims.ExpiresAt)
//	//	}

//	if err != nil {
//		http.NotFound(res, req)
//		return
//	}

//	// Grab the tokens claims and pass it into the original request
//	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
//		ctx := context.WithValue(req.Context(), MyKey, *claims)
//		page(res, req.WithContext(ctx))
//	} else {
//		http.NotFound(res, req)
//		return
//	}

//}

//func ValidateErrors(err error, w http.ResponseWriter) {
//	if err != nil {
//		w.Write(http.StatusUnauthorized)
//		return
//	}
//}
