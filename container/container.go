package container

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"sampleGoWebProject/config"
	models "sampleGoWebProject/models"
	repository "sampleGoWebProject/models/employee/repository/postgres"
	useCase "sampleGoWebProject/models/employee/usecase"
)

type Container struct {
	appConfig       *config.AppConfig
	initialized     bool
	db              *gorm.DB
	employeeUseCase models.EmployeeUseCase
}

func (c *Container) WireDependencies() {
	if c.initialized {
		fmt.Println("Already initialized....")
		return
	}
	c.appConfig = loadConfig()
	c.GetEmployeeUseCase()
	fmt.Println("Dependencies Injection Done!!")
	c.initialized = true
}

func (c *Container) GetEmployeeUseCase() models.EmployeeUseCase {
	if c.employeeUseCase == nil {
		c.employeeUseCase = useCase.NewEmployeeUseCase(repository.NewEmployeeRepository(c.GetDB()))
	}
	return c.employeeUseCase
}

func (c *Container) GetDB() *gorm.DB {
	if c.db == nil {
		con, _ := CreateDBConnection()
		c.db = con
	}
	return c.db
}

func CreateDBConnection() (*gorm.DB, error) {
	var err error
	con, err := gorm.Open("postgres", "postgres://sachin:123456@localhost:5432/mydb?sslmode=disable")
	if err != nil {
		fmt.Printf("Error {%v}", err)
		log.Fatal("DB Connection not found!!")
		return nil, err
	}
	return con, nil
}

func loadConfig() *config.AppConfig{
	config := &config.AppConfig{}
	fmt.Println("Inside Load Config")
	pwd, _ := os.Getwd()
	configFilePath := pwd + "/ws_skshukla_go_projects/sampleGoWebProject/config/config.yaml"
	fileVal, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Fatal(fmt.Sprintf("Config could not be loaded with error {%+v} from path {%s}!!", err, configFilePath))
	}
	err = yaml.Unmarshal(fileVal, config)
	if err != nil {
		log.Fatal("Config could not be Unmarshalled!!")
	}
	return config
}

func (c *Container) GetAppConfig()  config.AppConfig{
	return *c.appConfig
}