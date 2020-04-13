package graph

import (
	"cloud.google.com/go/firestore"
	"github.com/tryy3/go-cloudinary"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	Database *firestore.Client
	CloudinaryService *cloudinary.Service
}
