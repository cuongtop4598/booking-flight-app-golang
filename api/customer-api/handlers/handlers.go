package handlers

import (
	"net/http"

	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/api/customer-api/requests"
	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/api/customer-api/responses"
	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/pb"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type CustomerHandler interface {
	CreateCustom(c *gin.Context)
	//FindCustomerById(c *gin.Context)
	//FindCustomerByPhone(c *gin.Context)
	//FindCustomerByEmail(c *gin.Context)
	//UpdateCustomer(c *gin.Context)
	//ChangePassword(c *gin.Context)
	BookingHistory(c *gin.Context)
}

type customerHandler struct {
	customerClient pb.CustormerClient
}

func NewCustomerHandler(customerClient pb.CustormerClient) CustomerHandler {
	return &customerHandler{
		customerClient: customerClient,
	}
}

func (h *customerHandler) CreateCustom(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

//FindCustomerById(c *gin.Context)
//FindCustomerByPhone(c *gin.Context)
//FindCustomerByEmail(c *gin.Context)
//UpdateCustomer(c *gin.Context)
//ChangePassword(c *gin.Context)
func (h *customerHandler) BookingHistory(c *gin.Context) {
	req := requests.CustomerIdRequest{}

	if err := c.ShouldBind(&req); err != nil {
		if validateErrors, ok := err.(validator.ValidationErrors); ok {
			errMessages := make([]string, 0)
			for _, v := range validateErrors {
				errMessages = append(errMessages, v.Kind().String())
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  errMessages,
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})
		return
	}
	customerId := req.Id
	histories, err := h.customerClient.BookingHistory(c.Request.Context(), &pb.CustomerAuthen{Id: customerId})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
	}
	dto := []*responses.BookingHistoryReponse{}
	for _, v := range histories.Historys {
		dto = append(dto, &responses.BookingHistoryReponse{
			BookingCode:   v.BookingCode,
			BookingStatus: v.Status,
			BookedDate:    v.BookingDate,
			Flight:        v.Flight,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"payload": dto,
	})
}
