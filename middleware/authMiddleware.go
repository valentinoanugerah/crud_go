package middleware

import (
    "net/http"
    "os"
    "strings"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
            c.Abort()
            return
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        if tokenString == authHeader {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
            c.Abort()
            return
        }

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, jwt.ErrSignatureInvalid
            }
            return jwtSecret, nil
        })

        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            exp := int64(claims["exp"].(float64))
            if time.Now().Unix() > exp {
                c.JSON(http.StatusUnauthorized, gin.H{"error": "Token has expired"})
                c.Abort()
                return
            }
            
            // --- Perbaikan di sini ---
            // 1. Ambil role_id dari klaim
            roleID, roleOK := claims["role_id"].(float64)
            if !roleOK {
                c.JSON(http.StatusUnauthorized, gin.H{"error": "Role ID not found in token"})
                c.Abort()
                return
            }

            // 2. Simpan role_id ke context dengan kunci yang konsisten
            c.Set("user_role_id", uint(roleID)) // Menggunakan uint
            
            c.Next()
        } else {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
            c.Abort()
            return
        }
    }
}

func AdminOnlyMiddleware() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        // Ambil role_id dari context dengan kunci yang benar
        // Menggunakan MustGet akan menyebabkan panic jika key tidak ada, lebih aman Get
        roleID, exists := ctx.Get("user_role_id")
        if !exists {
            ctx.JSON(http.StatusForbidden, gin.H{"error": "Role ID not found"})
            ctx.Abort()
            return
        }
        
        // Konversi ke tipe data yang benar dan periksa
        if roleID.(uint) != 1 {
            ctx.JSON(http.StatusForbidden, gin.H{"error": "Forbidden, admin access only"})
            ctx.Abort()
            return
        }

        ctx.Next()
    }
}

func UserOnlyMiddleware() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        roleID, exists := ctx.Get("user_role_id")
        if !exists {
            ctx.JSON(http.StatusForbidden, gin.H{"error": "Role ID not found"})
            ctx.Abort()
            return
        }
        
        if roleID.(uint) != 2 {
            ctx.JSON(http.StatusForbidden, gin.H{"error": "Forbidden, user access only"})
            ctx.Abort()
            return
        }

        ctx.Next()
    }
}