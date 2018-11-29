package utils

import "sync"

var singleton *Singleton

type Singleton struct {
	sync.Once
	data interface{}
}

func newInstance(data interface{}) *Singleton {
	return &Singleton{
		data: data,
	}
}

func GetInstance(data interface{}) *Singleton {
	if singleton == nil {
		singleton.Once.Do(
			func() {
				newInstance(data)
			},
		)
	}
	return singleton
}
