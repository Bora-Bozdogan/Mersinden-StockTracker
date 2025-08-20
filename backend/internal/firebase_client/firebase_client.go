package firebase_client

import (
	"strings"

	"firebase.google.com/go/v4/auth"
	"github.com/gofiber/fiber/v2"
)

type firebaseClient struct {
	authClient *auth.Client
}

type FirebaseClientInterface interface {
	NewFirebaseClient(authClient *auth.Client) *firebaseClient
	AuthenticateToken(c *fiber.Ctx) error
}

func NewFirebaseClient(authClient *auth.Client) *firebaseClient {
	return &firebaseClient{authClient: authClient}
}

func (f *firebaseClient) AuthenticateToken(c *fiber.Ctx) error {
	//skip preflight checks
	if c.Method() == fiber.MethodOptions {
		return c.Next()
	}
	//get token from context
	token := c.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")

	//verify it with sdk function
	authToken, err := f.authClient.VerifyIDToken(c.UserContext(), token)
	
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "Unauthorized for request",
			"details": err.Error(),
		})
	}

	//add to local context
	c.Locals("uid", authToken.UID)

	//call next header for middleware functions
	return c.Next()
}
