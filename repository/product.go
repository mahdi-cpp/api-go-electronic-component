package repository

import (
	"github.com/mahdi-cpp/api-go-electronic-component/model"
	"github.com/mahdi-cpp/api-go-electronic-component/utils"
	"gorm.io/gorm"
	"strings"
)

var db *gorm.DB

func InitUser() {
	//db.Create(&model.User{Username: "mahdiabdolmaleki", Email: "mahdi.cpp@gmail.com", Phone: "09355512619", Avatar: "2018-10-23_13-55-58_UTC_profile_pic.jpg", Biography: "go lang programmer"})
}

func Init() {
	db = db
	db.AutoMigrate(&model.Products{})
	db.AutoMigrate(&model.Project{})
}

func InitProducts() {
	db.Create(&model.Products{PartNumber: "ESP32", Category: "Modules", Manufacturer: "Espressif", Photos: []string{"Mqz-LcNWat88zyN5OeBFryYNjJHYcz6i", "Mqz-LcNWat88zyN5OeBFryYNjJHYcz6i"}, Price: 1144961, Description: "Small­ sized 2.4 GHz Wi­Fi (802.11 b/g/n) and Bluetooth 5 module\n\n4 MB flash in chip package\n\nOn­board PCB antenna\n\n15 GPIOs", Count: 3})
	db.Create(&model.Products{PartNumber: "ESP32-S3", Category: "Modules", Manufacturer: "Espressif", Photos: []string{"Mqz-LcNWat88zyN5OeBFryYNjJHYcz6i", "Mqz-LcNWat88zyN5OeBFryYNjJHYcz6i"}, Price: 1144961, Description: "Small­ sized 2.4 GHz Wi­Fi (802.11 b/g/n) and Bluetooth 5 module\n\n4 MB flash in chip package\n\nOn­board PCB antenna\n\n15 GPIOs", Count: 3})
	db.Create(&model.Products{PartNumber: "ESP32-C3", Category: "Modules", Manufacturer: "Espressif", Photos: []string{"Mqz-LcNWat88zyN5OeBFryYNjJHYcz6i", "Mqz-LcNWat88zyN5OeBFryYNjJHYcz6i"}, Price: 1144961, Description: "Small­ sized 2.4 GHz Wi­Fi (802.11 b/g/n) and Bluetooth 5 module\n\n4 MB flash in chip package\n\nOn­board PCB antenna\n\n15 GPIOs", Count: 3})
	db.Create(&model.Products{PartNumber: "ESP32", Category: "Modules", Manufacturer: "Espressif", Photos: []string{"Mqz-LcNWat88zyN5OeBFryYNjJHYcz6i", "Mqz-LcNWat88zyN5OeBFryYNjJHYcz6i"}, Price: 1144961, Description: "Small­ sized 2.4 GHz Wi­Fi (802.11 b/g/n) and Bluetooth 5 module\n\n4 MB flash in chip package\n\nOn­board PCB antenna\n\n15 GPIOs", Count: 3})

	db.Create(&model.Products{PartNumber: "TPS63020DSJR", Category: "Switching Voltage Regulators\n", Manufacturer: "Texas Instruments", Photos: []string{"0PLuOmYV0xYspZeqsFJLUzcT5jd9rJOE.jpg"}, Price: 331812, Description: "رگولاتور سوئیچینگ 1.2 تا 5.5 ولت 4 آمپر با پکیج VSON-14", Count: 7})
	//db.Create(&model.Products{PartNumber: "", Category: "", Manufacturer: "", Photo: ".jpg", Price: 1144961, Description: "", Count: 7})
	//db.Create(&model.Products{PartNumber: "", Category: "", Manufacturer: "", Photo: ".jpg", Price: 1144961, Description: "", Count: 7})
	//db.Create(&model.Products{PartNumber: "", Category: "", Manufacturer: "", Photo: ".jpg", Price: 1144961, Description: "", Count: 7})
	//db.Create(&model.Products{PartNumber: "", Category: "", Manufacturer: "", Photo: ".jpg", Price: 1144961, Description: "", Count: 7})
	//db.Create(&model.Products{PartNumber: "", Category: "", Manufacturer: "", Photo: ".jpg", Price: 1144961, Description: "", Count: 7})
	//db.Create(&model.Products{PartNumber: "", Category: "", Manufacturer: "", Photo: ".jpg", Price: 1144961, Description: "", Count: 7})
	//db.Create(&model.Products{PartNumber: "", Category: "", Manufacturer: "", Photo: ".jpg", Price: 1144961, Description: "", Count: 7})
	//db.Create(&model.Products{PartNumber: "", Category: "", Manufacturer: "", Photo: ".jpg", Price: 1144961, Description: "", Count: 7})
}

func AddProduct(product model.Products) error {

	product.Description = strings.Replace(product.Description, ";", " ", 10)
	product.Url = strings.Replace(product.Url, "~", "#", 5)
	product.Photos = utils.AddPhotos()
	product.Datasheet = utils.AddDatasheet()

	err := db.Debug().Create(&product).Error

	return err
}

func GetProducts() (model.ProductsEntity, error) {

	var entity model.ProductsEntity
	var products []model.Products
	var count int64

	db.Debug().
		Where("id > 0").
		Limit(50).
		Order("id DESC").
		Find(&products)

	db.Debug().Model(&model.Products{}).Where("id > 0").Count(&count)

	entity.Products = products
	entity.Count = count

	println("......Count:", count)

	return entity, nil
}
