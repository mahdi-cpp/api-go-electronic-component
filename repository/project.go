package repository

import (
	"github.com/mahdi-cpp/api-go-electronic-component/model"
	"github.com/mahdi-cpp/api-go-electronic-component/utils"
	"strings"
)

func InitProjects() {
	db.Create(&model.Project{Name: "ESP32-S3", Category: "PCB", Photos: []string{"/static/images/ESP32-S3-WROOM-1-N4R2 (4MB).jpg"}, Description: "ESP32-S3 is a series of high-performance Wi-Fi and Bluetooth ® 5 (LE) SoCs.", Youtubes: []string{"swc23GK1PwI?si=eq1MOy0nTqAHPMJe", "aTYSqWPsQ1U?si=V9Cprqj7Fs8_llb0"}})
	db.Create(&model.Project{Name: "Quectel M65", Category: "PCB", Photos: []string{"/static/images/M65.png"}, Description: "", Youtubes: []string{"swc23GK1PwI?si=eq1MOy0nTqAHPMJe", "aTYSqWPsQ1U?si=V9Cprqj7Fs8_llb0"}})
	db.Create(&model.Project{Name: "Quectel MC65", Category: "PCB", Photos: []string{"/static/images/MC65.png"}, Description: "MC65 is a quad-band full-featured GSM/GPRS/GNSS module using LCC+LGA castellation package", Youtubes: []string{"swc23GK1PwI?si=eq1MOy0nTqAHPMJe", "aTYSqWPsQ1U?si=V9Cprqj7Fs8_llb0"}})
	db.Create(&model.Project{Name: "GPS", Category: "PCB", Photos: []string{"/static/images/GNSS-LC86L.png"}, Description: "Compact GNSS Module Integrating Patch Antenna", Youtubes: []string{"swc23GK1PwI?si=eq1MOy0nTqAHPMJe", "aTYSqWPsQ1U?si=V9Cprqj7Fs8_llb0"}})
	db.Create(&model.Project{Name: "trance", Category: "PCB", Photos: []string{"/static/images/trance.jpg"}, Description: "Covert 12V to 220V", Youtubes: []string{"swc23GK1PwI?si=eq1MOy0nTqAHPMJe", "aTYSqWPsQ1U?si=V9Cprqj7Fs8_llb0"}})
	db.Create(&model.Project{Name: "DRV8711", Category: "PCB", Photos: []string{"/static/images/image_5825.jpg"}, Description: "The DRV8711 device is a stepper motor controller\nthat uses external N-channel MOSFETs to drive a\nbipolar stepper motor or two brushed DC motors", Youtubes: []string{"swc23GK1PwI?si=eq1MOy0nTqAHPMJe", "aTYSqWPsQ1U?si=V9Cprqj7Fs8_llb0"}})

	db.Create(&model.Project{Name: "Printer 3D", Category: "Mechanic", Photos: []string{"/static/images/ESP32-S3-WROOM-1-N4R2 (4MB).jpg"}, Description: "ESP32-S3 is a series of high-performance Wi-Fi and Bluetooth ® 5 (LE) SoCs.", Youtubes: []string{"swc23GK1PwI?si=eq1MOy0nTqAHPMJe", "aTYSqWPsQ1U?si=V9Cprqj7Fs8_llb0"}})
	db.Create(&model.Project{Name: "CNC", Category: "Mechanic", Photos: []string{"/static/images/GNSS-LC86L.png"}, Description: "Compact GNSS Module Integrating Patch Antenna", Youtubes: []string{"swc23GK1PwI?si=eq1MOy0nTqAHPMJe", "aTYSqWPsQ1U?si=V9Cprqj7Fs8_llb0"}})

}

func AddProject(project model.Project) error {

	project.Description = strings.Replace(project.Description, ";", " ", 10)
	project.Photos = utils.AddPhotos()
	//project.Datasheet = addDatasheet()

	err := db.Debug().Create(&project).Error

	return err
}

func GetProject(id string) (model.Project, error) {

	var project model.Project

	db.Debug().
		Where("id", id).
		Find(&project)

	return project, nil
}
func GetProjectsByCategory(category string) ([]model.Project, error) {

	var projects []model.Project

	db.Debug().
		Where("category", category).
		Find(&projects)

	return projects, nil
}
