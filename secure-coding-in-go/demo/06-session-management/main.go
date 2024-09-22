package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	// JWT is not in the native Go packages
	"github.com/dgrijalva/jwt-go"
)

/*
Using the JWT (json web token) for enabling stateless session management and easy sharing of session across different systems or systems

In traditioinal session management, the session identifier would be a unique string stored on the server and referenced via a cookie (like session_id=abc123). In JWT-based systems, the JWT itself replaces the session identifier, and it is stored by the client and sent back in the Authorization header for each request. 
*/


type Key int

const MyKey Key = 0

// JWT schema of the data it will store
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Create a JWT and put in the clients cookie
func setToken(w http.ResponseWriter, r *http.Request) {
	// 30m Expiration for non-sensitive application - OWASP
	expireToken := time.Now().Add(time.Minute * 30).Unix()
	expireCookie := time.Now().Add(time.Minute * 30)

	// token Claims
	claims := Claims{
		"TestUser",
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "localhost:9000",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte("secret"))

	// Set Cookie parameters
	cookie := http.Cookie{
		Name:     "Auth",
		Value:    signedToken,  // the signed token is used in the cookie
		Expires:  expireCookie, // 30 mins
		HttpOnly: true,
		Path:     "/",
		Domain:   "127.0.0.1",
		Secure:   true,
	}
	// respne
	http.SetCookie(w, &cookie) // Comment: By using http.SetCookie(), the Go standard library appends the cookie to the header, ensuring that multiple cookies can be set withot overwriting each other. There can be multgiple Set-Cookie headers in a single HTTP respons (one for each cookie). If you use http.Header().Set("Set-Cookie", "..."), this overwrites any existing Set-Cookie headers, meaning only the last one will be set.
	http.Redirect(w, r, "/profile", http.StatusTemporaryRedirect)

}

// Middleware to protect private pages
func validate(page http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("Auth")
		if err != nil {
			w.Header().Set("Content-Type", "text/html")
			fmt.Fprintf(w, "Unauthorized - Please login <br>")
			fmt.Fprintf(w, "<a href=\"login\"> Login </a>") // This creates a hyperlink labeled Login, which points to the relative URL /login. Browser sends an HTTP GET request to the /login endpoint.

			// When the user clicks the "Login" link, the browser automatically changes the URL to /login and sends an HTTP request to your server.
			return
		}

		token, err := jwt.ParseWithClaims(cookie.Value, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte("secret"), nil
		})

		if err != nil {
			w.Header().Set("Content-Type", "text/html")
			fmt.Fprint(w, "Unauthorized - Please login <br>")
			fmt.Fprintf(w, "<a href=\"login\"> Login </a>")
			return
		}
		
		// token.Valid ensures that token is properly signed and has not expired. 
		// claims refers to the data or payload within the JWT (usually a Claims struct, which holds user info or other data). 
		// context.WithValue() stores the claims (user data, permissions, etc) in the request's context under the MyKey key. This allows other parts of the request lifecycle to access the claims without needing to re-parse the JWT token from the request headers. 

		// Using the context allows you to pass the claims without having to re-validate the token in every handler. 

		// claims := r.Context().Value(MyKey).(*Claims) to retrieve the claims from the context.

		// Parsing and verifying a JWT token multiple times (in each handler) is inefficient. By storing claims in the context, they are available across all handlers after a single verification.



		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			ctx := context.WithValue(r.Context(), MyKey, *claims)
			
			// Only viewable if the client has a valid token.
			page(w, r.WithContext(ctx))
		} else {
			w.Header().Set("Content-Type", "text/html")
			fmt.Fprint(w, "Unauthorized - Please login <br>")
			fmt.Fprintf(w, "<a href=\"login\"> Login </a>")
			return
		}
	}
}

// Only viewable if the client has a valid token
func protectedProfile(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(MyKey).(Claims)
	if !ok {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, "Unauthorized - Please login <br>")
		fmt.Fprintf(w, "<a href=\"login\"> Login </a>")

		return
	}

	w.Header().Set("Content-Type", "text/html")
	// We put claims in the context to extract the username. 
	fmt.Fprintf(w, "Hello %s<br>", claims.Username)
	fmt.Fprintf(w, "<a href=\"logout\"> Logout </a>")
}

// Delete the cookie
func logout(w http.ResponseWriter, r *http.Request) {
	deleteCookie := http.Cookie{
		// Since the cookie name is the same 'Auth', the original cookie is deleted.
		Name:    "Auth",
		Value:   "none",
		Expires: time.Now(),
	}

	http.SetCookie(w, &deleteCookie)
	http.Redirect(w, r, "/profile", http.StatusTemporaryRedirect)
}

func main() {
	http.HandleFunc("/", validate(protectedProfile))
	http.HandleFunc("/login", setToken)
	http.HandleFunc("/profile", validate(protectedProfile))
	http.HandleFunc("/logout", validate(logout))
	err := http.ListenAndServeTLS(":443", "cert/certificate.crt", "cert/private.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
