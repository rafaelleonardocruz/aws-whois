package pkg

//Resource used to store information related with AWS resources
type Resource struct {
	ResourceName string
	ResourceID   string
	PublicIP     string
}

//ResourceList will be used to store AWS API returns
var ResourceList []Resource
