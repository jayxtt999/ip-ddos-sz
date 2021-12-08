package service
//
//  RouteSubject
//  @Description: 处理集
//
type RouteSubject struct {
	observers []RouteInterFace
	context   string
}
func NewRouteSubject() *RouteSubject {
	return &RouteSubject{
		observers: make([]RouteInterFace, 0),
	}
}
//
//  Add
//  @Description: 加入处理列
//  @receiver s
//  @param o
//
func (s *RouteSubject) Add(o RouteInterFace) {
	s.observers = append(s.observers, o)
}

//
//  block
//  @Description: 封堵操作
//  @receiver s
//  @return int
//
func (s *RouteSubject) block() int {
	var ch = make(chan int)
	for _, o := range s.observers {
		go func(o RouteInterFace) {
			r := o.BlockIp()
			if r {
				ch <- 1
			} else {
				ch <- 0
			}
		}(o)
	}
	var result = 0 //成功次数计数器 0失败  等于任务长度成功 否则部分成功
	for i := 0; i < len(s.observers); i++ {
		result += <-ch
	}
	if result == 0 {
		return -1 //失败
	}

	if result == len(s.observers) {
		return 1 //成功
	} else {
		return 0 //部分成功
	}

}

func (s *RouteSubject) BlockIp(ip string) int {
	return s.block()
}
