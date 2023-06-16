package mms

type Controller struct{}

func NewController()*Controller{
	return &Controller{}
}

func (c *Controller) start(){

}

func (c *Controller) Init(){
	c.start()
}