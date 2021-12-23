package businesscustomer

import (
	"errors"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

const (
	DISABLED = "disabled"
	ENABLED = "enabled"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Customer struct {
	ID string `valid:"uuidv4"`
	Name string `valid:"required"`
	Email string `valid:"required"`
	Status string `valid:"required"`
	DeactivationReason string `valid:"optional"`
}

type CustomerInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetId() string
	GetName() string
	GetStatus() string
	GetEmail() string
	GetDeactivationReason() string
}

type CustomerServiceInterface interface {
	Get(id string) (CustomerInterface, error)
	Create(name string, price float64) (CustomerInterface, error)
	Enable(customer CustomerInterface) (CustomerInterface, error)
	Disable(customer CustomerInterface) (CustomerInterface, error)
}

type CustomerReaderInterface interface {
	Get(id string) (CustomerInterface, error)
}

type CustomerWriterInterface interface {
	Save(customer CustomerInterface) (CustomerInterface, error)
}

type CustomerPersistenceInterface interface {
	CustomerReaderInterface
	CustomerWriterInterface
}

func NewCustomer() *Customer {
	customer := Customer{
		ID: uuid.NewV4().String(),
		Status: ENABLED,
		DeactivationReason: "",
	}
	return &customer
}

func (p *Customer) IsValid() (bool, error) {
	if p.Status == "" {
		p.Status = DISABLED
	}

	if p.Name == "" {
		return false, errors.New("the name must be informed")
	}

	if p.Email == "" {
		return false, errors.New("the e-mail must be informed")
	}

	if p.Status != ENABLED && p.Status != DISABLED {
		return false, errors.New("the status must be enabled or disabled")
	}

	if p.Status == DISABLED && p.DeactivationReason == "" {
		return false, errors.New("the reason for deactivation must be informed")
	}

	if p.Status == ENABLED && p.DeactivationReason != "" {
		return false, errors.New("the reason for deactivation should not be informed")
	}

	_, err := govalidator.ValidateStruct(p)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *Customer) Enable() error {
	if p.Email != "" && p.Name != "" {
		p.Status = ENABLED
		p.DeactivationReason = ""
		return nil
	}
	return errors.New("the name and e-mail must be informed")
}

func (p *Customer) Disable() error {
	if p.DeactivationReason != "" {
		p.Status = DISABLED
		return nil
	}
	return errors.New("the reason for deactivation must be informed")
}

func (p *Customer) GetId() string {
	return p.ID
}

func (p *Customer) GetName() string {
	return p.Name
}

func (p *Customer) GetStatus() string {
	return p.Status
}

func (p *Customer) GetEmail() string {
	return p.Email
}

func (p *Customer) GetDeactivationReason() string {
	return p.DeactivationReason
}
