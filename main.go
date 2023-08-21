// main.go
package main

import (
	"po/views"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() // สร้าง Router ด้วย Gin
	views.WD(router)        // เรียกใช้ฟังก์ชั่น WD และส่ง Router เข้าไป
	router.Run(":8080")     // เริ่มเซิร์ฟเวอร์ที่พอร์ต 8080
}
