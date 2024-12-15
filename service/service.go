package service

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context) {
	// Get the uploaded file
	file, err := c.FormFile("image")
	if err != nil {
		c.String(http.StatusBadRequest, "Unable to get file: %v", err)
		log.Printf("FormFile error: %v", err)
		return
	}

	// Open the file
	fileContent, err := file.Open()
	if err != nil {
		c.String(http.StatusInternalServerError, "Unable to open uploaded file: %v", err)
		log.Printf("Open file error: %v", err)
		return
	}
	defer fileContent.Close()

	// Use the /tmp directory for temporary storage (Render writes to /tmp)
	tmpFile, err := os.CreateTemp("/tmp", "input-*.png")
	if err != nil {
		c.String(http.StatusInternalServerError, "Unable to create temporary file: %v", err)
		log.Printf("CreateTemp error: %v", err)
		return
	}
	defer os.Remove(tmpFile.Name()) // Ensure temporary file is deleted
	defer tmpFile.Close()

	// Save the uploaded image to the temporary file
	_, err = io.Copy(tmpFile, fileContent)
	if err != nil {
		c.String(http.StatusInternalServerError, "Unable to save temporary file: %v", err)
		log.Printf("Save temporary file error: %v", err)
		return
	}

	// Process the image using the primitive command
	// Ensure that the path to the primitive binary is correct in your container
	cmd := exec.Command("/usr/local/bin/primitive", "-i", tmpFile.Name(), "-o", tmpFile.Name(), "-n", "100")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		c.String(http.StatusInternalServerError, "Error running primitive command: %v", err)
		log.Printf("Primitive command error: %v", err)
		return
	}

	// Open the processed file to send as a response
	processedFile, err := os.Open(tmpFile.Name())
	if err != nil {
		c.String(http.StatusInternalServerError, "Unable to open processed file: %v", err)
		log.Printf("Open processed file error: %v", err)
		return
	}
	defer processedFile.Close()

	// Set response headers and return the processed image as a blob
	c.Header("Content-Disposition", "attachment; filename=processed_image.png")
	c.Header("Content-Type", "image/png")
	c.Status(http.StatusOK)
	if _, err := io.Copy(c.Writer, processedFile); err != nil {
		log.Printf("Error writing processed file to response: %v", err)
	}
}
