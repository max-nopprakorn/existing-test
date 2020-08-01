package hostel

import (
	"github.com/gin-gonic/gin"
)

// GetHostelsHandler is the handler for query all hostels.
func GetHostelsHandler(c *gin.Context) {
	hostels, err := getHostels()
	if err != nil {
		c.JSON(500, gin.H{
			"mesaage": "Something went wrong when trying to query all hostels.",
		})
		return
	}

	c.JSON(200, hostels)
}

// GetHostelByIDHandler is the handler for query a hostel by id.
func GetHostelByIDHandler(c *gin.Context) {
	hostelID := c.Param("hostelId")
	hostel, err := getHostelByID(hostelID)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Something went wrong when trying to query a hostel.",
		})
		return
	}
	if hostel == (&Hostel{}) {
		c.JSON(404, gin.H{
			"message": "Resource not found.",
		})
	}

	c.JSON(200, hostel)
}

// CreateHostelHandler is the handler for create a new hostel.
func CreateHostelHandler(c *gin.Context) {
	var hostel Hostel
	err := c.ShouldBindJSON(&hostel)
	if err != nil {
		c.JSON(422, gin.H{
			"message": "Invald json request.",
		})
		return
	}

	name := hostel.Name

	isExists := checkIfDuplicate(name)

	if isExists {
		c.JSON(409, gin.H{
			"message": "Hostel name already exists.",
		})
	}

	err = createHostel(hostel)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Something went wrong when creating hostel.",
		})
	}

	c.JSON(201, gin.H{
		"message": "Create hostel successfully.",
	})
}

func GetAvailableHostelsHandler(c *gin.Context) {
	availableHostels, err := getAvailableHostels()
	if err != nil {
		c.JSON(500, gin.H{
			"mesaage": "Something went wrong when trying to query available hostels.",
		})
		return
	}

	c.JSON(200, availableHostels)
}
