package routes


import(
	"github.com/gin-gonic/gin"
	"github.com/valentinoanugerah/crud_go/controller"
	 "github.com/valentinoanugerah/crud_go/middleware"

)

func AuthRoutes(r *gin.Engine) {
    
    r.POST("/auth/register", controller.RegisterUser)
    r.POST("/auth/login", controller.Login)

    
    protected := r.Group("/api")
    // Terapkan middleware JWTAuthMiddleware pada grup ini
    protected.Use(middleware.AuthMiddleware())
    {
        // Rute ini hanya bisa diakses jika token valid
        protected.GET("/users", controller.GetUsers)
        // Tambahkan rute lain yang memerlukan otentikasi di sini


        // Grup rute khusus untuk Admin saja
        // Middleware AdminOnlyMiddleware() akan melindungi semua rute di dalam grup ini
        admin := protected.Group("/admin/products")
        admin.Use(middleware.AdminOnlyMiddleware())
        {
            // Hanya admin yang bisa membuat, update, dan menghapus produk
            admin.POST("/", controller.CreateProduct)
            admin.PUT("/:id", controller.UpdateProduct)
            admin.DELETE("/:id", controller.DeleteProduct)
        }
    }
}