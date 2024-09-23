package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/yatoyun/todo-app/api/config"
	"github.com/yatoyun/todo-app/api/infrastructure/middleware"
	"github.com/yatoyun/todo-app/api/infrastructure/router/api"
	"log"
)

type Route struct {
	Method  string
	Route   string
	Handler func(*gin.Context)
}

func NewRouter(conn *sqlx.DB) *gin.Engine {
	cfg := config.LoadConfig()
	todoController := api.InitializeTodoController(conn)
	userController := api.InitializeUserController(conn)

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	authMiddleware := middleware.JWTMiddleware(cfg)
	apiGroup := r.Group("/api/v1")
	apiGroup.Use(authMiddleware)
	{
		todosGroup := apiGroup.Group("/todos")
		{
			RegisterRoutes(todosGroup, []Route{
				{"GET", "", todoController.GetTodos},
				{"GET", "/:id", todoController.GetTodoByID},
				{"POST", "", todoController.CreateTodo},
				{"PUT", "/:id", todoController.UpdateTodo},
				{"DELETE", "/:id", todoController.DeleteTodo},
			})
		}
		usersGroup := apiGroup.Group("/users")
		{
			RegisterRoutes(usersGroup, []Route{
				{"GET", "", userController.GetUsers},
				{"GET", "/:id", userController.GetUserByID},
				{"GET", "/auth0/:auth0_id", userController.GetUserByAuth0ID}, // 開発用
				{"POST", "", userController.CreateUser},
				{"PUT", "/:id", userController.UpdateUser},
				{"DELETE", "/:id", userController.DeleteUser},
			})
		}
	}
	return r
}

func RegisterRoutes(group *gin.RouterGroup, routes []Route) {
	for _, r := range routes {
		RegisterRoute(group, r.Method, r.Route, r.Handler)
	}
}

func RegisterRoute(group *gin.RouterGroup, method string, route string, handler func(*gin.Context)) {
	switch method {
	case "GET":
		group.GET(route, handler)
	case "POST":
		group.POST(route, handler)
	case "PUT":
		group.PUT(route, handler)
	case "PATCH":
		group.PATCH(route, handler)
	case "DELETE":
		group.DELETE(route, handler)
	default:
		log.Fatalf("Unsupported HTTP method: %s", method)
	}
}
