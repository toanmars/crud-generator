package cmd

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

// generateCmd đại diện cho lệnh generate
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate CRUD code",
	Long:  `Generate CRUD code for a given model.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Error: Model name is required.")
			return
		}
		modelName := args[0]
		generateCRUD(modelName)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

// Định nghĩa cấu trúc dữ liệu
type Field struct {
	Name    string
	Type    string
	JsonTag string
}

type Model struct {
	ModelName      string
	ModelNameLower string
	Fields         []Field
}

// generateCRUD thực hiện generate code CRUD
func generateCRUD(modelName string) {

	// Tạo dữ liệu mẫu
	model := Model{
		ModelName:      modelName,
		ModelNameLower: strings.ToLower(modelName),
		Fields: []Field{
			{Name: "ID", Type: "uint", JsonTag: "id"},
			{Name: "Name", Type: "string", JsonTag: "name"},
			{Name: "Email", Type: "string", JsonTag: "email"},
		},
	}

	// Tạo thư mục output nếu chưa tồn tại
	if err := os.MkdirAll("models", os.ModePerm); err != nil {
		fmt.Println("Error creating models directory:", err)
		return
	}

	if err := os.MkdirAll("repositories", os.ModePerm); err != nil {
		fmt.Println("Error creating repositories directory:", err)
		return
	}
	if err := os.MkdirAll("services", os.ModePerm); err != nil {
		fmt.Println("Error creating services directory:", err)
		return
	}
	if err := os.MkdirAll("handlers", os.ModePerm); err != nil {
		fmt.Println("Error creating handlers directory:", err)
		return
	}
	if err := os.MkdirAll("routes", os.ModePerm); err != nil {
		fmt.Println("Error creating routes directory:", err)
		return
	}

	// Generate model
	generateFile("templates/model.tmpl", "models/"+modelName+".go", model)
	generateFile("templates/repository.tmpl", "repositories/"+modelName+"_repository.go", model)
	generateFile("templates/service.tmpl", "services/"+modelName+"_service.go", model)
	generateFile("templates/handler.tmpl", "handlers/"+modelName+"_handler.go", model)
	generateFile("templates/routes.tmpl", "routes/"+modelName+"_routes.go", model)

	fmt.Println("CRUD code generated successfully!")
}

// generateFile tạo file từ template
func generateFile(templatePath, outputPath string, data interface{}) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	file, err := os.Create(outputPath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	if err := tmpl.Execute(file, data); err != nil {
		fmt.Println("Error executing template:", err)
		return
	}
}
