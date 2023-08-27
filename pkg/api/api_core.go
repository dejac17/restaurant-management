package api

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var ginLambda *ginadapter.GinLambda
var foodCollection *mongo.Collection = OpenCollection(Client, "food")

func SetApiRouter() *gin.Engine {
	router := gin.Default()
	router.Use(authentication())

	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", GetUsers())
		userRoutes.GET("/:userId", GetUser())
		userRoutes.POST("/signUp", SignUp())
		userRoutes.POST("/login", Login())
	}

	foodRoutes := router.Group("/foods")
	{
		foodRoutes.GET("/", GetFoods())
		foodRoutes.GET("/:foodId", GetFood())
		foodRoutes.POST("/", CreateFood())
		foodRoutes.PATCH("/:foodId", UpdateFood())
	}

	menuRoutes := router.Group("/menus")
	{
		menuRoutes.GET("/", GetMenus())
		menuRoutes.GET("/:menuId", GetMenu())
		menuRoutes.POST("/", CreateMenu())
		menuRoutes.PATCH("/:menuId", UpdateMenu())
	}

	tableRoutes := router.Group("/tables")
	{
		tableRoutes.GET("/", GetTables())
		tableRoutes.GET("/:tableId", GetTable())
		tableRoutes.POST("/", CreateTable())
		tableRoutes.PATCH("/:tableId", UpdateTable())
	}

	orderRoutes := router.Group("/orders")
	{
		orderRoutes.GET("/", GetOrders())
		orderRoutes.GET("/:orderId", GetOrder())
		orderRoutes.POST("/", CreateOrder())
		orderRoutes.PATCH("/:orderId", UpdateOrder())
	}

	orderItemRoutes := router.Group("/orderItems")
	{
		orderItemRoutes.GET("/", GetOrderItems())
		orderItemRoutes.GET("/:orderId", GetOrderItem())
		orderItemRoutes.GET("/order/orderId", GetOrderItemsByOrder())
		orderItemRoutes.POST("/", CreateOrderItem())
		orderItemRoutes.PATCH("/:orderId", UpdateOrderItem())
	}

	invoiceRoutes := router.Group("/invoices")
	{
		invoiceRoutes.GET("/", GetInvoices())
		invoiceRoutes.GET("/:invoiceId", GetInvoice())
		invoiceRoutes.POST("/", CreateInvoice())
		invoiceRoutes.PATCH("/:invoiceId", UpdateInvoice())
	}

	return router
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func authentication() {

}

func StartApi() {
	router := SetApiRouter()
	ginLambda = ginadapter.New(router)
	lambda.Start(Handler)
}
