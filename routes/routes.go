package routes

import (
	controllers "med-manager/controllers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	// Initialize controllers
	medicineController := controllers.NewMedicineController(db)

	// Medicine routes
	medicines := app.Group("/medicines")
	{
		medicines.Post("/", medicineController.CreateMedicine)
		medicines.Get("/", medicineController.GetAllMedicines)
		medicines.Get("/:id", medicineController.GetMedicine)
		medicines.Put("/:id", medicineController.UpdateMedicine)
		medicines.Delete("/:id", medicineController.DeleteMedicine)
	}

	// Medicine type routes
	medTypes := app.Group("/medtypes")
	{
		medTypes.Get("/", medicineController.GetAllMedTypes)
		medTypes.Get("/:id", medicineController.GetMedType)
		medTypes.Post("/", medicineController.CreateMedType)
		medTypes.Put("/:id", medicineController.UpdateMedType)
		medTypes.Delete("/:id", medicineController.DeleteMedType)
	}

	// Stock routes
	stockController := controllers.NewStockController(db)
	stock := app.Group("/stock")
	{
		stock.Post("/add", stockController.AddToStock)
		stock.Get("/additions", stockController.GetAllStockAdditions)

		stock.Post("/deduct", stockController.DeductFromStock)
		stock.Get("/deductions", stockController.GetAllStockDeductions)

		stock.Get("updations/:id", stockController.GetStockUpdation)
		stock.Put("updations/:id", stockController.UpdateStockUpdation)
		stock.Delete("updations/:id", stockController.DeleteStockUpdation)

		stock.Get("/medicine/:medicine_id", stockController.GetMedicineStockByMedicineID)
		stock.Get("/medicine/additions/:medicine_id", stockController.GetStockAdditionsByMedicineID)
		stock.Get("/medicine/deductions/:medicine_id", stockController.GetStockDeductionsByMedicineID)

	}

	// Patient routes
	patientController := controllers.NewPatientController(db)
	patients := app.Group("/patients")
	{
		patients.Post("/", patientController.CreatePatient)
		patients.Get("/", patientController.GetAllPatients)
		patients.Get("/:id", patientController.GetPatient)
		patients.Put("/:id", patientController.UpdatePatient)
		patients.Delete("/:id", patientController.DeletePatient)
		patients.Put("/undodelete/:id", patientController.UndoDeletePatient)
	}

	// Visit routes
	visits := app.Group("/visits")
	{
		visits.Post("/", patientController.CreateVisit)
		visits.Get("/", patientController.GetAllVisits)
		visits.Get("/:id", patientController.GetVisit)
		visits.Put("/:id", patientController.UpdateVisit)
		visits.Delete("/:id", patientController.DeleteVisit)
		visits.Get("/patient/:id", patientController.GetAllVisitsByPatientID)
	}
}
