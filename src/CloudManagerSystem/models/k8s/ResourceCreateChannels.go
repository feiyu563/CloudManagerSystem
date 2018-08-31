package k8s

import ("k8s.io/api/core/v1"
		 client "k8s.io/client-go/kubernetes")

type ResourceCreateChannels struct {
	ServiceCreate ServiceCreateChannel
}

type ServiceCreateChannel struct {
	//Param  chan []*v1.Service
	Result chan bool
	Eorror chan error
}

func GetServiceCreateChannel(client client.Interface, namespace string,paras []*v1.Service) ServiceCreateChannel{

	serviceCreate := ServiceCreateChannel{
		//Param:make(chan []*v1.Service,len(paras)),
		Result:make(chan bool,len(paras)),
		Eorror:make(chan error),
	}

	go func(){
		for _,v := range paras{
			_,err:=client.CoreV1().Services(namespace).Create(v)
			if err == nil {
				serviceCreate.Result <- true
			}else{
				serviceCreate.Result <- false
				serviceCreate.Eorror <- err
			}
		}
	}()

	return serviceCreate
}
