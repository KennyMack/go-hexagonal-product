package businesscustomer

type CustomerService struct {
	Persistence CustomerPersistenceInterface
}

func NewCustomerService(Persistence CustomerPersistenceInterface) *CustomerService {
	return &CustomerService{Persistence: Persistence}
}

func (s *CustomerService) Create(name string, email string) (CustomerInterface, error) {
	customer := NewCustomer()
	customer.Name = name
	customer.Email = email

	_, err := customer.IsValid()

	if err != nil {
		return &Customer{}, err
	}

	result, err := s.Persistence.Save(customer)
	if err != nil {
		return &Customer{}, err
	}

	return result, nil
}

func (s *CustomerService) Save(customer CustomerInterface) (CustomerInterface, error) {
	customer, err := s.Persistence.Save(customer)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (s *CustomerService) Disable(customer CustomerInterface) (CustomerInterface, error) {
	err := customer.Disable()
	if err != nil {
		return &Customer{}, err
	}

	result, err := s.Persistence.Save(customer)

	if err != nil {
		return &Customer{}, err
	}

	return result, nil
}

func (s *CustomerService) Enable(customer CustomerInterface) (CustomerInterface, error) {
	err := customer.Enable()
	if err != nil {
		return &Customer{}, err
	}

	result, err := s.Persistence.Save(customer)

	if err != nil {
		return &Customer{}, err
	}

	return result, nil
}

func (s *CustomerService) Get(id string) (CustomerInterface, error) {
	customer, err := s.Persistence.Get(id)
	if err != nil {
		return nil, err
	}

	return customer, nil
}
